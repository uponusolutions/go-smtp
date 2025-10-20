package benchmark

import (
	"bytes"
	"context"
	"crypto/tls"
	"embed"
	"io"
	"log/slog"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/client"
	"github.com/uponusolutions/go-smtp/server"
)

//go:embed testdata/*
var embedFSTestadata embed.FS

type message struct {
	From     string
	To       []string
	RcptOpts []*smtp.RcptOptions
	Data     []byte
	Opts     *smtp.MailOptions
}

type backend struct{}

func (be *backend) NewSession(ctx context.Context, _ *server.Conn) (context.Context, server.Session, error) {
	return ctx, &session{backend: be}, nil
}

type session struct {
	backend *backend

	msg *message
}

func (*session) Logger(_ context.Context) *slog.Logger {
	return nil
}

func (*session) AuthMechanisms(_ context.Context) []string {
	return nil
}

func (*session) Auth(_ context.Context, _ string) (sasl.Server, error) {
	return nil, nil
}

func (s *session) Reset(ctx context.Context, _ bool) (context.Context, error) {
	s.msg = &message{}
	return ctx, nil
}

func (*session) Close(_ context.Context, _ error) {
}

func (*session) STARTTLS(_ context.Context, tls *tls.Config) (*tls.Config, error) {
	return tls, nil
}

func (*session) Verify(_ context.Context, _ string, _ *smtp.VrfyOptions) error {
	return nil
}

func (s *session) Mail(ctx context.Context, from string, opts *smtp.MailOptions) error {
	_, _ = s.Reset(ctx, false)
	s.msg.From = from
	s.msg.Opts = opts
	return nil
}

func (s *session) Rcpt(_ context.Context, to string, opts *smtp.RcptOptions) error {
	s.msg.To = append(s.msg.To, to)
	s.msg.RcptOpts = append(s.msg.RcptOpts, opts)
	return nil
}

func (s *session) Data(_ context.Context, r func() io.Reader) (string, error) {
	b, err := io.ReadAll(r())
	if err != nil {
		return "", err
	}
	s.msg.Data = b

	// s.backend.messages = append(s.backend.messages, s.msg)

	return "", nil
}

func testServer(bei *backend, opts ...server.Option) (be *backend, s *server.Server, port string, err error) {
	if bei == nil {
		be = new(backend)
	} else {
		be = bei
	}

	curOpts := []server.Option{
		server.WithAddr("127.0.0.1:0"),
		server.WithBackend(be),
		server.WithMaxLineLength(2000),
		server.WithHostname("localhost"),
	}

	curOpts = append(curOpts, opts...)

	s = server.New(
		curOpts...,
	)

	ctx := context.Background()

	l, err := s.Listen()
	if err != nil {
		return nil, nil, "", err
	}

	go func() {
		// nolint: revive
		_ = s.Serve(ctx, l)
	}()

	return be, s, l.Addr().String(), nil
}

func sendMailCon(c *client.Client, data []byte) error {
	from := "alice@internal.com"
	recipients := []string{"bob@external.com", "tim@external.com"}

	in := bytes.NewReader(data)

	_, _, err := c.SendMail(from, recipients, in)
	return err
}

func sendMail(addr string, data []byte) error {
	c := client.New(client.WithServerAddresses(addr), client.WithSecurity(client.SecurityPlain))

	err := c.Connect(context.Background())
	if err != nil {
		return nil
	}

	err = sendMailCon(c, data)
	if err != nil {
		return nil
	}

	return c.Quit()
}
