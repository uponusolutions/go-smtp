package textsmtp_test

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"embed"
	"errors"
	"io"
	"net"
	legacy "net/textproto"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
	"github.com/uponusolutions/go-smtp/tester"
)

//go:embed testdata/reader/*
var embedFSReader embed.FS

func TestDotReaderCompare(t *testing.T) {
	input := []string{
		"dotlines\r\n.foo\r\n..bar\n...baz\nquux\r\n\r\n.\r\nanot.her\n",
		"anot.her\n",
		"\r\n",
		".\r\n",
	}

	for p, value := range input {
		t.Run(strconv.Itoa(p), func(t *testing.T) {
			readerOld := bufio.NewReader(strings.NewReader(value))
			reader := bufio.NewReader(strings.NewReader(value))

			dotReaderOld := NewDotReader(readerOld, 0)
			bufOld := make([]byte, 1)

			dotReader := textsmtp.NewDotReader(reader, 0)
			buf := make([]byte, 1)

			i := 0

			for {
				nOld, errOld := dotReaderOld.Read(bufOld)
				n, err := dotReader.Read(buf)

				require.Equal(t, bufOld, buf, i)
				require.Equal(t, nOld, n, i)

				if errOld != nil && err != io.EOF {
					require.Equal(t, errOld, err, i)
				}

				i++

				if errOld == io.EOF || errOld == io.ErrUnexpectedEOF {
					break
				}
			}

			bOld, errOld := io.ReadAll(readerOld)
			b, err := io.ReadAll(reader)

			require.Equal(t, errOld, err)
			require.Equal(t, bOld, b)
		})
	}
}

func TestDotReader(t *testing.T) {
	t.Run("CompareTest", func(t *testing.T) {
		tester.ReaderCompareTest(t, &embedFSReader, "testdata/reader", func(b io.Reader) ([]byte, error) {
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
		require.NoError(t, err)
		require.Equal(t, []byte("dotlines\r\nfoo\r\n.bar\n...baz\nquux\r\n\r\n"), b)

		r = textsmtp.NewDotReader(buf, 0)
		b, err = io.ReadAll(r)
		require.Error(t, io.ErrUnexpectedEOF, err)
		require.Equal(t, []byte("anot.her\n"), b)
	})

	// A live SMTP peer waits for the reply without closing the write side.
	// Adapted from Jabberwocky238's testcase.
	// https://github.com/Jabberwocky238/go-smtp/blob/b0673510e58009b47a2c9b6e6ca5fc189c3c5ba4/data_test.go
	t.Run("EmptyLiveConnection", func(t *testing.T) {
		server, client := net.Pipe()
		defer func() {
			_ = server.Close()
			_ = client.Close()
		}()
		_ = server.SetReadDeadline(time.Now().Add(2 * time.Second))
		done := make(chan error, 1)
		go func() { _, err := io.WriteString(client, ".\r\n"); done <- err }()
		buf := bufio.NewReader(server)
		r := textsmtp.NewDotReader(buf, 0)
		got, err := io.ReadAll(r)

		require.NoError(t, err)
		require.Len(t, got, 0, "empty DATA")

		if err := <-done; err != nil {
			t.Fatal(err)
		}
	})

	// Adapted from Jabberwocky238's testcase.
	// https://github.com/Jabberwocky238/go-smtp/blob/b0673510e58009b47a2c9b6e6ca5fc189c3c5ba4/data_test.go
	t.Run("ZeroRead", func(t *testing.T) {
		buf := bufio.NewReader(strings.NewReader("..first\r\n.\r\nNEXT\r\n"))
		r := textsmtp.NewDotReader(buf, 0)
		if n, err := r.Read(nil); n != 0 || err != nil {
			t.Fatalf("zero read: %d, %v", n, err)
		}
		got, err := io.ReadAll(r)
		require.NoError(t, err)
		require.Equal(t, ".first\r\n", string(got))

		got, err = io.ReadAll(buf)
		require.NoError(t, err)
		require.Equal(t, "NEXT\r\n", string(got))
	})

	t.Run("Limit", func(t *testing.T) {
		input := "dotlines\r\n.foo\r\n..bar\n...baz\nquux\r\n\r\n.\r\nanot.her\n"

		buf := bufio.NewReader(strings.NewReader(input))
		r := textsmtp.NewDotReader(buf, 35)
		b, err := io.ReadAll(r)
		require.NoError(t, err)
		require.Equal(t, []byte("dotlines\r\nfoo\r\n.bar\n...baz\nquux\r\n\r\n"), b)

		buf = bufio.NewReader(strings.NewReader(input))
		r = textsmtp.NewDotReader(buf, 34)
		b, err = io.ReadAll(r)
		require.Error(t, smtp.ErrDataTooLarge, err)
		require.Equal(t, []byte("dotlines\r\nfoo\r\n.bar\n...baz\nquux\r\n\r"), b)
	})
}

func BenchmarkDotReader(b *testing.B) {
	const size = 4 * 1024 * 1024
	var buf bytes.Buffer
	w := legacy.NewWriter(bufio.NewWriter(&buf)).DotWriter()
	_, _ = io.Copy(w, io.LimitReader(rand.Reader, size))
	data := buf.Bytes()

	b.Run("Legacy", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(size)
		}
		for b.Loop() {
			r := legacy.NewReader(bufio.NewReader(bytes.NewReader(data))).DotReader()
			_, _ = io.Copy(io.Discard, r)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		b.ResetTimer()
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(size)
		}
		for b.Loop() {
			r := textsmtp.NewDotReader(bufio.NewReader(bytes.NewReader(data)), 0)
			_, _ = io.Copy(io.Discard, r)
		}
	})

	b.Run("LegacySimpleReader", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(size)
		}
		for b.Loop() {
			r := legacy.NewReader(bufio.NewReader(tester.NewBuffer(data))).DotReader()
			_, _ = io.Copy(io.Discard, r)
		}
	})

	b.Run("OptimizedSimpleReader", func(b *testing.B) {
		b.ResetTimer()
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(size)
		}
		for b.Loop() {
			r := textsmtp.NewDotReader(bufio.NewReader(tester.NewBuffer(data)), 0)
			_, _ = io.Copy(io.Discard, r)
		}
	})
}

