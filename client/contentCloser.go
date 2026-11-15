package client

import (
	"errors"
	"io"
	"time"

	"github.com/uponusolutions/go-smtp"
)

// ContentCloser writes the message content.
// The whole transfer, until the reply to the end of the message, has to finish within the data timeout.
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
	if d.c.conn == nil {
		return 0, ErrNotConnected
	}
	n, err = d.writer.Write(p)
	// The message is only partly sent, the connection can't be used anymore.
	// A rejected BDAT chunk (*smtp.Status) leaves it in sync.
	if _, ok := err.(*smtp.Status); err != nil && !ok {
		_ = d.c.Close()
	}
	return n, err
}

// CloseWithResponse closes the data closer and returns code, msg.
func (d *ContentCloser) CloseWithResponse() (*smtp.Status, error) {
	if d.closed {
		return nil, errors.New("smtp: data writer closed twice")
	}
	d.closed = true

	// e.g. closed after a malformed reply to a BDAT chunk
	if d.c.conn == nil {
		return nil, ErrNotConnected
	}

	// clear the deadline set by newContentCloser
	defer func() {
		if d.c.conn != nil {
			_ = d.c.conn.SetDeadline(time.Time{})
		}
	}()

	if err := d.writer.Close(); err != nil {
		if _, ok := err.(*smtp.Status); !ok {
			_ = d.c.Close()
		}
		return nil, err
	}

	status, err := d.c.receiveResponse()

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

// newContentCloser starts a transfer of the message content,
// the deadline of the data timeout covers the whole transfer.
func (c *Client) newContentCloser(writer io.WriteCloser) *ContentCloser {
	var deadline time.Time
	if c.cfg.dataTimeout > 0 {
		deadline = time.Now().Add(c.cfg.dataTimeout)
	}
	_ = c.conn.SetDeadline(deadline)
	return &ContentCloser{c: c, writer: writer}
}
