package textsmtp_test

import (
	"bufio"
	"bytes"
	"io"
	legacy "net/textproto"
	"strings"
	"testing"

	"github.com/uponusolutions/go-smtp/internal/parse"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
	"github.com/uponusolutions/go-smtp/tester"
)

// canonical strips bare carriage returns so the payload only contains
// LF line endings. go-smtp intentionally handles bare CR differently
// from net/textproto (it does not rewrite CRLF), so the differential
// targets below only compare on this bare-CR-free subset where both
// implementations must agree.
func canonical(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("\r"), nil)
}

// readAllChunked reads r to the end with a fixed buffer size and
// returns everything read and the final error.
func readAllChunked(r io.Reader, chunk int) ([]byte, error) {
	var out []byte
	buf := make([]byte, chunk)
	for {
		n, err := r.Read(buf)
		out = append(out, buf[:n]...)
		if err != nil {
			return out, err
		}
	}
}

// FuzzDotWriter compares the dot writer against net/textproto for
// canonical payloads written in arbitrary chunk sizes.
func FuzzDotWriter(f *testing.F) {
	f.Add([]byte("abc\n.def\n..ghi\n.jkl\n."), uint8(4))
	f.Add([]byte(".leading dot\nsecond"), uint8(1))
	f.Add([]byte(""), uint8(1))
	f.Add([]byte(".\n."), uint8(2))

	f.Fuzz(func(t *testing.T, raw []byte, chunk uint8) {
		data := canonical(raw)
		size := int(chunk)%255 + 1

		var expectedBuf bytes.Buffer
		expected := legacy.NewWriter(bufio.NewWriter(&expectedBuf)).DotWriter()

		var actualBuf bytes.Buffer
		actual := textsmtp.NewDotWriter(bufio.NewWriter(&actualBuf))

		for _, w := range []io.WriteCloser{expected, actual} {
			for p := data; len(p) > 0; {
				k := min(size, len(p))
				if _, err := w.Write(p[:k]); err != nil {
					t.Fatalf("write: %v", err)
				}
				p = p[k:]
			}
			if err := w.Close(); err != nil {
				t.Fatalf("close: %v", err)
			}
		}

		if !bytes.Equal(expectedBuf.Bytes(), actualBuf.Bytes()) {
			t.Fatalf("output mismatch for %q:\nexpected %q\nactual   %q", data, expectedBuf.Bytes(), actualBuf.Bytes())
		}
	})
}

// FuzzDotReader encodes a canonical payload with the net/textproto
// writer (producing a valid dot stream) and compares decoding it with
// net/textproto against the optimized dot reader, byte for byte and
// including the read chunk size.
func FuzzDotReader(f *testing.F) {
	f.Add([]byte("dotlines\n.foo\n..bar\nquux\n"), uint8(16))
	f.Add([]byte(""), uint8(1)) // empty message -> ".\r\n" end marker
	f.Add([]byte(".leading"), uint8(1))
	f.Add([]byte("a\nb\nc"), uint8(3))

	f.Fuzz(func(t *testing.T, raw []byte, chunk uint8) {
		data := canonical(raw)
		size := int(chunk)%255 + 1

		// produce a valid dot-encoded stream
		var enc bytes.Buffer
		w := legacy.NewWriter(bufio.NewWriter(&enc)).DotWriter()
		if _, err := w.Write(data); err != nil {
			t.Fatalf("encode write: %v", err)
		}
		if err := w.Close(); err != nil {
			t.Fatalf("encode close: %v", err)
		}
		stream := enc.Bytes()

		expected, expErr := io.ReadAll(legacy.NewReader(bufio.NewReader(bytes.NewReader(stream))).DotReader())
		// net/textproto rewrites CRLF to LF, undo that for comparison
		expected = bytes.ReplaceAll(expected, []byte("\n"), []byte("\r\n"))

		actual, actErr := readAllChunked(textsmtp.NewDotReader(bufio.NewReader(bytes.NewReader(stream)), 0), size)

		if (expErr == nil) != (actErr == io.EOF) {
			t.Fatalf("error mismatch for %q: expected %v, actual %v", data, expErr, actErr)
		}
		if !bytes.Equal(expected, actual) {
			t.Fatalf("output mismatch for %q (stream %q):\nexpected %q\nactual   %q", data, stream, expected, actual)
		}
	})
}

