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
			serverAddresses: [][]string{{"127.0.0.1:25"}},
			security:        SecurityPreferStartTLS,
		},
		client: client.DefaultConfig(),
	}
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

// additionalConfig contains all the extra configuration needed to define an smtp client.
type additionalConfig struct {
	serverAddresses    [][]string  // Format address:port.
	serverAddressIndex int         // first server address to try
	saslClient         sasl.Client // support authentication
	security           Security    // 	// Defines the connection is secured
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

// WithSASLClient sets the SASL client.
func WithSASLClient(cl sasl.Client) Option {
	return func(c *Config) {
		c.extra.saslClient = cl
	}
}

// WithSecurity sets the TLS config.
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
