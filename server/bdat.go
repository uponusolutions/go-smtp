package server

import (
	"io"
	"strconv"
	"strings"

	"github.com/uponusolutions/go-smtp"
)

type bdat struct {
	size            int64
	last            bool
	bytesReceived   int64
	maxMessageBytes int64
	input           io.Reader
	chunk           io.Reader
	nextCommand     func() (string, string, error)
}

func bdatArg(arg string) (int64, bool, error) {
	args := strings.Fields(arg)
	if len(args) == 0 {
		return 0, true, smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Missing chunk size argument")
	}
	if len(args) > 2 {
		return 0, true, smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Too many arguments")
	}

	last := false
	if len(args) == 2 {
		if !strings.EqualFold(args[1], "LAST") {
			return 0, true, smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unknown BDAT argument")
		}
		last = true
	}

	// ParseUint instead of Atoi so we will not accept negative values.
	size, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil || (size == 0 && !last) {
		return 0, true, smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Malformed size argument")
	}

	return int64(size), last, nil
}

func (d *bdat) Read(b []byte) (n int, err error) {
	for {
		if d.maxMessageBytes != 0 && d.bytesReceived+d.size > d.maxMessageBytes {
			return 0, smtp.NewStatus(552, smtp.EnhancedCode{5, 3, 4}, "Max message size exceeded")
		}

		if d.size > 0 {
			if d.chunk == nil {
				d.chunk = io.LimitReader(d.input, int64(d.size))
			}

			for {
				in, err := d.chunk.Read(b[n:])
				d.bytesReceived += int64(in)
				n = in + n

				// error while reading, not eof
				if err != nil && err != io.EOF {
					return n, err
				}

				// limit reader has ended, need to end or next command
				if err == io.EOF {
					d.chunk = nil
					d.size = 0
				}

				// buffer full => return
				if n == len(b) {
					return n, err
				}

				// next comand needed
				if err == io.EOF {
					break
				}
			}

		}

		if d.last {
			return n, io.EOF
		}

		cmd, arg, err := d.nextCommand()
		if err != nil {
			if err == io.EOF {
				return n, smtp.ErrConnection
			}
			return n, err
		}

		switch cmd {
		case "RSET":
			return n, smtp.Reset
		case "QUIT":
			return n, smtp.Quit
		case "BDAT":
			d.size, d.last, err = bdatArg(arg)
			if err != nil {
				return n, err
			}

			if d.last && d.size == 0 {
				return n, io.EOF
			}
		default:
			return n, smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "BDAT command expected")
		}
	}

}
