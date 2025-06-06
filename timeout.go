package smtp

import (
	"net"
	"time"
)

func Timeout(conn net.Conn, duration time.Duration) func() {
	_ = conn.SetDeadline(time.Now().Add(duration))
	return func() {
		_ = conn.SetDeadline(time.Time{})
	}
}
