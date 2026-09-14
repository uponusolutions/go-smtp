package upstream

import (
	"bufio"
	"io"

	"github.com/uponusolutions/go-smtp"
)

// Copied and integrated over from
// https://github.com/emersion/go-smtp/blob/8d5af0d9db3ace5e4fdc0e8b427f5b157b0c6c44/data.go
// to serve as the upstream comparison.

// DotReader implements a smtp dot reader.
type DotReader struct {
	r     *bufio.Reader
	state int

	limited bool
	n       int64 // Maximum bytes remaining
}

// NewDotReader creates a new smtp dot reader.
func NewDotReader(r *bufio.Reader, maxMessageBytes int) *DotReader {
	dr := &DotReader{
		r: r,
	}

	if maxMessageBytes > 0 {
		dr.limited = true
		dr.n = int64(maxMessageBytes)
	}

	return dr
}

func (r *DotReader) Read(b []byte) (n int, err error) {
	if r.limited {
		if r.n <= 0 {
			return 0, smtp.ErrDataTooLarge
		}
		if int64(len(b)) > r.n {
			b = b[0:r.n]
		}
	}

	// Code below is taken from net/textproto with only one modification to
	// not rewrite CRLF -> LF.

	// Run data through a simple state machine to
	// elide leading dots and detect End-of-Data (<CR><LF>.<CR><LF>) line.
	const (
		stateBeginLine = iota // beginning of line; initial state; must be zero
		stateDot              // read . at beginning of line
		stateDotCR            // read .\r at beginning of line
		stateCR               // read \r (possibly at end of line)
		stateData             // reading data in middle of line
		stateEOF              // reached .\r\n end marker line
	)
	for n < len(b) && r.state != stateEOF {
		var c byte
		c, err = r.r.ReadByte()
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			break
		}
		switch r.state {
		case stateBeginLine:
			if c == '.' {
				r.state = stateDot
				continue
			}
			if c == '\r' {
				r.state = stateCR
				break
			}
			r.state = stateData
		case stateDot:
			if c == '\r' {
				r.state = stateDotCR
				continue
			}
			r.state = stateData
		case stateDotCR:
			if c == '\n' {
				r.state = stateEOF
				continue
			}
			r.state = stateData
		case stateCR:
			if c == '\n' {
				r.state = stateBeginLine
				break
			}
			r.state = stateData
		case stateData:
			if c == '\r' {
				r.state = stateCR
			}
		default:
		}
		b[n] = c
		n++
	}
	if err == nil && r.state == stateEOF {
		err = io.EOF
	}

	if r.limited {
		r.n -= int64(n)
	}
	return n, err
}
