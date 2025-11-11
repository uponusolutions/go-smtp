package textsmtp

import (
	"bufio"
	"bytes"
	"io"

	"github.com/uponusolutions/go-smtp"
)

var crlfdot = []byte{'\r', '\n', '.'}

type dotReader struct {
	r       *bufio.Reader
	state   int
	limited bool
	n       int64 // Maximum bytes remaining.
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

func noCrlfDotFound(err error, b []byte, c []byte) int {
	if err == nil {
		l := len(c)

		if l > 1 && c[l-2] == '\r' && c[l-1] == '\n' {
			// Ends with \r\n, write everything before.
			return copy(b, c[:l-2])
		}

		if l > 0 && c[l-1] == '\r' {
			// Ends with \r, write everything before.
			return copy(b, c[:l-1])
		}
	}

	return copy(b, c)
}

const (
	stateBeginLine = iota // Beginning of line; initial state; must be zero.
	stateCR               // Wrote \r.
	stateEOF              // Reached .\r\n end marker line.
)

// Read reads in some more bytes.
// Run data through a simple state machine to
// elide leading dots and detect End-of-Data
// (<CR><LF>.<CR><LF>) line.
func (r *dotReader) Read(b []byte) (int, error) {
	if r.limited {
		if r.n <= 0 {
			return 0, smtp.ErrDataTooLarge
		}

		if int64(len(b)) > r.n {
			b = b[0:r.n]
		}
	}

	var n int       // Data written to b.
	var skipped int // How many.

	// IMPORTANT: We cannot wait on read,
	// because no EOL returns.
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

			// No full \r\n. found.
			if i == -1 {
				n += noCrlfDotFound(err, b, c)
				break
			}

			if len(c)-1 < i+4 {
				// i is \r, \n.\r\n needs to be accessible
				if err != nil {
					// No more data, just read to the end.
					n += copy(b, c[:i+2])
					skipped++
				} else if i > 0 {
					// Not enough bytes to check for \r\n.\r\n,
					// write everything before
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
					r.state = stateCR // Next time we want to write '\n'.
					skipped--         // Prevent \r from being discarded
				}
				break
			}

			// The end \r\n.\n\r
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

	// n + skipped is always smaller then what was peeked,
	// so it is guaranteed to work
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
