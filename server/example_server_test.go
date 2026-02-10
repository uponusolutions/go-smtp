package server_test

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"log/slog"
	"time"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/server"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// NewSession is called after client greeting (EHLO, HELO).
func (*Backend) NewSession(ctx context.Context, _ *server.Conn) (context.Context, server.Session, error) {
	return ctx, &Session{}, nil
}

// A Session is returned after successful login.
type Session struct {
	auth bool
}

// Logger returns nil.
func (*Session) Logger(_ context.Context) *slog.Logger {
	return nil
}

// AuthMechanisms returns a slice of available auth mechanisms; only PLAIN is
// supported in this example.
func (*Session) AuthMechanisms(_ context.Context) []string {
	return []string{sasl.Plain}
}

// Auth is the handler for supported authenticators.
func (s *Session) Auth(_ context.Context, _ string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(_, username, password string) error {
		if username != "username" || password != "password" {
			return errors.New("Invalid username or password")
		}
		s.auth = true
		return nil
	}), nil
}

func (s *Session) Mail(_ context.Context, from string, _ *smtp.MailOptions) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	log.Println("Mail from:", from)
	return nil
}

func (*Session) Verify(_ context.Context, _ string, _ *smtp.VrfyOptions) error {
	return nil
}

func (s *Session) Rcpt(_ context.Context, to string, _ *smtp.RcptOptions) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	log.Println("Rcpt to:", to)
	return nil
}

func (*Session) STARTTLS(_ context.Context, tls *tls.Config) (*tls.Config, error) {
	return tls, nil
}

func (s *Session) Data(_ context.Context, r func() io.Reader) (string, error) {
	if !s.auth {
		return "", smtp.ErrAuthRequired
	}

	b, err := io.ReadAll(r())
	if err != nil {
		return "", err
	}
	log.Println("Data:", string(b))
	return "", nil
}

func (*Session) Reset(ctx context.Context, _ bool) (context.Context, error) {
	return ctx, nil
}

func (*Session) Close(_ context.Context, _ error) {
}

// ExampleServer runs an example SMTP server.
//
// It can be tested manually with e.g. netcat:
//
//	> netcat -C localhost 1025
//	EHLO localhost
//	AUTH PLAIN
//	AHVzZXJuYW1lAHBhc3N3b3Jk
//	MAIL FROM:<root@nsa.gov>
//	RCPT TO:<root@gchq.gov.uk>
//	DATA
//	Hey <3
//	.
func ExampleServer() {
	be := &Backend{}
	addr := "localhost:1025"

	s := server.New(
		server.WithBackend(be),
		server.WithAddr(addr),
		server.WithHostname("localhost"),
		server.WithWriteTimeout(10*time.Second),
		server.WithReadTimeout(10*time.Second),
		server.WithMaxMessageBytes(1024*1024),
		server.WithMaxRecipients(50),
	)

	log.Println("Starting server at", addr)
	if err := s.ListenAndServe(context.Background()); err != nil {
		log.Fatal(err)
	}
}
