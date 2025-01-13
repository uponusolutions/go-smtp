package server

import (
	"context"
	"io"
	"log/slog"

	"github.com/emersion/go-sasl"
	"github.com/uponusolutions/go-smtp"
)

// A SMTP server backend.
type Backend interface {
	NewSession(ctx context.Context, c *Conn) (context.Context, Session, error)
}

// BackendFunc is an adapter to allow the use of an ordinary function as a
// Backend.
type BackendFunc func(ctx context.Context, c *Conn) (context.Context, Session, error)

// NewSession calls f(c).
func (f BackendFunc) NewSession(ctx context.Context, c *Conn) (context.Context, Session, error) {
	return f(ctx, c)
}

// Session is used by servers to respond to an SMTP client.
//
// The methods are called when the remote client issues the matching command.
type Session interface {
	// Discard currently processed message.
	Reset(ctx context.Context, upgrade bool) (context.Context, error)

	// Free all resources associated with session.
	Close(ctx context.Context) error

	// Returns logger to use when an error occurs inside a session.
	Logger(ctx context.Context) *slog.Logger

	// Set return path for currently processed message.
	Mail(ctx context.Context, from string, opts *smtp.MailOptions) error
	// Add recipient for currently processed message.
	Rcpt(ctx context.Context, to string, opts *smtp.RcptOptions) error
	// Set currently processed message contents and send it.
	//
	// r must be consumed before Data returns.
	Data(ctx context.Context, r func() io.Reader) (string, error)

	AuthMechanisms(ctx context.Context) []string
	Auth(ctx context.Context, mech string) (sasl.Server, error)
}
