package smtpwriter_test

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/internal/smtpwriter"
)

func TestBdatWriterWithoutSize(t *testing.T) {
	t.Run("TwoWrites", func(t *testing.T) {
		var buf bytes.Buffer
		d := smtpwriter.NewBdat(bufio.NewWriter(&buf), func() error { return nil }, 0)

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

		want := "BDAT " + strconv.Itoa(len(input1)) + "\r\n" + string(input1) +
			"BDAT " + strconv.Itoa(len(input2)) + "\r\n" + string(input2) +
			"BDAT 0 LAST\r\n"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})

	t.Run("WithError", func(t *testing.T) {
		var buf bytes.Buffer
		d := smtpwriter.NewBdat(bufio.NewWriter(&buf), func() error { return errors.New("failed") }, 0)
		n, err := d.Write([]byte("ab"))
		require.Equal(t, 2, n)
		require.ErrorContains(t, err, "failed")
	})
}

func TestBdatWriterError(t *testing.T) {
	t.Run("ErrorTooMuch", func(t *testing.T) {
		var buf bytes.Buffer
		d := smtpwriter.NewBdat(bufio.NewWriter(&buf), func() error { return nil }, 1)

		input1 := []byte("a")
		n, err := d.Write(input1)
		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("b")
		_, err = d.Write(input2)
		require.ErrorContains(t, err, "got more bytes")
	})

	// Missing bytes are filled with \x00
	t.Run("ErrorTooLess", func(t *testing.T) {
		size := 3

		var buf bytes.Buffer
		d := smtpwriter.NewBdat(bufio.NewWriter(&buf), func() error { return nil }, size)

		input1 := []byte("a")
		n, err := d.Write(input1)

		require.NoError(t, err)
		require.Equal(t, len(input1), n)

		if n != len(input1) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input2 := []byte("b")
		n, err = d.Write(input2)

		require.NoError(t, err)
		require.Equal(t, len(input2), n)

		require.NoError(t, d.Close())

		want := "BDAT " + strconv.Itoa(size) + " LAST\r\n" + string(input1) + string(input2) + "\x00"
		require.Equal(t, want, buf.String())
	})
}

func TestBdatWriterWithSize(t *testing.T) {
	t.Run("TwoWrites", func(t *testing.T) {
		var buf bytes.Buffer
		d := smtpwriter.NewBdat(bufio.NewWriter(&buf), func() error { return nil }, 2)

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

	t.Run("ManyWrites", func(t *testing.T) {
		var buf bytes.Buffer
		d := smtpwriter.NewBdat(bufio.NewWriter(&buf), func() error { return nil }, 7)

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

		input3 := []byte("ef")
		n, err = d.Write(input3)
		if n != len(input3) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		input4 := []byte("g")
		n, err = d.Write(input4)
		if n != len(input4) || err != nil {
			t.Fatalf("Write: %d, %s", n, err)
		}

		require.NoError(t, d.Close())

		want := "BDAT 7 LAST\r\nabcdefg"
		if s := buf.String(); s != want {
			t.Fatalf("wrote %q", s)
		}
	})
}
