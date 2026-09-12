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
	"github.com/uponusolutions/go-smtp/mailer"
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

func newMailer(addr string, opts []client.Option) *mailer.Mailer {
	return mailer.New(
		mailer.WithServerAddresses(addr),
		mailer.WithSecurity(mailer.SecurityPlain),
		mailer.WithBasic(opts...),
	)
}

func sendMailCon(c *mailer.Mailer, in io.Reader) error {
	from := "alice@internal.com"
	recipients := []string{"bob@external.com", "tim@external.com"}

	_, _, err := c.Send(context.Background(), from, recipients, in)
	return err
}

type testcase struct {
	eml  []byte
	name string
}

// serverConfig is one server configuration to benchmark against.
type serverConfig struct {
	name string
	opts []server.Option
}

var serverConfigs = []serverConfig{
	{
		name: "",
		opts: []server.Option{server.WithEnableCHUNKING(false)},
	},
	{
		name: "Chunking",
		opts: []server.Option{server.WithEnableCHUNKING(true)},
	},
}

// clientConfig is one client configuration, applied to every mailer.
type clientConfig struct {
	name string
	opts []client.Option
}

var clientConfigs = []clientConfig{
	{
		name: "",
	},
	{
		name: "Pipelining",
		opts: []client.Option{client.WithPipelining(true)},
	},
}

// readerConfig is one way of handing the message body to the mailer.
type readerConfig struct {
	name string
	new  func(data []byte) io.Reader
}

var readerConfigs = []readerConfig{
	{
		name: "",
		new:  func(data []byte) io.Reader { return tester.NewBuffer(data) },
	},
	{
		name: "Bytes",
		new:  func(data []byte) io.Reader { return bytes.NewBuffer(data) },
	},
}

// connectionConfig is one connection lifecycle, driving the benchmark loop and
// calling send once per iteration.
type connectionConfig struct {
	name string
	run  func(b *testing.B, addr string, opts []client.Option, send func(c *mailer.Mailer) error)
}

var connectionConfigs = []connectionConfig{
	{
		name: "",
		run: func(b *testing.B, addr string, opts []client.Option, send func(c *mailer.Mailer) error) {
			for b.Loop() {
				c := newMailer(addr, opts)

				if err := c.Connect(context.Background()); err != nil {
					continue
				}

				if err := send(c); err != nil {
					continue
				}

				_ = c.Disconnect()
			}
		},
	},
	{
		name: "Reuse",
		run: func(b *testing.B, addr string, opts []client.Option, send func(c *mailer.Mailer) error) {
			c := newMailer(addr, opts)
			require.NotNil(b, c)
			require.NoError(b, c.Connect(context.Background()))

			for b.Loop() {
				_ = send(c)
			}

			require.NoError(b, c.Disconnect())
		},
	},
}

func setBytes(b *testing.B, eml []byte) {
	b.Helper()

	if os.Getenv("SETBYTES") == "" {
		b.SetBytes(int64(len(eml)))
	}
}

// benchmarkServer runs every client, connection and reader config against one
// server config.
func benchmarkServer(b *testing.B, t testcase, sc serverConfig) {
	b.Helper()

	_, s, addr, err := testServer(nil, sc.opts...)
	require.NoError(b, err)

	defer func() {
		require.NoError(b, s.Close())
	}()

	for _, cc := range clientConfigs {
		for _, conn := range connectionConfigs {
			for _, rc := range readerConfigs {
				b.Run(t.name+sc.name+cc.name+conn.name+rc.name, func(b *testing.B) {
					setBytes(b, t.eml)

					conn.run(b, addr, cc.opts, func(c *mailer.Mailer) error {
						return sendMailCon(c, rc.new(t.eml))
					})
				})
			}
		}
	}
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
		{eml: smallEml, name: "Small"},
		{eml: largeEml, name: "Large"},
	} {
		for _, sc := range serverConfigs {
			benchmarkServer(b, t, sc)
		}
	}
}