func write(w io.Writer, d string) {
	go func() { _, _ = w.Write([]byte(d)) }()
}

func TestDotReaderBytes(t *testing.T) {
	var n int
	var err error

	t.Run("Case1", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 255)
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "t\r\n.\r")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 1, n)
		require.Equal(t, []byte("t"), buf[:n])

		// reader finishes because of full crlf.crlf
		write(writer, "\n")
		n, err = r.Read(buf)
		require.Error(t, io.EOF, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte("\r\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})

	t.Run("Case2", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 1) // smallest buffer possible
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "t\r\n.\r")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 1, n)
		require.Equal(t, []byte("t"), buf[:n])

		// reader can only return \r as buffer is too small
		write(writer, "\n")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 1, n)
		require.Equal(t, []byte("\r"), buf[:n])

		// reader returns \n and eol
		write(writer, "\n")
		n, err = r.Read(buf)
		require.Error(t, io.EOF, err)
		require.Equal(t, 1, n)
		require.Equal(t, []byte("\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})

	t.Run("Case3", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 255)
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "testtest\r\n.")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 8, n)
		require.Equal(t, []byte("testtest"), buf[:n])

		// reader returns ending
		write(writer, "\r\n")
		n, err = r.Read(buf)
		require.Error(t, io.EOF, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte("\r\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})

	t.Run("Case4", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 255)
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "testtest\r\n.")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 8, n)
		require.Equal(t, []byte("testtest"), buf[:n])

		// close writer
		_ = writer.Close()

		n, err = r.Read(buf)
		require.Error(t, io.ErrUnexpectedEOF, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte("\r\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})

	t.Run("Case4RandomError", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 255)
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "testtest\r\n.")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 8, n)
		require.Equal(t, []byte("testtest"), buf[:n])

		// close writer
		serr := errors.New("test")
		_ = writer.CloseWithError(serr)

		n, err = r.Read(buf)
		require.Error(t, serr, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte("\r\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})

	t.Run("Case5", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 255)
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "testtest\r\n")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 8, n)
		require.Equal(t, []byte("testtest"), buf[:n])

		// close writer
		_ = writer.Close()

		n, err = r.Read(buf)
		require.Error(t, io.ErrUnexpectedEOF, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte("\r\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})

	t.Run("Case5RandomError", func(t *testing.T) {
		reader, writer := io.Pipe()
		buf := make([]byte, 255)
		bufio := bufio.NewReader(reader)
		r := textsmtp.NewDotReader(bufio, 0)

		// only t is read
		write(writer, "testtest\r\n")
		n, err = r.Read(buf)
		require.NoError(t, err)
		require.Equal(t, 8, n)
		require.Equal(t, []byte("testtest"), buf[:n])

		// close writer
		serr := errors.New("test")
		_ = writer.CloseWithError(serr)

		n, err = r.Read(buf)
		require.Error(t, serr, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte("\r\n"), buf[:n])

		// buffer must be empty
		require.Equal(t, 0, bufio.Buffered())
	})
}

// TestDotReaderNoBlockAfterEnd verifies that once the end marker has been
// consumed, a further Read returns io.EOF immediately and does not block
// waiting for more bytes on a still open connection. Without this a caller
// draining the reader would hang, wedging the server goroutine (denial of
// service).
func TestDotReaderNoBlockAfterEnd(t *testing.T) {
	pr, pw := io.Pipe()
	// Write the full message but keep the pipe open, as a live connection
	// waiting for the next command would be.
	go func() { _, _ = pw.Write([]byte("hi\r\n.\r\n")) }()

	r := textsmtp.NewDotReader(bufio.NewReader(pr), 0)

	// Consume the message up to the end marker.
	body, err := io.ReadAll(io.LimitReader(r, 4))
	require.NoError(t, err)
	require.Equal(t, []byte("hi\r\n"), body)

	done := make(chan error, 1)
	go func() {
		b := make([]byte, 8)
		_, e := r.Read(b)
		done <- e
	}()

	select {
	case e := <-done:
		require.Equal(t, io.EOF, e)
	case <-time.After(5 * time.Second):
		t.Fatal("Read blocked after end marker (denial of service)")
	}
}
