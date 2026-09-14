# go-smtp

[![coverage](https://raw.githubusercontent.com/uponusolutions/go-smtp/badges/.badges/main/coverage.svg)](https://uponusolutions.github.io/go-smtp/coverage/)
[![reference](https://pkg.go.dev/badge/github.com/uponusolutions/go-smtp.svg)](https://pkg.go.dev/github.com/uponusolutions/go-smtp)

An ESMTP client and server library written in Go.

It is a fork of [emersion/go-smtp](https://github.com/emersion/go-smtp) with a
reworked, context-aware API, a high-level mailer that resolves MX records and
reports per-recipient results, and support for extensions such as CHUNKING/BDAT,
DELIVERBY, MT-PRIORITY, RRVS and XOORG. Both client and server speak PIPELINING.

```sh
go get github.com/uponusolutions/go-smtp
```

## Packages

| Package | Description |
| --- | --- |
| [`smtp`](https://pkg.go.dev/github.com/uponusolutions/go-smtp) | Shared definitions: `Status`, `MailOptions`, `RcptOptions`, BATV/SRS sender parsing |
| [`mailer`](https://pkg.go.dev/github.com/uponusolutions/go-smtp/mailer) | High-level client: connection reuse, security, auth, sending to many servers |
| [`client`](https://pkg.go.dev/github.com/uponusolutions/go-smtp/client) | Low-level client, one method per SMTP command |
| [`server`](https://pkg.go.dev/github.com/uponusolutions/go-smtp/server) | SMTP server built around a `Backend`/`Session` pair |
| [`resolve`](https://pkg.go.dev/github.com/uponusolutions/go-smtp/resolve) | MX lookup that groups recipients by target server |
| [`tester`](https://pkg.go.dev/github.com/uponusolutions/go-smtp/tester) | Test utilities: fake connections, certificates, in-memory server |

See [examples](https://github.com/uponusolutions/go-smtp/tree/main/examples) for
runnable programs.

## Sending mail

`mailer.Send` resolves the MX records for every recipient, groups them by target
server and sends one transaction per server. The message is passed as a factory
because it is consumed once per server.

```go
report, err := mailer.Send(
    ctx,
    "alice@example.com",
    []string{"bob@example.org", "carol@example.net"},
    func() io.Reader { return strings.NewReader(eml) },
)
if err != nil {
    return err
}

for _, f := range report.Failures {
    fmt.Printf("failed for %v: %v\n", f.Rcpts, f.Error)
}
for _, r := range report.Responses {
    fmt.Printf("accepted for %v: %d %s\n", r.Rcpts, r.Status.Code, r.Status.Text())
}
```

Rejected recipients end up in `Failures` instead of aborting the whole send, so a
partial delivery is still reported as a success for the recipients that were
accepted.

To relay through a fixed host instead of resolving MX records, pass
`mailer.WithServerAddresses`. A `*mailer.Mailer` keeps its connection open across
calls, which is what you want for a queue runner:

```go
m := mailer.New(
    mailer.WithServerAddresses("mail.example.com:587"),
    mailer.WithSecurity(mailer.SecurityStartTLS),
    mailer.WithSASLClient(sasl.NewPlainClient("", "username", "password")),
    mailer.WithBasic(client.WithPipelining(true)),
)
defer func() { _ = m.Disconnect() }()

status, failures, err := m.Send(ctx, from, rcpts, msg)
```

`WithServerAddressesPrio` takes groups of addresses that are tried in order, and
`SendAdvanced` accepts explicit `MailOptions` and per-recipient `RcptOptions`.

## Running a server

A server is created with functional options and a `Backend` that returns a
`Session` per connection. Every method receives a `context.Context`; `NewSession`
and `Reset` may return a derived one that replaces the session context.

```go
type Backend struct{}

func (*Backend) NewSession(ctx context.Context, _ *server.Conn) (context.Context, server.Session, error) {
    return ctx, &Session{}, nil
}

func (s *Session) Mail(ctx context.Context, from string, opts *smtp.MailOptions) error { ... }
func (s *Session) Rcpt(ctx context.Context, to string, opts *smtp.RcptOptions) error   { ... }

// Data must consume the reader completely before returning.
// The returned queue id ends up in the 250 reply as "OK: queued as <id>".
func (s *Session) Data(ctx context.Context, r func() io.Reader) (string, error) { ... }
```

```go
s := server.New(
    server.WithBackend(&Backend{}),
    server.WithAddr("localhost:1025"),
    server.WithHostname("localhost"),
    server.WithReadTimeout(10*time.Second),
    server.WithWriteTimeout(10*time.Second),
    server.WithMaxMessageBytes(1024*1024),
    server.WithMaxRecipients(50),
    server.WithEnableCHUNKING(true),
    server.WithEnforceSecureConnection(true),
)

if err := s.ListenAndServe(ctx); err != nil {
    log.Fatal(err)
}
```

Errors returned from session methods are sent to the client. Return an
`*smtp.Status` to control the reply code, enhanced code and (multi-line) text;
any other error becomes a generic failure and is logged through the session
logger.

`server.WithEnforceSecureConnection` rejects everything except NOOP, EHLO,
STARTTLS and QUIT until the connection is encrypted, and
`server.WithEnforceAuthentication` does the same until authentication succeeded.

## Extensions

| Extension | RFC | Client | Server |
| --- | --- | --- | --- |
| PIPELINING | [2920](https://tools.ietf.org/html/rfc2920) | opt-in | always |
| 8BITMIME | [6152](https://tools.ietf.org/html/rfc6152) | yes | always |
| ENHANCEDSTATUSCODES | [2034](https://tools.ietf.org/html/rfc2034) | yes | always |
| SIZE | [1870](https://tools.ietf.org/html/rfc1870) | yes | always |
| STARTTLS | [3207](https://tools.ietf.org/html/rfc3207) | yes | with TLS config |
| AUTH | [4954](https://tools.ietf.org/html/rfc4954) | yes | via [go-sasl](https://github.com/uponusolutions/go-sasl) |
| CHUNKING / BDAT | [3030](https://tools.ietf.org/html/rfc3030) | yes | opt-in |
| BINARYMIME | [3030](https://tools.ietf.org/html/rfc3030) | yes | opt-in |
| SMTPUTF8 | [6531](https://tools.ietf.org/html/rfc6531) | yes | opt-in |
| DSN | [3461](https://tools.ietf.org/html/rfc3461) | yes | opt-in |
| REQUIRETLS | [8689](https://tools.ietf.org/html/rfc8689) | yes | opt-in |
| RRVS | [7293](https://tools.ietf.org/html/rfc7293) | yes | opt-in |
| DELIVERBY | [2852](https://tools.ietf.org/html/rfc2852) | yes | opt-in |
| MT-PRIORITY | [6710](https://tools.ietf.org/html/rfc6710) | yes | opt-in |
| LIMITS (RCPTMAX) | [9422](https://tools.ietf.org/html/rfc9422) | — | with max recipients |
| XOORG | non-standard | yes | opt-in |

Server-side extensions are advertised only when enabled, because advertising one
is a promise the backend has to keep. XOORG carries the accepted domain used by
Exchange Online outgoing connectors.

## Performance

The hot paths are the dot and BDAT readers and writers, which stream without
copying the whole message into memory. The client picks BDAT over DATA when the
server advertises CHUNKING; `client.WithChunkingMaxSize` caps the chunk size and
`client.WithChunkingBuffer` controls whether small reads are coalesced before
being written out. Pipelining removes a round trip per command in both the mailer
and the server.

Benchmarks run in CI and are published on the
[benchmark page](https://uponusolutions.github.io/go-smtp/dev/bench/); coverage
is on the [coverage page](https://uponusolutions.github.io/go-smtp/coverage/).

## Differences from emersion/go-smtp

The API is not compatible with upstream. The notable changes:

- **Split into packages.** Client, server, mailer, MX resolution and test helpers
  live in their own packages instead of a single `smtp` package.
- **Functional options.** `server.New(server.With...)` and
  `client.New(client.With...)` replace exported struct fields, so configuration
  cannot be mutated while a server is running.
- **Contexts everywhere.** Every `Backend` and `Session` method takes a
  `context.Context`, and `NewSession`/`Reset` can return a derived context that
  the rest of the session uses.
- **`Session.Data` takes `func() io.Reader`** and returns a queue id, which makes
  it possible to skip reading the body and to report an identifier back to the
  client.
- **Structured errors and logging.** Replies are `*smtp.Status` values with
  multi-line support, and each session can provide its own `*slog.Logger`.
- **More extensions.** Client CHUNKING/BDAT and XOORG are supported in addition to what upstream offers.
- **Pipelining on both sides**, opt-in for the client.
- **The mailer**, which resolves MX records, sends to several servers per call
  and reports per-recipient failures.
- **No LMTP.** Upstream supports LMTP; this fork does not.

If you need a very stable API or LMTP, use upstream. It is a well-maintained
library and this fork exists only because we needed a different set of tradeoffs.

## Relationship with net/smtp

The Go standard library provides an SMTP client implementation in `net/smtp`.
However `net/smtp` is frozen: it's not getting any new features. go-smtp provides
a server implementation and a number of client improvements.

## Development

```sh
make test    # go test ./...
make race    # tests with -race
make cover   # coverage profile and HTML report
make lint    # golangci-lint
make bench   # benchmarks, BENCH=... TIME=... to narrow them down
make stats   # benchstat over .bench/*.txt
```

## Licence

MIT
