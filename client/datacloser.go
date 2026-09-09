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
func (d *DataCloser) CloseWithResponse() (*smtp.Status, error) {
	if d.closed {
		return nil, errors.New("smtp: data writer closed twice")
	}

	if err := d.writer.Close(); err != nil {
		return nil, err
	}

	timeout := smtp.Timeout(d.c.conn, d.c.cfg.submissionTimeout)
	defer timeout()

	status, err := d.c.cfg.text.ReadResponse()

	if err == nil && status.Code != 250 {
		err = status
		status = nil
	}

	d.closed = true
	return status, err
}

// Close closes the data closer.
func (d *DataCloser) Close() error {
	_, err := d.CloseWithResponse()
	return err
}
