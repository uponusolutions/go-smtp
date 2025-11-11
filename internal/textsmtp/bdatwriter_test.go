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

func TestBdatWriter(t *testing.T) {
	t.Run("WithoutChunkSize", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriter(0, bufio.NewWriter(&buf), func() error { return nil })

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

		want := "BDAT " + strconv.Itoa(len(input1)) + " \r\n" + string(input1) + "BDAT " + strconv.Itoa(len(input2)) + " \r\n" + string(input2) + "BDAT 0 LAST\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("WithChunkSize", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriter(1, bufio.NewWriter(&buf), func() error { return nil })

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

		want := "BDAT 1 \r\naBDAT 1 \r\nbBDAT 1 \r\ncBDAT 1 \r\ndBDAT 0 LAST\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("WithError", func(t *testing.T) {
		var buf bytes.Buffer
		d := textsmtp.NewBdatWriter(0, bufio.NewWriter(&buf), func() error { return errors.New("failed") })
		n, err := d.Write([]byte("ab"))
		require.Equal(t, 2, n)
		require.ErrorContains(t, err, "failed")

		d = textsmtp.NewBdatWriter(1, bufio.NewWriter(&buf), func() error { return errors.New("failed") })
		n, err = d.Write([]byte("ab"))
		require.Equal(t, 1, n)
		require.ErrorContains(t, err, "failed")
	})
}
