package client

import (
	"syscall"
	"time"
)

// DefaultConfig returns the default configuration of a client.
func DefaultConfig() Config {
	return Config{
		localName: "localhost",
		// As recommended by RFC 5321. For DATA command reply (3xx one) RFC
		// recommends a slightly shorter timeout but we do not bother
		// differentiating these.
		commandTimeout: 5 * time.Minute,
		// 10 minutes + 2 minute buffer in case the server is doing transparent
		// forwarding and also follows recommended timeouts.
		submissionTimeout: 12 * time.Minute,
		// 30 seconds, very generous
		tlsHandshakeTimeout: 30 * time.Second,

		// 30 seconds, very generous
		dialTimeout: 30 * time.Second,

		// Doubled maximum line length per RFC 5321 (Section 4.5.3.1.6)
		maxLineLength: 2000,

		// Max lines of a single (multiline) server response.
		// Real world EHLO responses have less than 50 lines, so this
		// is generous while bounding a response to about 1 MB (with the default max
		// line length of 2000).
		maxResponseLines: 512,

		// Reader buffer of textproto
		readerSize: 4096,
		// Writer buffer of textproto
		writerSize: 4096,

		// Use chunking (BDAT) if the server supports it.
		chunking: true,

		// Messages of unknown size are buffered and sent in chunks of 2 MiB.
		chunkingBufferSize: 2 << 20,
	}
}

// Security describes how the connection is etablished.
type Security int32

const (
	// SecurityPreferStartTLS tries to use StartTls but fallbacks to plain.
	//
	// Security note: this is opportunistic TLS and an active attacker can
	// force the fallback to plain. Any authentication performed afterwards
	// sends the credentials unencrypted. Use SecurityStartTLS or SecurityTLS
	// if credentials must never be sent in plain text.
	SecurityPreferStartTLS Security = 0
	// SecurityPlain is always just a plain connection.
	SecurityPlain Security = 1
	// SecurityTLS does a implicit tls connection.
	SecurityTLS Security = 2
	// SecurityStartTLS always does starttls.
	SecurityStartTLS Security = 3
)

// UTF8 describes how SMTPUTF8 is used.
type UTF8 int32

const (
	// UTF8Prefer uses SMTPUTF8 if possible.
	UTF8Prefer UTF8 = 0
	// UTF8Force always uses SMTPUTF8.
	UTF8Force UTF8 = 1
	// UTF8Disabled never uses SMTPUTF8.
	UTF8Disabled UTF8 = 2
)

// Config contains all configuration needed to configure a smtp client.
type Config struct {
	localName string // the name to use in HELO/EHLO/LHLO

	// Time to wait for tls handshake to succeed.
	tlsHandshakeTimeout time.Duration

	// Time to wait for dial to succeed.
	dialTimeout time.Duration

	// Called with the resolved address of every connection before it is made, see net.Dialer.Control.
	dialControl func(network, address string, c syscall.RawConn) error

	// Time to wait for command responses (this includes 3xx reply to DATA).
	commandTimeout time.Duration

	// Time to wait for responses after final dot.
	submissionTimeout time.Duration

	// Max line length, defaults to 2000
	maxLineLength int

	// Max lines of a single (multiline) response, defaults to 512.
	// A zero or negative value disables the limit.
	maxResponseLines int

	// Reader size
	readerSize int

	// Writer size
	writerSize int

	// Use chunking (BDAT) if the server supports it, defaults to true.
	chunking bool

	// Size of the buffer used for chunking if the message size is unknown,
	// defaults to 2 MiB. Every full buffer is sent as one chunk.
	// A zero or negative value disables buffering.
	chunkingBufferSize int

	// Enable pipelining if the server supports it.
	pipelining bool
}

// Option defines a client option.
type Option func(c *Config)

// WithSubmissionTimeout sets the submission timeout.
func WithSubmissionTimeout(submissionTimeout time.Duration) Option {
	return func(c *Config) {
		c.submissionTimeout = submissionTimeout
	}
}

// WithCommandTimeout sets the command timeout.
func WithCommandTimeout(commandTimeout time.Duration) Option {
	return func(c *Config) {
		c.commandTimeout = commandTimeout
	}
}

// WithDialTimeout sets the dial timeout.
func WithDialTimeout(dialTimeout time.Duration) Option {
	return func(c *Config) {
		c.dialTimeout = dialTimeout
	}
}

// WithDialControl sets the net.Dialer.Control of the connections.
// It is called with the resolved address before connecting, an error refuses the connection.
// This lets callers refuse addresses, e.g. internal networks for servers users configure.
func WithDialControl(control func(network, address string, c syscall.RawConn) error) Option {
	return func(c *Config) {
		c.dialControl = control
	}
}

// WithTlsHandshakeTimeout sets tls handshake timeout.
func WithTlsHandshakeTimeout(tlsHandshakeTimeout time.Duration) Option {
	return func(c *Config) {
		c.tlsHandshakeTimeout = tlsHandshakeTimeout
	}
}

// WithLocalName sets the HELO local name.
func WithLocalName(localName string) Option {
	return func(c *Config) {
		c.localName = localName
	}
}

// WithMaxLineLength sets the max line length.
func WithMaxLineLength(maxLineLength int) Option {
	return func(c *Config) {
		c.maxLineLength = maxLineLength
	}
}

// WithMaxResponseLines sets the maximum number of lines of a single (multiline)
// server response. Defaults to 512.
// A zero or negative value disables the limit.
func WithMaxResponseLines(maxResponseLines int) Option {
	return func(c *Config) {
		c.maxResponseLines = maxResponseLines
	}
}

// WithChunking sets if chunking (BDAT, RFC 3030) is used when the server
// supports it. It is enabled by default. If disabled, DATA is used.
//
// A message of known size is always sent as a single chunk.
func WithChunking(enabled bool) Option {
	return func(c *Config) {
		c.chunking = enabled
	}
}

// WithChunkingBufferSize sets the size of the buffer used for chunking if
// the message size is unknown. Every full buffer is sent as one chunk, so it
// is the chunk size of such messages. Defaults to 2 MiB.
//
// A zero or negative value disables buffering, then every write is sent as
// its own chunk. Only do this if your writes are large enough.
func WithChunkingBufferSize(chunkingBufferSize int) Option {
	return func(c *Config) {
		c.chunkingBufferSize = chunkingBufferSize
	}
}

// WithPipelining sets if pipelining is used when the server supports it.
func WithPipelining(enabled bool) Option {
	return func(c *Config) {
		c.pipelining = enabled
	}
}

// WithReaderSize sets the reader size.
func WithReaderSize(readerSize int) Option {
	return func(c *Config) {
		c.readerSize = readerSize
	}
}

// WithWriterSize sets the reader size.
func WithWriterSize(writerSize int) Option {
	return func(c *Config) {
		c.writerSize = writerSize
	}
}
