package client

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"

	"github.com/uponusolutions/go-sasl"
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

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Verify(addr string) error {
	return c.c.Verify(addr)
}

// Close is a wrapper for Quit.
func (c *Client) Close() error {
	return c.Quit()
}

// Auth authenticates a client using the provided authentication mechanism.
// Only servers that advertise the AUTH extension support this function.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Auth(a sasl.Client) error {
	return c.c.Auth(a)
}

// Reset sends the RSET command to the server, aborting the current mail
// transaction.
func (c *Client) Reset() error {
	return c.c.Reset()
}

// Mail issues a MAIL command to the server using the provided email address.
// If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME
// parameter.
// This initiates a mail transaction and is followed by one or more Rcpt calls.
//
// If opts is not nil, MAIL arguments provided in the structure will be added
// to the command. Handling of unsupported options depends on the extension.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Mail(from string, opts *smtp.MailOptions) error {
	return c.c.Mail(from, opts)
}

// Rcpt issues a RCPT command to the server using the provided email address.
// A call to Rcpt must be preceded by a call to Mail and may be followed by
// a Data call or another Rcpt call.
//
// If opts is not nil, RCPT arguments provided in the structure will be added
// to the command. Handling of unsupported options depends on the extension.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Rcpt(to string, opts *smtp.RcptOptions) error {
	return c.c.Rcpt(to, opts)
}

// Data issues a DATA command to the server and returns a writer that
// can be used to write the mail headers and body. The caller should
// close the writer before calling any more methods on c. A call to
// Data must be preceded by one or more calls to Rcpt.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Data() (*client.DataCloser, error) {
	return c.c.Data()
}

func (c *Client) prepare(from string, rcpt []string) (*client.DataCloser, error) {
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

// SendMailWithResponse send a mail and returns the response code and message.
// Client must be connected to server.
func (c *Client) SendMailWithResponse(in io.Reader, from string, rcpt []string) (code int, msg string, err error) {
	return c.sendMail(in, from, rcpt)
}

// SendMail send a mail.
// Client must be connected to server.
func (c *Client) SendMail(in io.Reader, from string, rcpt []string) error {
	_, _, err := c.sendMail(in, from, rcpt)
	return err
}

func (c *Client) SetXOORG(xoorg *string) {
	c.options.XOORG = xoorg
}

func (c *Client) sendMail(in io.Reader, from string, rcpt []string) (code int, msg string, err error) {
	w, err := c.prepare(from, rcpt)
	if err != nil {
		return 0, "", err
	}

	_, err = io.Copy(w, in)
	if err != nil {
		return 0, "", err
	}

	return w.CloseWithResponse()
}

// Send implements enmime.Sender interface.
func (c *Client) Send(reversePath string, recipients []string, msg []byte) error {
	_, _, err := c.sendMail(bytes.NewBuffer(msg), reversePath, recipients)
	return err
}
