// Package smtp implements the client of the Simple Mail Transfer Protocol as defined in RFC 5321.
//
// It also implements the following extensions:
//
//   - 8BITMIME (RFC 1652)
//   - AUTH (RFC 2554)
//   - STARTTLS (RFC 3207)
//   - ENHANCEDSTATUSCODES (RFC 2034)
//   - SMTPUTF8 (RFC 6531)
//   - REQUIRETLS (RFC 8689)
//   - CHUNKING (RFC 3030)
//   - BINARYMIME (RFC 3030)
//   - DSN (RFC 3461, RFC 6533)
//
// Additional extensions may be handled by other packages.
package client

import (
	"context"
	"errors"
	"io"
	"net"
	"strings"

	"github.com/uponusolutions/go-smtp"
)

// Client implements a smtp client with .
type Client struct {
	basic BasicClient
	cfg   config
}

// New returns a new smtp client.
// When not set via options the address 127.0.0.1:25 is used.
// When not set via options a default tls.Config is used.
func New(opts ...Option) *Client {
	c := defaultConfig

	for _, o := range opts {
		o(&c)
	}

	return &Client{
		cfg: c,
		basic: BasicClient{
			cfg: c.basic,
		},
	}
}

// Connect connects to one of the available smtp server.
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// If an error occures, the connection is closed if open.
func (c *Client) Connect(ctx context.Context) error {
	// verify if local name is valid
	if strings.ContainsAny(c.cfg.basic.localName, "\n\r") {
		return errors.New("smtp: the local name must not contain CR or LF")
	}

	var err error

	for i := 0; i < len(c.cfg.extra.serverAddresses); i++ {
		for p := 0; p < len(c.cfg.extra.serverAddresses[i]); p++ {
			// use c.serverAddressIndex
			address := c.cfg.extra.serverAddresses[i][(p+c.cfg.extra.serverAddressIndex)%len(c.cfg.extra.serverAddresses[i])]
			err = c.connectAddress(ctx, address)
			if err == nil {
				c.basic.connAddress = address
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
func (c *Client) connectAddress(ctx context.Context, addr string) error {
	var err error

	switch c.cfg.extra.security {
	case SecurityPlain:
		fallthrough
	case SecurityStartTLS:
		fallthrough
	case SecurityPreferStartTLS:
		err = c.basic.Dial(ctx, addr)
	case SecurityTLS:
		err = c.basic.DialTLS(ctx, c.cfg.extra.tlsConfig, addr)
	}

	if err != nil {
		return err
	}

	if c.cfg.extra.security == SecurityStartTLS || c.cfg.extra.security == SecurityPreferStartTLS {
		if ok, _ := c.basic.Extension("STARTTLS"); !ok {
			if c.cfg.extra.security == SecurityStartTLS {
				_ = c.basic.Quit()
				return errors.New("smtp: server doesn't support STARTTLS")
			}
		} else {
			serverName, _, _ := net.SplitHostPort(addr)

			err = c.basic.StartTLS(c.cfg.extra.tlsConfig, serverName)
			if err != nil {
				return err
			}
		}
	}

	return c.auth()
}

func (c *Client) auth() error {
	// Authenticate if authentication is possible and sasl client available.
	if ok, _ := c.basic.Extension("AUTH"); ok && c.cfg.extra.saslClient != nil {
		if err := c.basic.Auth(c.cfg.extra.saslClient); err != nil {
			_ = c.basic.Quit()
			return err
		}
	}
	return nil
}

// Len defines the Len method existing in some structs to get the length of the internal []byte (e.g. bytes.Buffer)
type Len interface {
	Len() int
}

func (c *Client) prepare(ctx context.Context, from string, rcpt []string, size int) (*DataCloser, error) {
	if c.basic.conn == nil {
		err := c.Connect(ctx)
		if err != nil {
			return nil, err
		}
	}

	if len(rcpt) < 1 {
		return nil, errors.New("no recipients")
	}

	mailOptions := c.cfg.basic.mailOptions
	if size > 0 && c.cfg.basic.mailOptions.Size == 0 {
		mailOptions.Size = int64(size)
	}

	// MAIL FROM:
	if err := c.basic.Mail(from, &mailOptions); err != nil {
		return nil, err
	}

	// RCPT TO:
	for _, addr := range rcpt {
		if err := c.basic.Rcpt(addr, &smtp.RcptOptions{}); err != nil {
			return nil, err
		}
	}

	// DATA
	w, err := c.basic.Content(size)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// SendMail send an email from
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
func (c *Client) SendMail(ctx context.Context, from string, rcpt []string, in io.Reader) (code int, msg string, err error) {
	size := 0
	if wt, ok := in.(Len); ok {
		size = wt.Len()
	}

	w, err := c.prepare(ctx, from, rcpt, size)
	if err != nil {
		return 0, "", err
	}

	_, err = io.Copy(w.Writer(), in)
	if err != nil {
		// if err isn't smtp.Status we are in an unknown state, close connection
		if _, ok := err.(*smtp.Status); !ok {
			err = errors.Join(err, c.basic.Close())
		}
		return 0, "", err
	}

	code, msg, err = w.CloseWithResponse()

	// if err isn't smtp.Status we are in an unknown state, close connection
	if _, ok := err.(*smtp.Status); err != nil && !ok {
		err = errors.Join(err, c.basic.Close())
	}

	return code, msg, err
}

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Verify(addr string, opts *VrfyOptions) error {
	return c.basic.Verify(addr, opts)
}

// Disconnect ends current connection gracefully, if any exists.
func (c *Client) Disconnect() error {
	return c.basic.Quit()
}

// Terminate ends current connection forcefully.
func (c *Client) Terminate() error {
	return c.basic.Close()
}

// ServerAddress returns the current server address.
func (c *Client) ServerAddress() string {
	return c.basic.connAddress
}

// ServerName returns the current server name.
func (c *Client) ServerName() string {
	return c.basic.connName
}
