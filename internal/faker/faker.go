package faker

import (
	"bytes"
	"io"
	"net"
	"strings"
	"time"
)

type Conn struct {
	io.ReadWriter
}

func NewConnStream(in io.Reader, out *bytes.Buffer) *Conn {
	rw := struct {
		io.Reader
		io.Writer
	}{
		Reader: in,
		Writer: out,
	}

	return &Conn{
		ReadWriter: rw,
	}
}

func NewConn(in string, out *bytes.Buffer) *Conn {
	rw := struct {
		io.Reader
		io.Writer
	}{
		Reader: strings.NewReader(in),
		Writer: out,
	}

	return &Conn{
		ReadWriter: rw,
	}
}

func (f Conn) Close() error                     { return nil }
func (f Conn) LocalAddr() net.Addr              { return nil }
func (f Conn) RemoteAddr() net.Addr             { return nil }
func (f Conn) SetDeadline(time.Time) error      { return nil }
func (f Conn) SetReadDeadline(time.Time) error  { return nil }
func (f Conn) SetWriteDeadline(time.Time) error { return nil }