// FuzzDotReaderRobust feeds arbitrary bytes into the dot reader and
// asserts it never panics, always terminates and respects invariants.
func FuzzDotReaderRobust(f *testing.F) {
	f.Add([]byte(".\r\n"), int64(0), uint8(7))
	f.Add([]byte("a\r\n.\rx\r\n.\r\n"), int64(0), uint8(1))
	f.Add([]byte("\r\r\n"), int64(4), uint8(2))
	f.Add([]byte("no end marker"), int64(0), uint8(16))

	f.Fuzz(func(t *testing.T, data []byte, maxBytes int64, chunk uint8) {
		if maxBytes < 0 {
			maxBytes = -maxBytes
		}
		size := int(chunk)%255 + 1

		out, _ := readAllChunked(textsmtp.NewDotReader(bufio.NewReader(bytes.NewReader(data)), maxBytes), size)

		if len(out) > len(data) {
			t.Fatalf("read %d bytes from %d byte input", len(out), len(data))
		}
		if maxBytes > 0 && int64(len(out)) > maxBytes {
			t.Fatalf("read %d bytes, more than the limit %d", len(out), maxBytes)
		}
	})
}

// FuzzBdatReader feeds an arbitrary BDAT command stream into the bdat
// reader. The first input line is the argument of the initial BDAT
// command, the rest is the wire stream. It asserts no panic and that
// the size limit is respected.
func FuzzBdatReader(f *testing.F) {
	f.Add("5 LAST\r\nhello", int64(0), uint8(16))
	f.Add("5\r\nhelloBDAT 0 LAST\r\n", int64(0), uint8(3))
	f.Add("3\r\nabcRSET\r\n", int64(1024), uint8(1))
	f.Add("1\r\naBDAT 2\r\nbcBDAT 0 LAST\r\n", int64(2), uint8(7))

	f.Fuzz(func(t *testing.T, input string, maxBytes int64, chunk uint8) {
		if maxBytes < 0 {
			maxBytes = -maxBytes
		}
		size := int(chunk)%255 + 1

		idx := strings.IndexByte(input, '\n')
		if idx < 0 {
			return
		}
		arg := strings.TrimRight(input[:idx], "\r")
		r := bufio.NewReader(strings.NewReader(input[idx+1:]))

		nextCommand := func() (string, string, error) {
			line, err := r.ReadString('\n')
			if err != nil {
				return "", "", err
			}
			return parse.Cmd(strings.TrimRight(line, "\r\n"))
		}

		br, err := textsmtp.NewBdatReader(arg, maxBytes, r, nextCommand)
		if err != nil {
			return
		}

		out, _ := readAllChunked(br, size)

		if maxBytes > 0 && int64(len(out)) > maxBytes {
			t.Fatalf("read %d bytes, more than the limit %d", len(out), maxBytes)
		}
	})
}

// FuzzReadResponse parses arbitrary bytes as SMTP responses and lines,
// asserting no panic and termination.
func FuzzReadResponse(f *testing.F) {
	f.Add("250 ok\r\n", 250)
	f.Add("250-a\r\n250-b\r\n250 c\r\n", 2)
	f.Add("550-x\r\ngarbage\r\n550 y\r\n", 25)
	f.Add("99\r\n", 0)

	f.Fuzz(func(_ *testing.T, input string, expect int) {
		tp := textsmtp.NewTextproto(tester.NewFakeConn(input, &bytes.Buffer{}), 64, 64, 1000)
		_, _, _ = tp.ReadResponse(((expect % 1000) + 1000) % 1000)

		tp = textsmtp.NewTextproto(tester.NewFakeConn(input, &bytes.Buffer{}), 64, 64, 32)
		for {
			if _, err := tp.ReadLine(); err != nil {
				break
			}
		}
	})
}
