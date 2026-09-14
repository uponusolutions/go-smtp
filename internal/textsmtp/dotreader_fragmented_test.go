package textsmtp_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

type fragmentReader struct {
	io.Reader
	size int
}

func (r fragmentReader) Read(p []byte) (int, error) {
	if len(p) > r.size {
		p = p[:r.size]
	}
	return r.Reader.Read(p)
}

func TestDataReaderFragmented(t *testing.T) {
	tests := []struct {
		name, body, want string
		err              error
	}{
		{"empty", ".\r\n", "", nil},
		{"suffix-4", "a\r\n.\r", "a\r\n", io.ErrUnexpectedEOF},
		{"suffix-3", "a\r\n", "a\r\n", io.ErrUnexpectedEOF},
		{"suffix-2", "a\r", "a\r", io.ErrUnexpectedEOF},
		{"suffix-1", "a", "a", io.ErrUnexpectedEOF},
		{"lines", "hello\r\nworld\r\n.\r\n", "hello\r\nworld\r\n", nil},
		{"dots", "..one\r\n...two\r\n.\r\n", ".one\r\n..two\r\n", nil},
		{"bare-lf", "first\n.second\r\n.\r\n", "first\n.second\r\n", nil},
		{"bare-cr", "first\rsecond\r\n.\r\n", "first\rsecond\r\n", nil},
		{"dot-cr", ".\rx\r\n.\ry\r\n.\r\n", "\rx\r\n\ry\r\n", nil}, // original "x\r\n" -- why should \r be removed?
		{"loose-cr-lf", "a\rb\nc\rd\r\n.\r\n", "a\rb\nc\rd\r\n", nil},
		{"unterminated", "hello\r\n", "hello\r\n", io.ErrUnexpectedEOF},
		{"split-dot", "hello\r\n.", "hello\r\n", io.ErrUnexpectedEOF},
		{"long-run", strings.Repeat("a", 8193) + "\r\n.\r\n", strings.Repeat("a", 8193) + "\r\n", nil},
	}
	for _, tc := range tests {
		for _, input := range []int{1, 2, 3, 17, 4096} {
			for _, output := range []int{64, 4096} {
				t.Run(fmt.Sprintf("%s/in%d/out%d", tc.name, input, output), func(t *testing.T) {
					wire := tc.body
					if tc.err == nil {
						wire += "NEXT\r\n"
					}
					buffered := bufio.NewReader(fragmentReader{strings.NewReader(wire), input})
					r := textsmtp.NewDotReader(buffered, 0)

					var got bytes.Buffer
					buffer := make([]byte, output)
					_, err := io.CopyBuffer(&writeOnlyBuffer{&got}, r, buffer)
					if err != tc.err || got.String() != tc.want {
						t.Fatalf("got %q, %v; want %q, %v", got.String(), err, tc.want, tc.err)
					}
					if tc.err == nil {
						remaining, err := io.ReadAll(buffered)
						if err != nil || string(remaining) != "NEXT\r\n" {
							t.Fatalf("read beyond DATA: %q, %v", remaining, err)
						}
					}
				})
			}
		}
	}
}

type writeOnlyBuffer struct{ b *bytes.Buffer }

func (w *writeOnlyBuffer) Write(p []byte) (int, error) { return w.b.Write(p) }

func TestDataReaderBulkLimit(t *testing.T) {
	for _, limit := range []int64{1, 63, 4095, 4096, 4097} {
		r := textsmtp.NewDotReader(bufio.NewReader(strings.NewReader(strings.Repeat("a", 8192)+"\r\n.\r\n")), limit)
		data, err := io.ReadAll(r)
		if err != smtp.ErrDataTooLarge || int64(len(data)) != limit {
			t.Fatalf("limit %d: read %d, %v", limit, len(data), err)
		}
	}
}
