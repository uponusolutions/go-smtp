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
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
	"github.com/uponusolutions/go-smtp/tester"
)

type dotReader struct {
	r     *bufio.Reader
	state int

	limited bool
	n       int64 // Maximum bytes remaining
}

// NewDotReader creates a new dot reader.
func NewDotReader(reader *bufio.Reader, maxMessageBytes int64) io.Reader {
	dr := &dotReader{
		r: reader,
	}

	if maxMessageBytes > 0 {
		dr.limited = true
		dr.n = maxMessageBytes
	}

	return dr
}

// Read reads in some more bytes.
func (r *dotReader) Read(b []byte) (n int, err error) {
	if r.limited {
		if r.n <= 0 {
			return 0, smtp.ErrDataTooLarge
		}
		if int64(len(b)) > r.n {
			b = b[0:r.n]
		}
	}

	// Code below is taken from net/textproto with only one modification to
	// not rewrite CRLF -> LF.

	// Run data through a simple state machine to
	// elide leading dots and detect End-of-Data (<CR><LF>.<CR><LF>) line.
	const (
		stateBeginLine = iota // beginning of line; initial state; must be zero
		stateDot              // read . at beginning of line
		stateDotCR            // read .\r at beginning of line
		stateCR               // read \r (possibly at end of line)
		stateData             // reading data in middle of line
		stateEOF              // reached .\r\n end marker line
	)
	for n < len(b) && r.state != stateEOF {
		var c byte
		c, err = r.r.ReadByte()
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			break
		}
		switch r.state {
		case stateBeginLine:
			if c == '.' {
				r.state = stateDot
				continue
			}
			if c == '\r' {
				r.state = stateCR
				break
			}
			r.state = stateData
		case stateDot:
			if c == '\r' {
				r.state = stateDotCR
				continue
			}
			r.state = stateData
		case stateDotCR:
			if c == '\n' {
				r.state = stateEOF
				continue
			}
			r.state = stateData
		case stateCR:
			if c == '\n' {
				r.state = stateBeginLine
				break
			}
			r.state = stateData
		case stateData:
			if c == '\r' {
				r.state = stateCR
			}
		}
		b[n] = c
		n++
	}
	if err == nil && r.state == stateEOF {
		err = io.EOF
	}

	if r.limited {
		r.n -= int64(n)
	}
	return n, err
}

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

	t.Run("CompareTestByte", func(t *testing.T) {
		input := "dotlines\r\n.foo\r\n..bar\n...baz\nquux\r\n\r\n.\r\nanot.her\n"
		readerOld := bufio.NewReader(strings.NewReader(input))
		reader := bufio.NewReader(strings.NewReader(input))

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

			if errOld == io.EOF {
				break
			}
		}

		dotReaderOld = NewDotReader(readerOld, 0)
		dotReader = textsmtp.NewDotReader(reader, 0)

		i = 0

		for {
			nOld, errOld := dotReaderOld.Read(bufOld)
			n, err := dotReader.Read(buf)

			require.Equal(t, bufOld, buf, i)
			require.Equal(t, nOld, n, i)

			if errOld != nil && err != io.ErrUnexpectedEOF {
				require.Equal(t, errOld, err, i)
			}

			i++

			if errOld == io.ErrUnexpectedEOF {
				break
			}
		}
	})

	t.Run("Decode", func(t *testing.T) {
		buf := bufio.NewReader(strings.NewReader("dotlines\r\n.foo\r\n..bar\n...baz\nquux\r\n\r\n.\r\nanot.her\n"))
		r := textsmtp.NewDotReader(buf, 0)
		b, err := io.ReadAll(r)
		require.NoError(t, err)
		require.Equal(t, "dotlines\r\nfoo\r\n.bar\n...baz\nquux\r\n\r\n", b)

		r = textsmtp.NewDotReader(buf, 0)
		b, err = io.ReadAll(r)
		require.NoError(t, err)
		require.Equal(t, "anot.her\n", b)
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
}
