// Package deadline sets connection deadlines from timeouts.
// A zero or negative timeout always means no deadline.
package deadline

import (
	"net"
	"time"
)

// Set sets the read and write deadline of conn to now plus timeout.
func Set(conn net.Conn, timeout time.Duration) {
	_ = conn.SetDeadline(at(timeout))
}

// SetRead sets the read deadline of conn to now plus timeout.
func SetRead(conn net.Conn, timeout time.Duration) {
	_ = conn.SetReadDeadline(at(timeout))
}

// SetWrite sets the write deadline of conn to now plus timeout.
func SetWrite(conn net.Conn, timeout time.Duration) {
	_ = conn.SetWriteDeadline(at(timeout))
}

func at(timeout time.Duration) time.Time {
	if timeout <= 0 {
		return time.Time{}
	}
	return time.Now().Add(timeout)
}
