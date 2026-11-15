package smtp

import (
	"net"
	"time"
)

// Timeout sets a timeout by deadline to the connection and relieves it when returning func is used.
// A zero or negative duration means no deadline.
func Timeout(conn net.Conn, duration time.Duration) func() {
	var deadline time.Time
	if duration > 0 {
		deadline = time.Now().Add(duration)
	}
	_ = conn.SetDeadline(deadline)
	return func() {
		_ = conn.SetDeadline(time.Time{})
	}
}
