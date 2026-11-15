package mailer

import (
	"crypto/tls"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp/client"
)

// DefaultConfig returns the default configuration of a mailer.
func DefaultConfig() Config {
	return Config{
		extra: additionalConfig{
			// Mandatory STARTTLS (RFC 3207) with certificate verification, so
			// mail and credentials are never sent unencrypted by default.
			security: SecurityStartTLS,
		},
		client: client.DefaultConfig(),
	}
}

// NewConfig creates a new config with the given options.
func NewConfig(opts ...Option) Config {
	cfg := DefaultConfig()

	for _, o := range opts {
		o(&cfg)
	}

	return cfg
}

// Security describes how the connection is etablished.
type Security int32

const (
	// SecurityPreferStartTLS tries to use StartTLS but fallbacks to plain if
	// StartTLS is not available, rejected by the server or the TLS handshake
	// failed. A handshake failure includes the client rejecting the server
	// certificate (e.g. expired or not matching the host name): the mail is
	// then sent over a plain connection instead of returning an error.
	//
	// Security note: this is opportunistic TLS and provides no protection
	// against an active attacker. An attacker on the network path can strip
	// STARTTLS from the EHLO response or break the TLS handshake to force the
	// fallback to a plain connection.
	//
	// If a SASL client is configured, authentication is still performed after
	// such a fallback, i.e. credentials are sent over the unencrypted
	// connection (for mechanisms like PLAIN or LOGIN the password is only
	// base64 encoded). This is intended: the mode prefers delivering over
	// failing. If credentials must never be sent in plain text, use
	// SecurityStartTLS or SecurityTLS instead.
	SecurityPreferStartTLS Security = 0
	// SecurityPlain is always just a plain connection.
	SecurityPlain Security = 1
	// SecurityTLS does a implicit tls connection.
	SecurityTLS Security = 2
	// SecurityStartTLS always does starttls and fails if the server doesn't
	// support it or the TLS handshake fails. This is the default.
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

// additionalConfig contains all the extra configuration needed to define an smtp client.
type additionalConfig struct {
	serverAddresses    [][]string  // Format address:port.
	serverAddressIndex int         // first server address to try
	saslClient         sasl.Client // support authentication
	security           Security    // Defines the connection is secured
	abortOnRcptReject  bool        // Send a mail even if some recipients aren't accepted
	tlsConfig          *tls.Config
}

// Config contains a client config and the mailer config additions.
type Config struct {
	extra  additionalConfig
	client client.Config
}

// Option defines a client option.
type Option func(c *Config)

// WithBasic allows to set all settings of the basic smtp client.
func WithBasic(opts ...client.Option) Option {
	return func(c *Config) {
		for _, e := range opts {
			e(&c.client)
		}
	}
}

// WithServerAddresses sets the SMTP servers address.
func WithServerAddresses(addrs ...string) Option {
	return func(c *Config) {
		c.extra.serverAddresses = [][]string{addrs}
	}
}

// WithServerAddressesPrio sets the SMTP servers address.
func WithServerAddressesPrio(addrs ...[]string) Option {
	return func(c *Config) {
		c.extra.serverAddresses = addrs
	}
}

// WithServerAddressIndex sets the SMTP server index.
func WithServerAddressIndex(index int) Option {
	return func(c *Config) {
		c.extra.serverAddressIndex = index
	}
}

// WithSASLClient sets the SASL client used to authenticate if the server
// supports AUTH. Defaults to nil, no authentication.
//
// With SecurityPreferStartTLS or SecurityPlain the credentials may be sent
// unencrypted, the default SecurityStartTLS and SecurityTLS never do that.
func WithSASLClient(cl sasl.Client) Option {
	return func(c *Config) {
		c.extra.saslClient = cl
	}
}

// WithSecurity sets how the connection is secured.
// Defaults to SecurityStartTLS: STARTTLS is required and the server
// certificate is verified, otherwise connecting fails.
func WithSecurity(security Security) Option {
	return func(c *Config) {
		c.extra.security = security
	}
}

// WithTLSConfig sets the TLS config.
func WithTLSConfig(cfg *tls.Config) Option {
	return func(c *Config) {
		c.extra.tlsConfig = cfg
	}
}

// WithAbortOnRcptReject aborts sending if at last one recipient is rejected by the server.
func WithAbortOnRcptReject(abortOnRcptReject bool) Option {
	return func(c *Config) {
		c.extra.abortOnRcptReject = abortOnRcptReject
	}
}
