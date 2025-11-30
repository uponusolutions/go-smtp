package textsmtp

import (
	"io"
	"strconv"
	"strings"

	"github.com/uponusolutions/go-smtp"
)

type bdatReader struct {
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

// NewBdatReader creates a new bdat reader.
func NewBdatReader(arg string, maxMessageBytes int64, input io.Reader, nextCommand func() (string, string, error)) (io.Reader, error) {
	size, last, err := bdatArg(arg)
	if err != nil {
		return nil, err
	}

	return &bdatReader{
		maxMessageBytes: maxMessageBytes,
		size:            size,
		last:            last,
		bytesReceived:   0,
		input:           input,
		nextCommand:     nextCommand,
	}, nil
}

func (d *bdatReader) Read(b []byte) (int, error) {
	if d.size == 0 {
		if d.last {
			return 0, io.EOF
		}

		d.chunk = nil

		cmd, arg, err := d.nextCommand()
		if err != nil {
			if err == io.EOF {
				return 0, smtp.ErrConnection
			}
			return 0, err
		}

		switch cmd {
		case "RSET":
			return 0, smtp.Reset
		case "QUIT":
			return 0, smtp.Quit
		case "BDAT":
			d.size, d.last, err = bdatArg(arg)
			if err != nil {
				return 0, err
			}

			if d.last && d.size == 0 {
				return 0, io.EOF
			}
		default:
			return 0, smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "BDAT command expected")
		}
	}

	if d.maxMessageBytes != 0 && d.bytesReceived+d.size > d.maxMessageBytes {
		return 0, smtp.NewStatus(552, smtp.EnhancedCode{5, 3, 4}, "Max message size exceeded")
	}

	if d.chunk == nil {
		d.chunk = io.LimitReader(d.input, int64(d.size))
	}

	n, err := d.chunk.Read(b)
	d.bytesReceived += int64(n)
	d.size -= int64(n)

	// this isn't the end
	if err == io.EOF && !d.last {
		// stream broke in the middle
		if d.size > 0 {
			err = smtp.ErrConnection
		} else {
			err = nil
		}
	}

	return n, err
}
