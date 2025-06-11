// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textsmtp_test

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"embed"
	"io"
	legacy "net/textproto"
	"reflect"
	"strings"
	"testing"

	"github.com/uponusolutions/go-smtp/internal/textsmtp"
	"github.com/uponusolutions/go-smtp/tester"
)

//go:embed testdatareader/*
var embedFSReader embed.FS

func TestDotReader(t *testing.T) {
	t.Run("CompareTest", func(t *testing.T) {
		tester.ReaderCompareTest(t, &embedFSReader, "testdatareader", func(b io.Reader) ([]byte, error) {
			reader := legacy.NewReader(bufio.NewReader(b)).DotReader()
			buf, err := io.ReadAll(reader)
			buf = bytes.ReplaceAll(buf, []byte("\n"), []byte("\r\n"))
			return buf, err
		}, func(b io.Reader) ([]byte, error) {
			reader := textsmtp.NewDotReader(bufio.NewReader(b), 0) // textsmtp.NewDotReader(bufio.NewReader(b), 999999)
			return io.ReadAll(reader)
		})
	})

	t.Run("Decode", func(t *testing.T) {
		buf := bufio.NewReader(strings.NewReader("dotlines\r\n.foo\r\n..bar\n...baz\nquux\r\n\r\n.\r\nanot.her\n"))
		r := textsmtp.NewDotReader(buf, 0)
		b, err := io.ReadAll(r)
		want := []byte("dotlines\r\nfoo\r\n.bar\n...baz\nquux\r\n\r\n")

		if !reflect.DeepEqual(b, want) || err != nil {
			t.Fatalf("ReadDotBytes: %q, %v", b, err)
		}

		r = textsmtp.NewDotReader(buf, 0)
		b, err = io.ReadAll(r)
		want = []byte("anot.her\n")
		if !reflect.DeepEqual(b, want) || err != io.ErrUnexpectedEOF {
			t.Fatalf("ReadDotBytes2: %q, %v", b, err)
		}
	})
}

func BenchmarkDotReader(b *testing.B) {
	const size = 4 * 1024 * 1024
	var buf bytes.Buffer
	w := legacy.NewWriter(bufio.NewWriter(&buf)).DotWriter()
	_, _ = io.Copy(w, io.LimitReader(rand.Reader, size))
	data := buf.Bytes()

	b.Run("Legacy", func(b *testing.B) {
		b.ResetTimer()
		b.SetBytes(size)
		for b.Loop() {
			r := legacy.NewReader(bufio.NewReader(bytes.NewReader(data))).DotReader()
			_, _ = io.Copy(io.Discard, r)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ResetTimer()
		b.SetBytes(size)
		for b.Loop() {
			r := textsmtp.NewDotReader(bufio.NewReader(bytes.NewReader(data)), 0)
			_, _ = io.Copy(io.Discard, r)
		}
	})
}
