// Package smtptester implements a simple SMTP server for testing. All
// received mails are saved in a sync.Map with a key:
//
//	From+Recipient1+Recipient2
//
// Mails to the same sender and recipients will overwrite a previous
// received mail, when the recipients slice has the same order as
// in the mail received before.
package tester

import (
	"context"
	"crypto/tls"
	"io"
	"log/slog"
	"sync"
	"time"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/server"
)

// Standard returns a standard SMTP server listening on a random Port.
func Standard(opts ...server.Option) *server.Server {
	defaultOpts := []server.Option{
		server.WithAddr(":0"),
		server.WithReadTimeout(10 * time.Second),
		server.WithWriteTimeout(10 * time.Second),
		server.WithMaxMessageBytes(1024 * 1024),
		server.WithMaxRecipients(100),
		server.WithBackend(NewBackend()),
	}

	return server.New(
		append(defaultOpts, opts...)...,
	)
}

///////////////////////////////////////////////////////////////////////////
// Backend
///////////////////////////////////////////////////////////////////////////

// Backend is the backend for out test server.
// It contains a sync.Map with all mails received.
type Backend struct {
	Mails sync.Map
}

// NewBackend returns a new Backend with an empty (not nil) Mails map.
func NewBackend() *Backend {
	return &Backend{Mails: sync.Map{}}
}

// NewSession returns a new Session.
func (b *Backend) NewSession(ctx context.Context, _ *server.Conn) (context.Context, server.Session, error) {
	return ctx, newSession(b), nil
}

// GetBackend returns the concrete type *Backend from SMTP server.
func GetBackend(s *server.Server) *Backend {
	if s.Backend() == nil {
		return nil
	}

	b, ok := s.Backend().(*Backend)
	if !ok {
		return nil
	}

	return b
}

// Add adds mail to backends map.
func (b *Backend) Add(m *Mail) {
	b.Mails.Store(m.LookupKey(), m)
}

// Load loads mail from 'from' to recipients 'recipients'. The ok
// result indicates whether value was found in the map.
func (b *Backend) Load(from string, recipients []string) (*Mail, bool) {
	i, ok := b.Mails.Load(LookupKey(from, recipients))
	if !ok {
		return nil, ok
	}

	return i.(*Mail), ok //nolint
}

///////////////////////////////////////////////////////////////////////////
// Session
///////////////////////////////////////////////////////////////////////////

// A Session is returned after successful login.
type Session struct {
	backend *Backend
	mail    *Mail
}

func newSession(b *Backend) *Session {
	return &Session{
		backend: b,
		mail:    &Mail{},
	}
}

// Reset implements Reset interface.
func (s *Session) Reset(ctx context.Context, _ bool) (context.Context, error) {
	s.mail = &Mail{}

	return ctx, nil
}

// Close implements the Close interface.
func (s *Session) Close(_ context.Context, _ error) {
	s.mail = &Mail{}
}

// Logger implements the Logger interface.
func (Session) Logger(_ context.Context) *slog.Logger {
	return nil
}

// Verify implements the Verify interface.
func (Session) Verify(_ context.Context, _ string, _ *smtp.VrfyOptions) error {
	return nil
}

// Mail implements the Mail interface.
func (s *Session) Mail(_ context.Context, from string, _ *smtp.MailOptions) error {
	s.mail.From = from

	return nil
}

// Rcpt implements the Rcpt interface.
func (s *Session) Rcpt(_ context.Context, to string, _ *smtp.RcptOptions) error {
	s.mail.Recipients = append(s.mail.Recipients, to)

	return nil
}

// Data implements the Data interface.
func (s *Session) Data(_ context.Context, r func() io.Reader) (string, error) {
	var err error

	if s.mail.Data, err = io.ReadAll(r()); err != nil {
		return "", err
	}

	s.backend.Add(s.mail)

	return "", nil
}

// AuthMechanisms implements the AuthMechanisms interface.
func (Session) AuthMechanisms(_ context.Context) []string {
	return nil
}

// Auth implements the Auth interface.
func (Session) Auth(_ context.Context, _ string) (sasl.Server, error) {
	return nil, nil
}

// STARTTLS implements the STARTTLS interface.
func (Session) STARTTLS(_ context.Context, config *tls.Config) (*tls.Config, error) {
	return config, nil
}
