// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textsmtp_test

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

func TestBdatWriterBuffered(t *testing.T) {
	t.Run("WithoutChunkSize", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriterBuffered(0, bufio.NewWriter(&buf), func() error { return nil }, 0, make([]byte, 1048576*2))

		input1 := []byte("a")
		n, err := d.Write(input1)
		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("b")
		n, err = d.Write(input2)
		if n != len(input2) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}
		require.NoError(t, d.Close())

		// it's buffered, so he always knows when the last chunk is send
		want := "BDAT " + strconv.Itoa(len(input1)+len(input2)) + " LAST\r\n" + string(input1) + string(input2)
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("WithChunkSize", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriterBuffered(1, bufio.NewWriter(&buf), func() error { return nil }, 0, make([]byte, 1))

		input1 := []byte("ab")
		n, err := d.Write(input1)
		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("cd")
		n, err = d.Write(input2)
		if n != len(input2) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}
		require.NoError(t, d.Close())

		// it's buffered but fast path prevents last detection (edge case on small buffers)
		want := "BDAT 1\r\naBDAT 1\r\nbBDAT 1\r\ncBDAT 1\r\ndBDAT 0 LAST\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("WithError", func(t *testing.T) {
		var buf bytes.Buffer

		// read is never called as buffer lead to a single bdat
		d := textsmtp.NewBdatWriterBuffered(0, bufio.NewWriter(&buf), func() error { return errors.New("failed") }, 0, make([]byte, 1048576*2))
		n, err := d.Write([]byte("ab"))
		require.Equal(t, 2, n)
		require.NoError(t, err)
		err = d.Close()
		require.NoError(t, err)

		d = textsmtp.NewBdatWriterBuffered(1, bufio.NewWriter(&buf), func() error { return errors.New("failed") }, 0, make([]byte, 1))
		n, err = d.Write([]byte("ab"))
		require.Equal(t, 1, n)
		require.ErrorContains(t, err, "failed")
	})

	t.Run("WithSize", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriterBuffered(0, bufio.NewWriter(&buf), func() error { return nil }, 2, make([]byte, 1048576*2))

		input1 := []byte("a")
		n, err := d.Write(input1)
		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("b")
		n, err = d.Write(input2)
		if n != len(input2) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}
		require.NoError(t, d.Close())

		want := "BDAT " + strconv.Itoa(len(input1)+len(input2)) + " LAST\r\n" + string(input1) + string(input2)
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("WithSizeErrorTooMuch", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriterBuffered(0, bufio.NewWriter(&buf), func() error { return nil }, 1, make([]byte, 1048576*2))

		input1 := []byte("a")
		n, err := d.Write(input1)
		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("b")
		_, err = d.Write(input2)
		require.NoError(t, err)
		require.ErrorContains(t, d.Close(), "got more bytes")
	})

	t.Run("WithSizeErrorTooLess", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriterBuffered(0, bufio.NewWriter(&buf), func() error { return nil }, 3, make([]byte, 1048576*2))

		input1 := []byte("a")
		n, err := d.Write(input1)
		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("b")
		if n != len(input2) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		require.ErrorContains(t, d.Close(), "got less bytes")
	})
}
