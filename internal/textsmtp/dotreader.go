package textsmtp

import (
	"bufio"
	"bytes"
	"io"

	"github.com/uponusolutions/go-smtp"
)

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

const (
	stateBegin = iota // Initial state, beginning of first line.
	stateLine         // Somewhere inside a line.
	stateCR           // Wrote \r.
	stateEOF          // Reached .\r\n end marker line.
)

// Read reads in some more bytes.
// Run data through a simple state machine to
// elide leading dots and detect End-of-Data
// (<CR><LF>.<CR><LF>) line.
func (r *dotReader) Read(b []byte) (int, error) {
	if r.state == stateEOF {
		return 0, io.EOF
	}

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

	c, err := r.peek(len(b))

	// To reach this state it is necessary that 5 bytes are still in the buffer ready to be consumed.
	// So there shouldn't be a case where err happens and there aren't 5 bytes in c at this point.
	switch r.state {
	case stateCR:
		// write \n
		b[0] = '\n'
		n++
		skipped += 2
		if c[3] == '\r' && c[4] == '\n' {
			r.state = stateEOF
			skipped += 2 // skip .\n\r
			return r.finalize(n, skipped, err)
		}
		r.state = stateLine
		b = b[1:]
		c = c[3:]
	case stateBegin:
		// nothing to do
		if len(c) == 0 {
			return r.finalize(n, skipped, err)
		}

		if c[0] == '.' {
			// need to check .\r\n
			if len(c) < 3 {
				return r.finalize(n, skipped, err)
			}

			// already finish, empty data
			if c[1] == '\r' && c[2] == '\n' {
				skipped += 3
				r.state = stateEOF
				return r.finalize(n, skipped, err)
			}

			skipped++
			c = c[1:]
		}
		r.state = stateLine
	default:
	}

	for {
		i := bytes.Index(c, smtp.Crlfdot)

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

	return r.finalize(n, skipped, err)
}

func (r *dotReader) finalize(n int, skipped int, err error) (int, error) {
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

func (r *dotReader) peek(blen int) ([]byte, error) {
	minimumPeek := 5
	if r.state == stateBegin {
		minimumPeek = 3
	}

	// IMPORTANT: We cannot wait on read,
	// because no EOL returns. So we call peek with minimumPeek to fill the buffer probably with more
	// to get as much data as possible in the second peek.
	if r.r.Buffered() < minimumPeek {
		_, _ = r.r.Peek(minimumPeek)
	}
	// min 5, max buffer size, default len(b)
	return r.r.Peek(max(min(blen, r.r.Buffered()), minimumPeek))
}
