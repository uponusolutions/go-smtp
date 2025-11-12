// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on the modifications from
// https://github.com/go-textproto/textproto/blob/v0/writer.go

package textsmtp

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

var (
	ending   = []byte("BDAT 0 LAST\r\n")
	prefix   = []byte("BDAT ")
	lastcrlf = []byte(" LAST\r\n")
	crlf     = []byte("\r\n")
)

// NewBdatWriter returns a writer that can be used to write bdat commands to w.
// The caller should close the BdatWriter before the next call to a method on w.
func NewBdatWriter(maxChunkSize int, writer *bufio.Writer, read func() error, size int) io.WriteCloser {
	return &bdatWriter{
		w:             writer,
		read:          read,
		maxChunkSize:  maxChunkSize,
		remainingSize: size,
		knowSize:      size > 0,
	}
}

type bdatWriter struct {
	w    *bufio.Writer
	read func() error
	// maximum bdat chunk size
	maxChunkSize int
	// size still expected to be written
	remainingSize int
	// if true then size is known, so remainingSize is respected
	knowSize bool
}

// Write writes bytes as multiple bdat commands split by max chunk size.
func (d *bdatWriter) Write(b []byte) (n int, err error) {
	var p int

	for d.maxChunkSize > 0 && len(b) > d.maxChunkSize {
		p, err = d.write(b[:d.maxChunkSize], false)
		n += p
		b = b[d.maxChunkSize:]

		if err != nil {
			return n, err
		}
	}

	if d.knowSize {
		d.remainingSize -= n + len(b)

		if d.remainingSize < 0 {
			return n, errors.New("got more bytes than expected, check length")
		}
	}

	p, err = d.write(b, d.knowSize && d.remainingSize == 0)
	n += p

	return n, err
}

// write writes b as bdat command and checks return
func (d *bdatWriter) write(b []byte, last bool) (n int, err error) {
	if err = d.bdat(len(b), last); err != nil {
		return n, err
	}

	if n, err = d.w.Write(b); err != nil {
		return n, err
	}

	if err = d.w.Flush(); err != nil {
		return n, err
	}

	if last {
		return n, nil
	}

	return n, d.read()
}

// write BDAT <SIZE> \r\n
func (d *bdatWriter) bdat(size int, last bool) (err error) {
	if _, err = d.w.Write(prefix); err != nil {
		return err
	}

	if _, err = d.w.Write([]byte(strconv.Itoa(size))); err != nil {
		return err
	}

	if last {
		if _, err = d.w.Write(lastcrlf); err != nil {
			return err
		}
	} else {
		if _, err = d.w.Write(crlf); err != nil {
			return err
		}
	}

	return nil
}

func (d *bdatWriter) Close() error {
	if d.knowSize {
		if d.remainingSize == 0 {
			return nil
		}
		return errors.New("got less bytes than expected, check length")
	}
	if _, err := d.w.Write(ending); err != nil {
		return err
	}
	return d.w.Flush()
}
