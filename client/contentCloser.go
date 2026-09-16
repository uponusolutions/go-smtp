package client

import (
	"errors"
	"io"

	"github.com/uponusolutions/go-smtp"
)

// ContentCloser function.
type ContentCloser struct {
	writer io.WriteCloser
	c      *Client
	closed bool
}

// Writer returns inner io.Writer which possible implement more methods (e.g. ReadFrom)
func (d *ContentCloser) Writer() io.Writer {
	return d.writer
}

// Write writes do underlying writer.
func (d *ContentCloser) Write(p []byte) (n int, err error) {
	return d.writer.Write(p)
}

// CloseWithResponse closes the data closer and returns code, msg.
func (d *ContentCloser) CloseWithResponse() (*smtp.Status, error) {
	if d.closed {
		return nil, errors.New("smtp: data writer closed twice")
	}
	d.closed = true

	if err := d.writer.Close(); err != nil {
		return nil, err
	}

	timeout := smtp.Timeout(d.c.conn, d.c.cfg.submissionTimeout)
	defer timeout()

	status, err := d.c.receiver.ReadResponse()

	if err == nil && status.Code != 250 {
		err = status
		status = nil
	}

	return status, err
}

// Close closes the data closer.
func (d *ContentCloser) Close() error {
	_, err := d.CloseWithResponse()
	return err
}
