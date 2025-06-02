// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textsmtp_test

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"io"
	legacy "net/textproto"
	"testing"

	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

func TestDotWriter(t *testing.T) {
	t.Run("Encode", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
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

	t.Run("EncodeNoError", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
		n, err := d.Write([]byte("abc\r\n.def\r\n..ghi\r\n.jkl\r\n."))
		if n != 25 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}
		d.Close()
		want := "abc\r\n..def\r\n...ghi\r\n..jkl\r\n..\r\n.\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("EncodeSeparateWrites", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
		n, err := d.Write([]byte("abc\r"))
		if n != 4 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		n, err = d.Write([]byte("\n.def\r\n..ghi\r\n.jkl\r\n."))
		if n != 21 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		d.Close()
		want := "abc\r\n..def\r\n...ghi\r\n..jkl\r\n..\r\n.\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("EncodeSeparateWritesEndsWithR", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
		n, err := d.Write([]byte("abc\r"))
		if n != 4 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		n, err = d.Write([]byte("\n.def\r\n..ghi\r\n.jkl\r"))
		if n != 19 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		d.Close()
		// \r gets interpreted as line break
		want := "abc\r\n..def\r\n...ghi\r\n..jkl\r\n.\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("EncodeSeparateWritesContainsR", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
		n, err := d.Write([]byte("abc\r"))
		if n != 4 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		n, err = d.Write([]byte("\n.def\r..ghi\r\n.jkl\r"))
		if n != 18 || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		d.Close()
		// \r gets interpreted as normal character
		want := "abc\r\n..def\r..ghi\r\n..jkl\r\n.\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})
}

func TestDotWriterCloseEmptyWrite(t *testing.T) {
	var buf bytes.Buffer
	d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
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
	d := textsmtp.NewDotWriter(bufio.NewWriter(&buf))
	d.Close()
	want := "\r\n.\r\n"
	if s := buf.String(); s != want {
		t.Fatalf("wrote %q; want %q", s, want)
	}
}

func BenchmarkDotWriter(b *testing.B) {
	const size = 256 * 1024 * 1024
	data, _ := io.ReadAll(io.LimitReader(rand.Reader, size))

	b.Run("Legacy", func(b *testing.B) {
		b.ResetTimer()
		b.SetBytes(size)
		for b.Loop() {
			r := bytes.NewReader(data)
			w := legacy.NewWriter(bufio.NewWriter(io.Discard)).DotWriter()
			io.Copy(w, r)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ResetTimer()
		b.SetBytes(size)
		for b.Loop() {
			r := bytes.NewReader(data)
			w := textsmtp.NewDotWriter(bufio.NewWriter(io.Discard))
			io.Copy(w, r)
		}
	})
}
