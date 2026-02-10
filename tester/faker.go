package tester

import (
	"bytes"
	"io"
	"net"
	"strings"
	"time"
)

// FakeConn fakes a conn for testing.
type FakeConn struct {
	io.ReadWriter
	RemoteAddrReturn net.Addr
}

// NewFakeConnStream creates a new FakeConn with a stream as a input.
func NewFakeConnStream(in io.Reader, out *bytes.Buffer) *FakeConn {
	rw := struct {
		io.Reader
		io.Writer
	}{
		Reader: in,
		Writer: out,
	}

	return &FakeConn{
		ReadWriter: rw,
	}
}

// NewFakeConn creates a new FakeConn with a string as a input.
func NewFakeConn(in string, out *bytes.Buffer) *FakeConn {
	rw := struct {
		io.Reader
		io.Writer
	}{
		Reader: strings.NewReader(in),
		Writer: out,
	}

	return &FakeConn{
		ReadWriter: rw,
	}
}

// Close always returns nil.
func (FakeConn) Close() error { return nil }

// LocalAddr always returns nil.
func (FakeConn) LocalAddr() net.Addr { return nil }

// RemoteAddr always returns RemoteAddrReturn.
func (f FakeConn) RemoteAddr() net.Addr { return f.RemoteAddrReturn }

// SetDeadline always returns nil and does nothing.
func (FakeConn) SetDeadline(time.Time) error { return nil }

// SetReadDeadline always returns nil and does nothing.
func (FakeConn) SetReadDeadline(time.Time) error { return nil }

// SetWriteDeadline always returns nil and does nothing.
func (FakeConn) SetWriteDeadline(time.Time) error { return nil }
