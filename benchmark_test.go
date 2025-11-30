package smtp_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"embed"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/client"
	"github.com/uponusolutions/go-smtp/server"
	"github.com/uponusolutions/go-smtp/tester"
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
	return []string{"PLAIN"}
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

func sendMailCon(c *client.Client, data []byte, simplereader bool) error {
	from := "alice@internal.com"
	recipients := []string{"bob@external.com", "tim@external.com"}

	var in io.Reader

	if simplereader {
		in = tester.NewBuffer(data)
	} else {
		in = bytes.NewBuffer(data)
	}

	_, _, err := c.SendMail(context.Background(), from, recipients, in)
	return err
}

func sendMail(addr string, data []byte, simplereader bool) error {
	c := client.New(
		client.WithServerAddresses(addr),
		client.WithSecurity(client.SecurityPlain),
		client.WithMailOptions(client.MailOptions{Size: int64(len(data))}),
	)

	err := c.Connect(context.Background())
	if err != nil {
		return nil
	}

	err = sendMailCon(c, data, simplereader)
	if err != nil {
		return nil
	}

	return c.Quit()
}

type testcase struct {
	eml  []byte
	name string
}

func Benchmark(b *testing.B) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
	slog.SetDefault(l)

	smallEml, err := embedFSTestadata.ReadFile("testdata/small.eml")
	require.NoError(b, err)

	largeEml, err := embedFSTestadata.ReadFile("testdata/large.eml")
	require.NoError(b, err)

	for _, t := range []testcase{
		{
			eml:  smallEml,
			name: "Small",
		},
		{
			eml:  largeEml,
			name: "Large",
		},
	} {
		s1(b, t)
		s2(b, t)
	}

	// require.EqualValues(b, be1.messages, be2.messages)
}

func s1(b *testing.B, t testcase) {
	_, s1, addr1, err := testServer(nil, server.WithEnableCHUNKING(true))
	require.NoError(b, err)

	b.Run(t.name+"WithChunking", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(int64(len(t.eml)))
		}
		for b.Loop() {
			_ = sendMail(addr1, t.eml, false)
		}
	})

	b.Run(t.name+"WithChunkingSameConnection", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(int64(len(t.eml)))
		}
		c := client.New(
			client.WithServerAddresses(addr1),
			client.WithSecurity(client.SecurityPlain),
			client.WithMailOptions(client.MailOptions{Size: int64(len(t.eml))}),
		)
		require.NotNil(b, c)
		require.NoError(b, c.Connect(context.Background()))

		for b.Loop() {
			_ = sendMailCon(c, t.eml, false)
		}

		err = c.Quit()
		require.NoError(b, err)
	})

	b.Run(t.name+"WithChunkingSameConnectionSimpleReader", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(int64(len(t.eml)))
		}
		c := client.New(
			client.WithServerAddresses(addr1),
			client.WithSecurity(client.SecurityPlain),
			client.WithMailOptions(client.MailOptions{Size: int64(len(t.eml))}),
		)
		require.NotNil(b, c)
		require.NoError(b, c.Connect(context.Background()))

		for b.Loop() {
			_ = sendMailCon(c, t.eml, true)
		}

		err = c.Quit()
		require.NoError(b, err)
	})

	require.NoError(b, s1.Close())
}

func s2(b *testing.B, t testcase) {
	_, s2, addr2, err := testServer(nil, server.WithEnableCHUNKING(false))
	require.NoError(b, err)

	b.Run(t.name+"WithoutChunking", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(int64(len(t.eml)))
		}
		for b.Loop() {
			_ = sendMail(addr2, t.eml, false)
		}
	})

	b.Run(t.name+"WithoutChunkingSameConnection", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(int64(len(t.eml)))
		}
		c := client.New(
			client.WithServerAddresses(addr2),
			client.WithSecurity(client.SecurityPlain),
			client.WithMailOptions(client.MailOptions{Size: int64(len(t.eml))}),
		)
		require.NotNil(b, c)

		require.NoError(b, c.Connect(context.Background()))

		for b.Loop() {
			_ = sendMailCon(c, t.eml, false)
		}

		err = c.Quit()
		require.NoError(b, err)
	})

	b.Run(t.name+"WithoutChunkingSameConnectionSimpleReader", func(b *testing.B) {
		if os.Getenv("SETBYTES") == "" {
			b.SetBytes(int64(len(t.eml)))
		}
		c := client.New(
			client.WithServerAddresses(addr2),
			client.WithSecurity(client.SecurityPlain),
			client.WithMailOptions(client.MailOptions{Size: int64(len(t.eml))}),
		)
		require.NotNil(b, c)

		require.NoError(b, c.Connect(context.Background()))

		for b.Loop() {
			_ = sendMailCon(c, t.eml, true)
		}

		err = c.Quit()
		require.NoError(b, err)
	})

	require.NoError(b, s2.Close())
}
