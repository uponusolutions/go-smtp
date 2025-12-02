package client

import (
	"errors"
	"io"

	"github.com/uponusolutions/go-smtp"
)

// DataCloser implement an io.WriteCloser with the additional
// CloseWithResponse function.
type DataCloser struct {
	writer io.WriteCloser
	c      *Client
	closed bool
}

// Writer returns inner io.Writer which possible implement more methods (e.g. ReadFrom)
func (d *DataCloser) Writer() io.Writer {
	return d.writer
}

// Write writes do underlying writer.
func (d *DataCloser) Write(p []byte) (n int, err error) {
	return d.writer.Write(p)
}

// CloseWithResponse closes the data closer and returns code, msg.
func (d *DataCloser) CloseWithResponse() (code int, msg string, err error) {
	if d.closed {
		return 0, "", errors.New("smtp: data writer closed twice")
	}

	if err := d.writer.Close(); err != nil {
		return 0, "", err
	}

	timeout := smtp.Timeout(d.c.conn, d.c.submissionTimeout)
	defer timeout()

	code, msg, err = d.c.readResponse(250)

	d.closed = true
	return code, msg, err
}

// Close closes the data closer.
func (d *DataCloser) Close() error {
	_, _, err := d.CloseWithResponse()
	return err
}
