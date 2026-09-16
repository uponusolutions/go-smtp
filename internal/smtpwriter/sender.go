package smtpwriter

import (
	"bufio"
	"io"
)

// Sender is used as a wrapper around a connection to write to it.
type Sender struct {
	*bufio.Writer
}

// NewSender creates a new writer wrapper.
func NewSender(
	conn io.Writer,
	writerSize int,
) *Sender {
	if writerSize == 0 {
		writerSize = 4096 // default
	}

	return &Sender{
		Writer: bufio.NewWriterSize(conn, writerSize),
	}
}
