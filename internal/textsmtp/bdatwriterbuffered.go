// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on the modifications from
// https://github.com/go-textproto/textproto/blob/v0/writer.go

package textsmtp

import (
	"bufio"
	"io"
)

// NewBdatWriterBuffered returns a writer that can be used to write bdat commands to w.
// The caller should close the BdatWriter before the next call to a method on w.
func NewBdatWriterBuffered(maxChunkSize int, writer *bufio.Writer, read func() error, size int, buffer []byte) io.WriteCloser {
	return &bdatWriterBuffered{
		writer: bdatWriter{
			w:             writer,
			read:          read,
			maxChunkSize:  maxChunkSize,
			remainingSize: size,
			knownSize:     size > 0,
		},
		buffer:   buffer,
		position: 0,
	}
}

type bdatWriterBuffered struct {
	buffer   []byte
	position int
	writer   bdatWriter
}

// ReadFrom implements io.ReadFrom.
func (d *bdatWriterBuffered) ReadFrom(r io.Reader) (n int64, err error) {
	var p int

	for err == nil {
		p, err = r.Read(d.buffer[d.position:])
		n += int64(p)
		if p+d.position >= len(d.buffer) {
			_, err = d.writer.Write(d.buffer)
			d.position = 0
		} else {
			d.position += p
		}
	}

	// io.EOF is not returned (see io.Copy)
	if err == io.EOF {
		return n, nil
	}

	return n, err
}

// Write writes bytes as multiple bdat commands split by max chunk size.
func (d *bdatWriterBuffered) Write(b []byte) (n int, err error) {
	var p int

	for n < len(b) {
		p, err = d.write(b[n:])
		n += p
		if err != nil {
			return n, err
		}
	}

	return n, err
}

func (d *bdatWriterBuffered) write(b []byte) (n int, err error) {
	available := len(d.buffer) - d.position

	if available == 0 {
		_, err = d.writer.Write(d.buffer)
		d.position = 0
	}

	if len(b) > available {
		// ignore buffer, just pass it directly
		if d.position == 0 {
			n, err = d.writer.Write(b)
			return n, err
		}
		n = available
		copy(d.buffer[d.position:], b[:available])
	} else {
		n = len(b)
		copy(d.buffer[d.position:], b)
	}

	d.position += n

	return n, err
}

func (d *bdatWriterBuffered) Close() error {
	if d.position > 0 {
		// force bdat last, if size is unknown
		if !d.writer.knownSize {
			d.writer.knownSize = true
			d.writer.remainingSize = d.position
		}
		if _, err := d.writer.Write(d.buffer[:d.position]); err != nil {
			return err
		}
		d.position = 0
	}

	return d.writer.Close()
}
