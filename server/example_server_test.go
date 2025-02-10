package server_test

import (
	"context"
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
func (bkd *Backend) NewSession(ctx context.Context, c *server.Conn) (context.Context, server.Session, error) {
	return ctx, &Session{}, nil
}

// A Session is returned after successful login.
type Session struct{}

func (s *Session) Logger(ctx context.Context) *slog.Logger {
	return nil
}

// AuthMechanisms returns a slice of available auth mechanisms; only PLAIN is
// supported in this example.
func (s *Session) AuthMechanisms(ctx context.Context) []string {
	return []string{sasl.Plain}
}

// Auth is the handler for supported authenticators.
func (s *Session) Auth(ctx context.Context, mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		if username != "username" || password != "password" {
			return errors.New("Invalid username or password")
		}
		return nil
	}), nil
}

func (s *Session) Mail(ctx context.Context, from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(ctx context.Context, to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(ctx context.Context, r func() io.Reader) (string, error) {
	if b, err := io.ReadAll(r()); err != nil {
		return "", err
	} else {
		log.Println("Data:", string(b))
	}
	return "", nil
}

func (s *Session) Reset(ctx context.Context, _ bool) (context.Context, error) {
	return ctx, nil
}

func (s *Session) Close(ctx context.Context) error {
	return nil
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

	s := server.NewServer(
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
