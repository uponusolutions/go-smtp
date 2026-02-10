package tester

import (
	"io"
)

// A Buffer is a variable-sized buffer of bytes with [Buffer.Read] and [Buffer.Write] methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	buf []byte // contents are the bytes buf[off : len(buf)]
	off int    // read at &buf[off], write at &buf[len(buf)]
}

// empty reports whether the unread portion of the buffer is empty.
func (b *Buffer) empty() bool { return len(b.buf) <= b.off }

// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as [Buffer.Truncate](0).
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is [io.EOF] (unless len(p) is zero);
// otherwise it is nil.
func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.empty() {
		// Buffer is empty, reset to recover space.
		b.Reset()
		if len(p) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	return n, nil
}

// NewBuffer creates and initializes a new [Buffer] using buf as its
// initial contents. The new [Buffer] takes ownership of buf, and the
// caller should not use buf after this call. NewBuffer is intended to
// prepare a [Buffer] to read existing data.
//
// In most cases, new([Buffer]) (or just declaring a [Buffer] variable) is
// sufficient to initialize a [Buffer].
//
// Buffer just implements a reader without any extra functionality.
func NewBuffer(buf []byte) io.Reader { return &Buffer{buf: buf} }
