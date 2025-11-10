// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on the modifications from
// https://github.com/go-textproto/textproto/blob/v0/writer.go

package textsmtp

import (
	"bufio"
	"io"
	"strconv"
)

var (
	maxChunkSize = 1048576
	last         = []byte("BDAT 0 LAST\r\n")
	prefix       = []byte("BDAT ")
	spacecrlf    = []byte(" \r\n")
)

// NewBdatWriter returns a writer that can be used to write bdat commands to w.
// The caller should close the BdatWriter before the next call to a method on w.
func NewBdatWriter(writer *bufio.Writer, read func() error) io.WriteCloser {
	return &bdatWriter{
		w:    writer,
		read: read,
	}
}

type bdatWriter struct {
	w    *bufio.Writer
	read func() error
}

func (d *bdatWriter) Write(b []byte) (n int, err error) {
	var p int

	for len(b) > maxChunkSize {
		p, err = d.write(b[:maxChunkSize])
		n += p
		b = b[maxChunkSize:]

		if err != nil {
			return n, err
		}
	}

	p, err = d.write(b)
	n += p

	return n, err
}

func (d *bdatWriter) write(b []byte) (n int, err error) {
	if _, err = d.w.Write(prefix); err != nil {
		return n, err
	}

	if _, err = d.w.Write([]byte(strconv.Itoa(len(b)))); err != nil {
		return n, err
	}

	if _, err = d.w.Write(spacecrlf); err != nil {
		return n, err
	}

	if n, err = d.w.Write(b); err != nil {
		return n, err
	}

	if err = d.w.Flush(); err != nil {
		return n, err
	}

	return n, d.read()
}

func (d *bdatWriter) Close() error {
	if _, err := d.w.Write(last); err != nil {
		return err
	}
	return d.w.Flush()
}
