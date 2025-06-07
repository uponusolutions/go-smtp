package smtp

import (
	"net"
	"time"
)

// Timeout sets a timeout by deadline to the connection and relieves it when returning func is used.
func Timeout(conn net.Conn, duration time.Duration) func() {
	_ = conn.SetDeadline(time.Now().Add(duration))
	return func() {
		_ = conn.SetDeadline(time.Time{})
	}
}
