// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textsmtp

import (
	"bufio"
	"bytes"
	"io"
	"unsafe"
)

var crnl = []byte{'\r', '\n'}
var dotcrnl = []byte{'.', '\r', '\n'}

// DotWriter returns a writer that can be used to write a dot-encoding to w.
// It takes care of inserting leading dots when necessary,
// translating line-ending \n into \r\n, and adding the final .\r\n line
// when the DotWriter is closed. The caller should close the
// DotWriter before the next call to a method on w.
//
// See the documentation for Reader's DotReader method for details about dot-encoding.
func NewDotWriter(writer *bufio.Writer) io.WriteCloser {
	return &dotWriter{
		W: writer,
	}
}

type dotWriter struct {
	W     *bufio.Writer
	state int
}

const (
	wstateBegin     = iota // starting state
	wstateBeginLine        // beginning of line
	wstateCR               // wrote \r (possibly at end of line)
	wstateData             // writing data in middle of line
)

func (d *dotWriter) Write(b []byte) (n int, err error) {
	var (
		i    int
		p    []byte
		pLen int
		bw   = d.W
	)
	for len(b) > 0 {
		i = bytes.IndexByte(b, '\n')
		if i >= 0 {
			p, b = b[:i+1], b[i+1:]
		} else {
			p, b = b, nil
		}
		pLen = len(p)
		if d.state == wstateBeginLine && p[0] == '.' {
			err = bw.WriteByte('.')
			if err != nil {
				return
			}
		}
		if (pLen >= 2 && *(*[2]byte)(unsafe.Pointer(&p[pLen-2])) == [2]byte{'\r', '\n'}) ||
			(d.state == wstateCR && pLen == 1 && p[0] == '\n') {
			d.state = wstateBeginLine
			if _, err = bw.Write(p); err != nil {
				return
			}
		} else if pLen >= 1 && p[pLen-1] == '\n' {
			d.state = wstateBeginLine
			_, _ = bw.Write(p[:pLen-1])
			if _, err = bw.Write(crnl); err != nil {
				return
			}
		} else if pLen >= 1 && p[pLen-1] == '\r' {
			d.state = wstateCR
			if _, err = bw.Write(p[:pLen-1]); err != nil {
				return
			}
		} else {
			d.state = wstateData
			if _, err = bw.Write(p); err != nil {
				return
			}
		}
		n += pLen
	}
	return
}

func (d *dotWriter) Close() error {
	bw := d.W
	switch d.state {
	default:
		bw.WriteByte('\r')
		fallthrough
	case wstateCR:
		bw.WriteByte('\n')
		fallthrough
	case wstateBeginLine:
		bw.Write(dotcrnl)
	}
	return bw.Flush()
}
