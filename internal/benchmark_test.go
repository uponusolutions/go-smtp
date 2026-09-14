package internal

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"math/rand/v2"
	upstreamtextproto "net/textproto"
	"os"
	"testing"

	"github.com/uponusolutions/go-smtp/internal/smtpreader"
	"github.com/uponusolutions/go-smtp/internal/smtpwriter"
	"github.com/uponusolutions/go-smtp/tester"
	"github.com/uponusolutions/go-smtp/tester/upstream"
)

// payloadSize is used for both directions so read and write throughput are
// directly comparable. 256 MiB per iteration, as the writer benchmark used
// before, costs a few hundred milliseconds per op and leaves the default
// benchtime with a handful of samples — far too few to say anything.
const payloadSize = 4 << 20

// payload is one message body in both its plain and its dot-encoded form.
type payload struct {
	name string
	raw  []byte // plain body, input of the write direction
	enc  []byte // dot-encoded body, input of the read direction
	dec  int64  // number of bytes decoding enc must yield
}

// payloadConfig is one shape of body to benchmark against.
type payloadConfig struct {
	name string
	new  func(size int) []byte
}

var payloadConfigs = []payloadConfig{
	{
		name: "Binary",
		new:  binaryPayload,
	},
	{
		name: "Text",
		new:  textPayload,
	},
}

// binaryPayload returns pseudo-random bytes. This is the codec's fast path:
// random data holds roughly one line break every 256 bytes and no leading dots,
// so neither stuffing nor unstuffing is ever exercised.
// nolint:gosec
func binaryPayload(size int) []byte {
	rng := rand.New(rand.NewPCG(1, 0x5eed))

	p := make([]byte, size)
	for i := 0; i+8 <= len(p); i += 8 {
		binary.LittleEndian.PutUint64(p[i:], rng.Uint64())
	}

	for i := len(p) - len(p)%8; i < len(p); i++ {
		p[i] = byte(rng.Uint32())
	}

	return p
}

// textPayload returns CRLF-terminated ASCII lines, every eighth of which starts
// with a dot. That is what a real message looks like to the codec, and it is
// the only variant that makes the encoder stuff and the decoder unstuff.
// nolint:gosec
func textPayload(size int) []byte {
	rng := rand.New(rand.NewPCG(2, 0x5eed))

	buf := bytes.NewBuffer(make([]byte, 0, size+80))
	line := make([]byte, 0, 80)

	for i := 0; buf.Len() < size; i++ {
		line = line[:0]
		if i%8 == 0 {
			line = append(line, '.')
		}

		for len(line) < 76 {
			line = append(line, byte('a'+rng.IntN(26)))
		}

		line = append(line, '\r', '\n')
		buf.Write(line)
	}

	return buf.Bytes()
}

// encode dot-encodes raw with the upstream implementation, which serves as the
// reference for both directions.
func encode(b *testing.B, raw []byte) []byte {
	b.Helper()

	var buf bytes.Buffer

	bw := bufio.NewWriter(&buf)
	w := upstreamtextproto.NewWriter(bw).DotWriter()

	if _, err := w.Write(raw); err != nil {
		b.Fatalf("encode payload: %v", err)
	}

	// Close writes the terminating dot line and flushes bw. Without it the
	// encoded body is truncated to a multiple of the bufio buffer size and
	// carries no terminator, so every decoder would stop at
	// io.ErrUnexpectedEOF instead of at the end of the message.
	if err := w.Close(); err != nil {
		b.Fatalf("close encoder: %v", err)
	}

	return buf.Bytes()
}

// decodedLen returns how many bytes enc decodes to, so the read direction can
// verify it consumed the whole message rather than stopping early.
func decodedLen(b *testing.B, enc []byte) int64 {
	b.Helper()

	r := upstream.NewDotReader(bufio.NewReader(bytes.NewReader(enc)), 0)

	n, err := io.Copy(io.Discard, r)
	if err != nil {
		b.Fatalf("decode payload: %v", err)
	}

	return n
}

// implConfig is one dot codec implementation.
type implConfig struct {
	name      string
	dotReader func(br *bufio.Reader) io.Reader
	dotWriter func(bw *bufio.Writer) io.WriteCloser
}

