// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textsmtp

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"io"
	legacy "net/textproto"
	"testing"
)

func TestDotWriter(t *testing.T) {
	t.Run("Encode", func(t *testing.T) {
		var buf bytes.Buffer
		d := NewDotWriter(bufio.NewWriter(&buf))
		n, err := d.Write([]byte("abc\n.def\n..ghi\n.jkl\n."))
		if n != 21 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}
		d.Close()
		want := "abc\r\n..def\r\n...ghi\r\n..jkl\r\n..\r\n.\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})
}

func TestDotWriterCloseEmptyWrite(t *testing.T) {
	var buf bytes.Buffer
	d := NewDotWriter(bufio.NewWriter(&buf))
	n, err := d.Write([]byte{})
	if n != 0 || err != nil {
		t.Fatalf("Write: %d, %s", n, err)
	}
	d.Close()
	want := "\r\n.\r\n"
	if s := buf.String(); s != want {
		t.Fatalf("wrote %q; want %q", s, want)
	}
}

func TestDotWriterCloseNoWrite(t *testing.T) {
	var buf bytes.Buffer
	d := NewDotWriter(bufio.NewWriter(&buf))
	d.Close()
	want := "\r\n.\r\n"
	if s := buf.String(); s != want {
		t.Fatalf("wrote %q; want %q", s, want)
	}
}

func BenchmarkDotWriter(b *testing.B) {
	const size = 4 * 1024 * 1024
	data, _ := io.ReadAll(io.LimitReader(rand.Reader, size))

	b.Run("Legacy", func(b *testing.B) {
		b.ResetTimer()
		b.SetBytes(size)
		for i := 0; i < b.N; i++ {
			r := bytes.NewReader(data)
			w := legacy.NewWriter(bufio.NewWriter(io.Discard)).DotWriter()
			io.Copy(w, r)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ResetTimer()
		b.SetBytes(size)
		for i := 0; i < b.N; i++ {
			r := bytes.NewReader(data)
			w := NewDotWriter(bufio.NewWriter(io.Discard))
			io.Copy(w, r)
		}
	})
}
