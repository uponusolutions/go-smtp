package client

import (
	"crypto/tls"
	"io"
	"time"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

const defaultChunkingMaxSize = 1048576 * 2

var defaultConfig = config{
	extra: extraConfig{
		serverAddresses: [][]string{{"127.0.0.1:25"}},
		security:        SecurityPreferStartTLS,
	},
	basic: basicConfig{
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

		// Reader buffer of textproto
		readerSize: 4096,
		// Writer buffer of textproto
		writerSize: 4096,

		// Default chunking max size, 2 MiB
		chunkingMaxSize: defaultChunkingMaxSize,

		// chunking buffer enabled by default
		chunkingBufferEnabled: true,
	},
}

// Security describes how the connection is etablished.
type Security int32

const (
	// SecurityPreferStartTLS tries to use StartTls but fallbacks to plain.
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

// basicConfig contains all configuration needed to configure a smtp client.
type basicConfig struct {
	text      *textsmtp.Textproto
	localName string // the name to use in HELO/EHLO/LHLO

	// Time to wait for tls handshake to succeed.
	tlsHandshakeTimeout time.Duration

	// Time to wait for dial to succeed.
	dialTimeout time.Duration

	// Time to wait for command responses (this includes 3xx reply to DATA).
	commandTimeout time.Duration

	// Time to wait for responses after final dot.
	submissionTimeout time.Duration

	// Max line length, defaults to 2000
	maxLineLength int

	// Reader size
	readerSize int

	// Writer size
	writerSize int

	// Logger for all network activity.
	debug io.Writer

	mailOptions MailOptions

	// Chunking max size
	// A zero value disables chunk size limitation.
	// A negative value disables chunking from the client.
	chunkingMaxSize int

	// If no size is available and chunkingMaxSize > 4096 then
	// the buffer is automatically used.
	// If you guarantee that you reader has large enough chunks,
	// you can disable the chunking buffer here.
	chunkingBufferEnabled bool
}

// extraConfig contains all the extra configuration needed to define an smtp client.
type extraConfig struct {
	serverAddresses    [][]string  // Format address:port.
	serverAddressIndex int         // first server address to try
	saslClient         sasl.Client // support authentication
	security           Security    // 	// Defines the connection is secured
	tlsConfig          *tls.Config
}

// extraConfig contains all the extra configuration needed to define an smtp client.
type config struct {
	extra extraConfig
	basic basicConfig
}

// Option defines a client option.
type Option func(c *config)

// WithBasic allows to set all settings of the basic smtp client.
func WithBasic(opts ...BasicOption) Option {
	return func(c *config) {
		for _, e := range opts {
			e(&c.basic)
		}
	}
}

// WithServerAddresses sets the SMTP servers address.
func WithServerAddresses(addrs ...string) Option {
	return func(c *config) {
		c.extra.serverAddresses = [][]string{addrs}
	}
}

// WithServerAddressesPrio sets the SMTP servers address.
func WithServerAddressesPrio(addrs ...[]string) Option {
	return func(c *config) {
		c.extra.serverAddresses = addrs
	}
}

// WithServerAddressIndex sets the SMTP server index.
func WithServerAddressIndex(index int) Option {
	return func(c *config) {
		c.extra.serverAddressIndex = index
	}
}

// WithSASLClient sets the SASL client.
func WithSASLClient(cl sasl.Client) Option {
	return func(c *config) {
		c.extra.saslClient = cl
	}
}

// WithSecurity sets the TLS config.
func WithSecurity(security Security) Option {
	return func(c *config) {
		c.extra.security = security
	}
}

// WithTLSConfig sets the TLS config.
func WithTLSConfig(cfg *tls.Config) Option {
	return func(c *config) {
		c.extra.tlsConfig = cfg
	}
}

// BasicOption defines a client option.
type BasicOption func(c *basicConfig)

// WithMailOptions sets the mail options.
func WithMailOptions(mailOptions MailOptions) BasicOption {
	return func(c *basicConfig) {
		c.mailOptions = mailOptions
	}
}

// WithSubmissionTimeout sets the submission timeout.
func WithSubmissionTimeout(submissionTimeout time.Duration) BasicOption {
	return func(c *basicConfig) {
		c.submissionTimeout = submissionTimeout
	}
}

// WithCommandTimeout sets the command timeout.
func WithCommandTimeout(commandTimeout time.Duration) BasicOption {
	return func(c *basicConfig) {
		c.commandTimeout = commandTimeout
	}
}

// WithDialTimeout sets the dial timeout.
func WithDialTimeout(dialTimeout time.Duration) BasicOption {
	return func(c *basicConfig) {
		c.dialTimeout = dialTimeout
	}
}

// WithTlsHandshakeTimeout sets tls handshake timeout.
func WithTlsHandshakeTimeout(tlsHandshakeTimeout time.Duration) BasicOption {
	return func(c *basicConfig) {
		c.tlsHandshakeTimeout = tlsHandshakeTimeout
	}
}

// WithLocalName sets the HELO local name.
func WithLocalName(localName string) BasicOption {
	return func(c *basicConfig) {
		c.localName = localName
	}
}

// WithMaxLineLength sets the max line length.
func WithMaxLineLength(maxLineLength int) BasicOption {
	return func(c *basicConfig) {
		c.maxLineLength = maxLineLength
	}
}

// WithChunkingMaxSize sets the chunking max size.
// A zero value disables chunk size limitation.
// A negative value disables chunking from the client.
func WithChunkingMaxSize(chunkingMaxSize int) BasicOption {
	return func(c *basicConfig) {
		c.chunkingMaxSize = chunkingMaxSize
	}
}

// WithChunkingBuffer sets if the chunking buffer is used when necessary
// If no size is available and chunkingMaxSize > 4096 then
// the buffer is automatically used.
// If you guarantee that you reader has large enough chunks,
// you can disable the chunking buffer here.
// It is enabled by default.
func WithChunkingBuffer(enabled bool) BasicOption {
	return func(c *basicConfig) {
		c.chunkingBufferEnabled = enabled
	}
}

// WithReaderSize sets the reader size.
func WithReaderSize(readerSize int) BasicOption {
	return func(c *basicConfig) {
		c.readerSize = readerSize
	}
}

// WithWriterSize sets the reader size.
func WithWriterSize(writerSize int) BasicOption {
	return func(c *basicConfig) {
		c.writerSize = writerSize
	}
}
