package textsmtp

import (
	"bufio"
	"bytes"
	"io"

	"github.com/uponusolutions/go-smtp"
)

var crlfdot = []byte{'\r', '\n', '.'}

type dotReader struct {
	r     *bufio.Reader
	state int

	limited bool
	n       int64 // Maximum bytes remaining
}

// NewDotReader creates a new dot reader.
func NewDotReader(reader *bufio.Reader, maxMessageBytes int64) io.Reader {
	dr := &dotReader{
		r: reader,
	}

	if maxMessageBytes > 0 {
		dr.limited = true
		dr.n = maxMessageBytes
	}

	return dr
}

//go:inline
func noCrlfDotFound(err error, b []byte, c []byte) int {
	if err == nil && len(c) > 1 && c[len(c)-2] == '\r' && c[len(c)-1] == '\n' {
		// ends with \r\n, write everything before
		return copy(b, c[:len(c)-2])
	} else if err == nil && len(c) > 0 && c[len(c)-1] == '\r' {
		// ends with \r, write everything before
		return copy(b, c[:len(c)-1])
	}
	return copy(b, c)
}

// Read reads in some more bytes.
func (r *dotReader) Read(b []byte) (int, error) {
	// Run data through a simple state machine to
	// elide leading dots and detect End-of-Data (<CR><LF>.<CR><LF>) line.
	const (
		stateBeginLine = iota // beginning of line; initial state; must be zero
		stateCR               // wrote \r
		stateEOF              // reached .\r\n end marker line
	)

	if r.limited {
		if r.n <= 0 {
			return 0, smtp.ErrDataTooLarge
		}
		if int64(len(b)) > r.n {
			b = b[0:r.n]
		}
	}

	var n int       // data written to b
	var skipped int // how many

	// IMPORTANT: We cannot wait on read, because no EOL returns
	if r.r.Buffered() < 5 {
		_, _ = r.r.Peek(5)
	}

	// min 5, max buffer size, default len(b)
	c, err := r.r.Peek(max(min(len(b), r.r.Buffered()), 5))

	// write \n
	if r.state == stateCR {
		b[0] = '\n'
		n++
		skipped += 2
		r.state = stateBeginLine
		if c[3] == '\r' && c[4] == '\n' {
			r.state = stateEOF
			skipped += 2 // skip .\n\r
		} else {
			b = b[1:]
			c = c[3:]
		}
	}

	if r.state != stateEOF {
		for {
			i := bytes.Index(c, crlfdot)

			// no full \r\n. found
			if i == -1 {
				n += noCrlfDotFound(err, b, c)
				break
			}

			if len(c)-1 < i+4 {
				// i is \r, \n.\r\n needs to be accessible

				if err != nil {
					// no more data, just read to the end
					n += copy(b, c[:i+2])
					skipped++
				} else if i > 0 {
					// not enough bytes to check for \r\n.\r\n, write everything before
					n += copy(b, c[:i])
				}

				break
			}

			p := copy(b, c[:i+2])
			n += p

			// b was to small
			if p < i+2 {
				// we only wrote \r
				if i+2-p == 1 {
					r.state = stateCR // next time we want to write \n
					skipped--         // prevent \r from being discarded
				}
				break
			}

			// the end \r\n.\n\r
			if c[i+3] == '\r' && c[i+4] == '\n' {
				r.state = stateEOF
				skipped += 3 // skip .\r\n
				break
			}

			skipped++ // . isn't written
			b = b[i+2:]
			c = c[i+3:]
		}
	}

	// n + skipped is always smaller then what was peeked, so it is guaranteed to work
	_, _ = r.r.Discard(n + skipped)

	if err == io.EOF && r.state != stateEOF {
		err = io.ErrUnexpectedEOF
	} else if err == nil && r.state == stateEOF {
		err = io.EOF
	}

	if r.limited {
		r.n -= int64(n)
	}

	return n, err
}
