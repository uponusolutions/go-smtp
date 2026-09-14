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
var embedTestdata embed.FS

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

	return "", nil
}

// testServer starts a server on a random loopback port and returns the address
// it is listening on.
func testServer(bei *backend, opts ...server.Option) (be *backend, s *server.Server, addr string, err error) {
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
		name: "NoChunking",
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
		name: "NoPipelining",
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
		name: "MinimalReader",
		new:  func(data []byte) io.Reader { return tester.NewBuffer(data) },
	},
	{
		name: "BufferReader",
		new:  func(data []byte) io.Reader { return bytes.NewBuffer(data) },
	},
}

// connectionConfig is one connection lifecycle, driving the benchmark loop and
// calling send once per iteration.
//
// Every error inside the loop aborts the benchmark. A failing connection would
// otherwise race through the iterations and report an excellent — and
// meaningless — result.
type connectionConfig struct {
	name string
	run  func(b *testing.B, addr string, opts []client.Option, send func(c *mailer.Mailer) error)
}

var connectionConfigs = []connectionConfig{
	{
		name: "CloseConn",
		run: func(b *testing.B, addr string, opts []client.Option, send func(c *mailer.Mailer) error) {
			for b.Loop() {
				c := newMailer(addr, opts)

				if err := c.Connect(context.Background()); err != nil {
					b.Fatalf("connect: %v", err)
				}

				if err := send(c); err != nil {
					b.Fatalf("send: %v", err)
				}

				if err := c.Disconnect(); err != nil {
					b.Fatalf("disconnect: %v", err)
				}
			}
		},
	},
	{
		name: "ReuseConn",
		run: func(b *testing.B, addr string, opts []client.Option, send func(c *mailer.Mailer) error) {
			c := newMailer(addr, opts)
			require.NotNil(b, c)
			require.NoError(b, c.Connect(context.Background()))

			for b.Loop() {
				if err := send(c); err != nil {
					b.Fatalf("send: %v", err)
				}
			}

			require.NoError(b, c.Disconnect())
		},
	},
}

// setBytes reports the message size so the benchmark prints throughput in MB/s.
// Set NOSETBYTES=1 to suppress it, e.g. when comparing ns/op or allocations
// across different message sizes, where a MB/s column is only noise.
func setBytes(b *testing.B, eml []byte) {
	b.Helper()

	if os.Getenv("NOSETBYTES") != "" {
		return
	}

	b.SetBytes(int64(len(eml)))
}

// benchmarkServer runs every client, connection and reader config against one
// server config. The configs are nested as sub-benchmarks so the resulting
// names are paths, e.g. Large/Chunking/Pipelining/Reuse/BytesBuffer, which can
// be filtered with -bench 'Large/Chunking/.*/Reuse'.
func benchmarkServer(b *testing.B, tc testcase, sc serverConfig) {
	b.Helper()

	_, s, addr, err := testServer(nil, sc.opts...)
	require.NoError(b, err)

	defer func() {
		require.NoError(b, s.Close())
	}()

	for _, cc := range clientConfigs {
		b.Run(cc.name, func(b *testing.B) {
			for _, conn := range connectionConfigs {
				b.Run(conn.name, func(b *testing.B) {
					for _, rc := range readerConfigs {
						b.Run(rc.name, func(b *testing.B) {
							setBytes(b, tc.eml)

							conn.run(b, addr, cc.opts, func(c *mailer.Mailer) error {
								return sendMailCon(c, rc.new(tc.eml))
							})
						})
					}
				})
			}
		})
	}
}

func BenchmarkMailer(b *testing.B) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
	slog.SetDefault(l)

	smallEml, err := embedTestdata.ReadFile("testdata/small.eml")
	require.NoError(b, err)

	largeEml, err := embedTestdata.ReadFile("testdata/large.eml")
	require.NoError(b, err)

	for _, tc := range []testcase{
		{eml: smallEml, name: "Small"},
		{eml: largeEml, name: "Large"},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for _, sc := range serverConfigs {
				b.Run(sc.name, func(b *testing.B) {
					benchmarkServer(b, tc, sc)
				})
			}
		})
	}
}
