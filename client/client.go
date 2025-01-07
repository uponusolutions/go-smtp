package client

import (
	"crypto/tls"
	"errors"
	"io"

	"github.com/emersion/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/client"
)

// Client is an SMTP client.
// It sends one or more mails to a SMTP server over a single connection.
// TODO: Add context support.
type Client struct {
	ServerAddress string // Format address:port.
	UseTLS        bool
	TLSConfig     *tls.Config
	SASLClient    sasl.Client
	HeloName      string

	c       *client.Client
	options *smtp.MailOptions
}

// NewClient returns a new SMTP client.
// When not set via options the address 127.0.0.1:25 is used.
// When not set via options a default tls.Config is used.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		ServerAddress: "127.0.0.1:25",
		HeloName:      "localhost",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// ClientOption defines a client option.
type ClientOption func(c *Client)

// WithServerAddress Sets the SMTP servers address.
func WithServerAddress(addr string) ClientOption {
	return func(c *Client) {
		c.ServerAddress = addr
	}
}

// WithHeloName sets the HELO local name.
func WithHeloName(localName string) ClientOption {
	return func(c *Client) {
		c.HeloName = localName
	}
}

// WithUseTLS sets use TLS.
func WithUseTLS(useTLS bool) ClientOption {
	return func(c *Client) {
		c.UseTLS = useTLS
	}
}

// WithTLSConfig sets the TLS config.
func WithTLSConfig(cfg *tls.Config) ClientOption {
	return func(c *Client) {
		c.TLSConfig = cfg
	}
}

// WithSASLClient sets the SASL client.
func WithSASLClient(cl sasl.Client) ClientOption {
	return func(c *Client) {
		c.SASLClient = cl
	}
}

// Connect connects to the SMTP server.
// When server supports STARTTLS, it's called with clients TLS config.
// When server supports auth and clients SaslClient is set, auth is called.
// SMTP-UTF8 is enabled, when server supports it.
func (c *Client) Connect() error {
	var err error
	if c.UseTLS {
		c.c, err = client.DialStartTLS(c.ServerAddress, c.TLSConfig)
	} else {
		c.c, err = client.Dial(c.ServerAddress)
	}
	if err != nil {
		return err
	}

	// Try STARTTLS.
	//if ok, _ := c.c.Extension("STARTTLS"); ok {
	//	if err = c.c.StartTLS(c.TLSConfig); err != nil {
	//		return err
	//	}
	//}

	return c.heloAuthAndUTF8()
}

// ConnectTLS directly connects via TLS using clients TLS config.
// When server supports auth and clients SaslClient is set, auth is called.
// SMTP-UTF8 is enabled, when server supports it.
func (c *Client) ConnectTLS() error {
	var err error
	c.c, err = client.DialTLS(c.ServerAddress, c.TLSConfig)
	if err != nil {
		return err
	}

	return c.heloAuthAndUTF8()
}

func (c *Client) heloAuthAndUTF8() error {
	// Set LocalName (otherwise localhost inside c.c.Mail)
	if err := c.c.Hello(c.HeloName); err != nil {
		return err
	}

	ok, _ := c.c.Extension("AUTH")
	if ok && c.SASLClient != nil {
		if err := c.c.Auth(c.SASLClient); err != nil {
			return err
		}
	}

	c.options = &smtp.MailOptions{}
	if ok, _ := c.c.Extension("SMTPUTF8"); ok {
		c.options.UTF8 = true
	}

	return nil
}

// Quit sends the QUIT command and closes the connection to the server.
// If Quit fails, it tries to close the server connection.
// It's safe to be called multiple times.
func (c *Client) Quit() error {
	if c.c == nil {
		return nil
	}

	err := c.c.Quit()
	if err != nil {
		err = c.c.Close()
	}

	c.c = nil

	return err
}

// Close is a wrapper for Quit.
func (c *Client) Close() error {
	return c.Quit()
}

func (c *Client) prepare(from string, rcpt []string) (io.WriteCloser, error) {
	if c.c == nil {
		return nil, errors.New("client is nil or not connected")
	}

	if len(rcpt) < 1 {
		return nil, errors.New("no recipients")
	}

	// MAIL FROM:
	if err := c.c.Mail(from, c.options); err != nil {
		return nil, err
	}

	// RCPT TO:
	// TODO: Maybe usful for SMTP-UTF8?
	rcptOpts := &smtp.RcptOptions{}

	for _, addr := range rcpt {
		if err := c.c.Rcpt(addr, rcptOpts); err != nil {
			return nil, err
		}
	}

	// DATA
	w, err := c.c.Data()
	if err != nil {
		return nil, err
	}
	return w, nil
}

// SendMail send a mail.
// Client must be connected to server.
func (c *Client) SendMail(in io.Reader, from string, rcpt []string) error {
	w, err := c.prepare(from, rcpt)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, in)
	if err != nil {
		return err
	}

	return w.Close()
}

// Send implements enmime.Sender interface.
func (c *Client) Send(reversePath string, recipients []string, msg []byte) error {

	w, err := c.prepare(reversePath, recipients)
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	return w.Close()
}