var implConfigs = []implConfig{
	{
		name: "Upstream",
		dotReader: func(br *bufio.Reader) io.Reader {
			return upstream.NewDotReader(br, 0)
		},
		dotWriter: func(bw *bufio.Writer) io.WriteCloser {
			return upstreamtextproto.NewWriter(bw).DotWriter()
		},
	},
	{
		name: "Fork",
		dotReader: func(br *bufio.Reader) io.Reader {
			return smtpreader.NewDot(br, 0)
		},
		dotWriter: func(bw *bufio.Writer) io.WriteCloser {
			return smtpwriter.NewDot(bw)
		},
	},
}

// sourceConfig is one way of handing the body to the codec.
//
// bytes.Reader implements io.WriterTo, so io.Copy hands the whole body over in
// a single call. tester.Buffer does not, so io.Copy falls back to its 32 KiB
// staging buffer. The two configs therefore exercise noticeably different code
// paths, which is the point of keeping both.
type sourceConfig struct {
	name string
	new  func(data []byte) io.Reader
}

var sourceConfigs = []sourceConfig{
	{
		name: "MinimalReader",
		new:  func(data []byte) io.Reader { return bytes.NewReader(data) },
	},
	{
		name: "BufferReader",
		new:  func(data []byte) io.Reader { return tester.NewBuffer(data) },
	},
}

// directionConfig drives the benchmark loop for one direction of the codec.
//
// Every error aborts the benchmark, and the number of bytes moved is checked
// against the payload. A codec that stops early would otherwise report a very
// good — and very wrong — result.
type directionConfig struct {
	name string
	run  func(b *testing.B, impl implConfig, src sourceConfig, p payload)
}

var directionConfigs = []directionConfig{
	{
		name: "Read",
		run: func(b *testing.B, impl implConfig, src sourceConfig, p payload) {
			setBytes(b, len(p.enc))

			for b.Loop() {
				r := impl.dotReader(bufio.NewReader(src.new(p.enc)))

				n, err := io.Copy(io.Discard, r)
				if err != nil {
					b.Fatalf("read: %v", err)
				}

				if n != p.dec {
					b.Fatalf("read %d bytes, want %d", n, p.dec)
				}
			}
		},
	},
	{
		name: "Write",
		run: func(b *testing.B, impl implConfig, src sourceConfig, p payload) {
			setBytes(b, len(p.raw))

			for b.Loop() {
				bw := bufio.NewWriter(io.Discard)
				w := impl.dotWriter(bw)

				n, err := io.Copy(w, src.new(p.raw))
				if err != nil {
					b.Fatalf("write: %v", err)
				}

				if n != int64(len(p.raw)) {
					b.Fatalf("wrote %d bytes, want %d", n, len(p.raw))
				}

				// Terminator plus flush belong to the work being measured;
				// leaving them out dropped the tail of every message on the
				// floor and hid the cost of Close entirely.
				if err := w.Close(); err != nil {
					b.Fatalf("close: %v", err)
				}

				if err := bw.Flush(); err != nil {
					b.Fatalf("flush: %v", err)
				}
			}
		},
	},
}

// setBytes reports the size of the input so the benchmark prints throughput in
// MB/s. Set NOSETBYTES=1 to suppress it, e.g. when comparing ns/op or
// allocations across payload shapes, where a MB/s column is only noise.
func setBytes(b *testing.B, n int) {
	b.Helper()

	if os.Getenv("NOSETBYTES") != "" {
		return
	}

	b.SetBytes(int64(n))
}

// BenchmarkDot benchmarks both directions of the dot codec against every
// payload, source and implementation. The configs are nested as sub-benchmarks,
// so names are paths, e.g. BenchmarkDot/Read/Text/BytesReader/Fork, and can be
// filtered with -bench 'Dot/Read/Text/.*/Fork'.
//
// Upstream and Fork sit innermost so the two implementations appear on adjacent
// lines under otherwise identical conditions.
func BenchmarkDot(b *testing.B) {
	payloads := make([]payload, 0, len(payloadConfigs))

	for _, pc := range payloadConfigs {
		raw := pc.new(payloadSize)
		enc := encode(b, raw)

		payloads = append(payloads, payload{
			name: pc.name,
			raw:  raw,
			enc:  enc,
			dec:  decodedLen(b, enc),
		})
	}

	for _, dc := range directionConfigs {
		b.Run(dc.name, func(b *testing.B) {
			for _, p := range payloads {
				b.Run(p.name, func(b *testing.B) {
					for _, sc := range sourceConfigs {
						b.Run(sc.name, func(b *testing.B) {
							for _, ic := range implConfigs {
								b.Run(ic.name, func(b *testing.B) {
									dc.run(b, ic, sc, p)
								})
							}
						})
					}
				})
			}
		})
	}
}
