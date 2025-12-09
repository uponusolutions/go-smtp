// Package mailer implements a mailer to send mail by smtp.
// It uses the client library and provides high level functionality.
package mailer

import (
	"context"
	"errors"
	"io"
	"net"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/client"
)

// Mailer implements a smtp client with .
type Mailer struct {
	client *client.Client
	cfg    additionalConfig
}

// New returns a new smtp client.
// When not set via options the address 127.0.0.1:25 is used.
// When not set via options a default tls.Config is used.
func New(opts ...Option) *Mailer {
	cfg := DefaultConfig()

	for _, o := range opts {
		o(&cfg)
	}

	return NewFromConfig(cfg)
}

// NewFromConfig returns a new smtp client from existing config.
func NewFromConfig(cfg Config) *Mailer {
	return &Mailer{
		cfg:    cfg.extra,
		client: client.NewFromConfig(cfg.client),
	}
}

// Connect connects to one of the available smtp server.
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// If an error occures, the connection is closed if open.
func (c *Mailer) Connect(ctx context.Context) error {
	var err error

	for i := 0; i < len(c.cfg.serverAddresses); i++ {
		for p := 0; p < len(c.cfg.serverAddresses[i]); p++ {
			// use c.serverAddressIndex
			address := c.cfg.serverAddresses[i][(p+c.cfg.serverAddressIndex)%len(c.cfg.serverAddresses[i])]
			err = c.connectAddress(ctx, address)
			if err == nil {
				return nil
			}
		}
	}

	return err
}

// Connect connects to the SMTP server (addr).
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// If an error occures, the connection is closed if open.
func (c *Mailer) connectAddress(ctx context.Context, addr string) error {
	var err error

	switch c.cfg.security {
	case SecurityPlain:
		fallthrough
	case SecurityStartTLS:
		fallthrough
	case SecurityPreferStartTLS:
		err = c.client.Dial(ctx, addr)
	case SecurityTLS:
		err = c.client.DialTLS(ctx, c.cfg.tlsConfig, addr)
	}

	if err != nil {
		return err
	}

	if c.cfg.security == SecurityStartTLS || c.cfg.security == SecurityPreferStartTLS {
		if ok, _ := c.client.Extension("STARTTLS"); !ok {
			if c.cfg.security == SecurityStartTLS {
				_ = c.client.Quit()
				return errors.New("smtp: server doesn't support STARTTLS")
			}
		} else {
			serverName, _, _ := net.SplitHostPort(addr)

			err = c.client.StartTLS(c.cfg.tlsConfig, serverName)
			if err != nil {
				return err
			}
		}
	}

	return c.auth()
}

func (c *Mailer) auth() error {
	// Authenticate if authentication is possible and sasl client available.
	if ok, _ := c.client.Extension("AUTH"); ok && c.cfg.saslClient != nil {
		if err := c.client.Auth(c.cfg.saslClient); err != nil {
			_ = c.client.Quit()
			return err
		}
	}
	return nil
}

// Len defines the Len method existing in some structs to get the length of the internal []byte (e.g. bytes.Buffer)
type Len interface {
	Len() int
}

func (c *Mailer) prepare(ctx context.Context, from string, options *client.MailOptions, rcpt []string, size int) (*client.DataCloser, error) {
	if c.client.Connected() {
		err := c.Connect(ctx)
		if err != nil {
			return nil, err
		}
	}

	if len(rcpt) < 1 {
		return nil, errors.New("no recipients")
	}

	if options == nil && size > 0 {
		options = &client.MailOptions{
			Size: int64(size),
		}
	}

	// MAIL FROM:
	if err := c.client.Mail(from, options); err != nil {
		return nil, err
	}

	// RCPT TO:
	for _, addr := range rcpt {
		if err := c.client.Rcpt(addr, &smtp.RcptOptions{}); err != nil {
			return nil, err
		}
	}

	// DATA
	w, err := c.client.Content(size)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// Send send an email from
// address from, to addresses to, with message r.
//
// It will use an existing connection if possible or create a new one otherwise.
//
// This function does not start TLS, nor does it perform authentication. Use
// DialStartTLS and Auth before-hand if desirable.
//
// The addresses in the to parameter are the SMTP RCPT addresses.
//
// The r parameter should be an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of r
// should be CRLF terminated. The r headers should usually include
// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
// messages is accomplished by including an email address in the to
// parameter but not including it in the r headers.
func (c *Mailer) Send(ctx context.Context, from string, rcpt []string, in io.Reader) (code int, msg string, err error) {
	return c.SendAdvanced(ctx, from, nil, rcpt, in)
}

// SendAdvanced send an email from
// address from, to addresses to, with message r.
//
// It will use an existing connection if possible or create a new one otherwise.
//
// This function does not start TLS, nor does it perform authentication. Use
// DialStartTLS and Auth before-hand if desirable.
//
// The addresses in the to parameter are the SMTP RCPT addresses.
//
// The r parameter should be an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of r
// should be CRLF terminated. The r headers should usually include
// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
// messages is accomplished by including an email address in the to
// parameter but not including it in the r headers.
func (c *Mailer) SendAdvanced(ctx context.Context, from string, options *client.MailOptions, rcpt []string, in io.Reader) (code int, msg string, err error) {
	size := 0
	if wt, ok := in.(Len); ok {
		size = wt.Len()
	}

	w, err := c.prepare(ctx, from, options, rcpt, size)
	if err != nil {
		return 0, "", err
	}

	_, err = io.Copy(w.Writer(), in)
	if err != nil {
		// if err isn't smtp.Status we are in an unknown state, close connection
		if _, ok := err.(*smtp.Status); !ok {
			err = errors.Join(err, c.client.Close())
		}
		return 0, "", err
	}

	code, msg, err = w.CloseWithResponse()

	// if err isn't smtp.Status we are in an unknown state, close connection
	if _, ok := err.(*smtp.Status); err != nil && !ok {
		err = errors.Join(err, c.client.Close())
	}

	return code, msg, err
}

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
//
// If server returns an error, it will be of type *smtp.
func (c *Mailer) Verify(addr string, opts *client.VrfyOptions) error {
	return c.client.Verify(addr, opts)
}

// Disconnect ends current connection gracefully, if any exists.
func (c *Mailer) Disconnect() error {
	return c.client.Quit()
}

// Terminate ends current connection forcefully.
func (c *Mailer) Terminate() error {
	return c.client.Close()
}

// Connected returns the current server name.
func (c *Mailer) Connected() bool {
	return c.client.Connected()
}

// ServerAddress returns the current server address.
func (c *Mailer) ServerAddress() string {
	return c.client.ServerAddress()
}

// ServerName returns the current server name.
func (c *Mailer) ServerName() string {
	return c.client.ServerName()
}
