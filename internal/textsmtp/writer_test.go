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
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

func getAllFilenames(fs *embed.FS, path string) (out []string, err error) {
	if len(path) == 0 {
		path = "."
	}
	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}

func chunkSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func CheckAgainstOld(t *testing.T, b []byte) {
	var buf bytes.Buffer
	var err error

	f := legacy.NewWriter(bufio.NewWriter(&buf)).DotWriter()
	_, err = f.Write(b)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	size := 1

	for size < 4048 && len(b) >= size {
		bsplitted := chunkSlice(b, size)

		var buf1 bytes.Buffer
		f = legacy.NewWriter(bufio.NewWriter(&buf1)).DotWriter()

		for _, b := range bsplitted {
			_, err = f.Write(b)
			require.NoError(t, err)
		}

		require.NoError(t, f.Close())

		require.Equal(t, buf, buf1)

		size++
	}

}

//go:embed testdata/*
var embedFS embed.FS

func TestDotWriter(t *testing.T) {
	t.Run("Testdata", func(t *testing.T) {

		files, err := getAllFilenames(&embedFS, "testdata")
		require.NoError(t, err)

		for _, file := range files {
			dat, err := embedFS.ReadFile(file)
			require.NoError(t, err)
			CheckAgainstOld(t, []byte(dat))
		}

	})

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
