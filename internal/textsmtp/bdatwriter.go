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
	ending = []byte("BDAT 0 LAST\r\n")
	prefix = []byte("BDAT ")
	suffix = []byte(" \r\n")
)

// NewBdatWriter returns a writer that can be used to write bdat commands to w.
// The caller should close the BdatWriter before the next call to a method on w.
func NewBdatWriter(maxChunkSize int, writer *bufio.Writer, read func() error) io.WriteCloser {
	return &bdatWriter{
		w:            writer,
		read:         read,
		maxChunkSize: maxChunkSize,
	}
}

type bdatWriter struct {
	w            *bufio.Writer
	read         func() error
	maxChunkSize int
}

// Write writes bytes as multiple bdat commands split by max chunk size.
func (d *bdatWriter) Write(b []byte) (n int, err error) {
	var p int

	for d.maxChunkSize > 0 && len(b) > d.maxChunkSize {
		p, err = d.write(b[:d.maxChunkSize])
		n += p
		b = b[d.maxChunkSize:]

		if err != nil {
			return n, err
		}
	}

	p, err = d.write(b)
	n += p

	return n, err
}

// write BDAT <SIZE> \r\n
func (d *bdatWriter) bdat(size int) (err error) {
	if _, err = d.w.Write(prefix); err != nil {
		return err
	}

	if _, err = d.w.Write([]byte(strconv.Itoa(size))); err != nil {
		return err
	}

	if _, err = d.w.Write(suffix); err != nil {
		return err
	}

	return nil
}

// write writes b as bdat command and checks return
func (d *bdatWriter) write(b []byte) (n int, err error) {
	if err = d.bdat(len(b)); err != nil {
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
	if _, err := d.w.Write(ending); err != nil {
		return err
	}
	return d.w.Flush()
}
