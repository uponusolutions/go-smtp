package server

import (
	"io"

	"github.com/emersion/go-sasl"
	"github.com/uponusolutions/go-smtp"
)

// A SMTP server backend.
type Backend interface {
	NewSession(c *Conn) (Session, error)
}

// BackendFunc is an adapter to allow the use of an ordinary function as a
// Backend.
type BackendFunc func(c *Conn) (Session, error)

var _ Backend = (BackendFunc)(nil)

// NewSession calls f(c).
func (f BackendFunc) NewSession(c *Conn) (Session, error) {
	return f(c)
}

// Session is used by servers to respond to an SMTP client.
//
// The methods are called when the remote client issues the matching command.
type Session interface {
	// Discard currently processed message.
	Reset()

	// Free all resources associated with session.
	Logout() error

	// Set return path for currently processed message.
	Mail(from string, opts *smtp.MailOptions) error
	// Add recipient for currently processed message.
	Rcpt(to string, opts *smtp.RcptOptions) error
	// Set currently processed message contents and send it.
	//
	// r must be consumed before Data returns.
	Data(r func() io.Reader) error
}

// StatusCollector allows a backend to provide per-recipient status
// information.
type StatusCollector interface {
	SetStatus(rcptTo string, err error)
}

// AuthSession is an add-on interface for Session. It provides support for the
// AUTH extension.
type AuthSession interface {
	Session

	AuthMechanisms() []string
	Auth(mech string) (sasl.Server, error)
}
