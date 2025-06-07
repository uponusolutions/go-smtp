package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"strings"
	"time"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

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

// Client is an SMTP client.
// It sends one or more mails to a SMTP server over a single connection.
// TODO: Add context support.
type Client struct {
	serverAddress string // Format address:port.
	tlsConfig     *tls.Config
	saslClient    sasl.Client

	// keep a reference to the connection so it can be used to create a TLS
	// connection later
	conn       net.Conn
	text       *textsmtp.Textproto
	serverName string
	ext        map[string]string // supported extensions
	localName  string            // the name to use in HELO/EHLO/LHLO

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

	// Defines the connection is secured
	security Security

	mailOptions smtp.MailOptions
}

// NewClient returns a new SMTP client.
// When not set via options the address 127.0.0.1:25 is used.
// When not set via options a default tls.Config is used.
func NewClient(opts ...Option) *Client {
	c := &Client{
		serverAddress: "127.0.0.1:25",

		security: SecurityPreferStartTLS,

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

		// UTF8 is enabled by default but gets disabled if the server doesn't support it
		mailOptions: smtp.MailOptions{
			UTF8: true,
		},
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// Option defines a client option.
type Option func(c *Client)

// WithServerAddress Sets the SMTP servers address.
func WithServerAddress(addr string) Option {
	return func(c *Client) {
		c.serverAddress = addr
	}
}

// WithMailOptions sets the mail options.
func WithMailOptions(mailOptions smtp.MailOptions) Option {
	return func(c *Client) {
		c.mailOptions = mailOptions
	}
}

// WithSubmissionTimeout sets the submission timeout.
func WithSubmissionTimeout(submissionTimeout time.Duration) Option {
	return func(c *Client) {
		c.submissionTimeout = submissionTimeout
	}
}

// WithCommandTimeout sets the command timeout.
func WithCommandTimeout(commandTimeout time.Duration) Option {
	return func(c *Client) {
		c.commandTimeout = commandTimeout
	}
}

// WithDialTimeout sets the dial timeout.
func WithDialTimeout(dialTimeout time.Duration) Option {
	return func(c *Client) {
		c.dialTimeout = dialTimeout
	}
}

// WithTlsHandshakeTimeout sets tls handshake timeout.
func WithTlsHandshakeTimeout(tlsHandshakeTimeout time.Duration) Option {
	return func(c *Client) {
		c.tlsHandshakeTimeout = tlsHandshakeTimeout
	}
}

// WithLocalName sets the HELO local name.
func WithLocalName(localName string) Option {
	return func(c *Client) {
		c.localName = localName
	}
}

// WithTLSConfig sets the TLS config.
func WithTLSConfig(cfg *tls.Config) Option {
	return func(c *Client) {
		c.tlsConfig = cfg
	}
}

// WithSecurity sets the TLS config.
func WithSecurity(security Security) Option {
	return func(c *Client) {
		c.security = security
	}
}

// WithSASLClient sets the SASL client.
func WithSASLClient(cl sasl.Client) Option {
	return func(c *Client) {
		c.saslClient = cl
	}
}

// WithMaxLineLength sets the max line length.
func WithMaxLineLength(maxLineLength int) Option {
	return func(c *Client) {
		c.maxLineLength = maxLineLength
	}
}

// WithReaderSize sets the reader size.
func WithReaderSize(readerSize int) Option {
	return func(c *Client) {
		c.readerSize = readerSize
	}
}

// WithWriterSize sets the reader size.
func WithWriterSize(writerSize int) Option {
	return func(c *Client) {
		c.writerSize = writerSize
	}
}

// Connect connects to the SMTP server.
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// SMTP-UTF8 is enabled, when server supports it.
// If an error occures, the connection is closed if open.
func (c *Client) Connect(ctx context.Context) error {
	// verify if local name is valid
	if strings.ContainsAny(c.localName, "\n\r") {
		return errors.New("smtp: the local name must not contain CR or LF")
	}

	var err error
	var conn net.Conn

	switch c.security {
	case SecurityPlain:
		fallthrough
	case SecurityStartTLS:
		fallthrough
	case SecurityPreferStartTLS:
		conn, err = c.dial(ctx)
	case SecurityTLS:
		conn, err = c.dialTLS(ctx)
	}

	if err != nil {
		return err
	}

	c.setConn(conn)
	c.serverName, _, _ = net.SplitHostPort(c.serverAddress)

	if err = c.greet(); err != nil {
		return err
	}

	if err = c.hello(); err != nil {
		return err
	}

	if c.security == SecurityStartTLS || c.security == SecurityPreferStartTLS {
		if ok, _ := c.Extension("STARTTLS"); !ok {
			if c.security == SecurityStartTLS {
				_ = c.Quit()
				return errors.New("smtp: server doesn't support STARTTLS")
			}
		} else {
			err = c.startTLS()
			if err != nil {
				return err
			}
		}
	}

	return c.authAndUTF8()
}

func (c *Client) authAndUTF8() error {
	// Authenticate if authentication is possible and sasl client available.
	if ok, _ := c.Extension("AUTH"); ok && c.saslClient != nil {
		if err := c.Auth(c.saslClient); err != nil {
			_ = c.Quit()
			return err
		}
	}

	// Disable UTF8 if not supported by the server.
	if c.mailOptions.UTF8 {
		if ok, _ := c.Extension("SMTPUTF8"); !ok {
			c.mailOptions.UTF8 = false
		}
	}

	return nil
}

func (c *Client) prepare(from string, rcpt []string) (*DataCloser, error) {
	if c.conn == nil {
		return nil, errors.New("client is nil or not connected")
	}

	if len(rcpt) < 1 {
		return nil, errors.New("no recipients")
	}

	// MAIL FROM:
	if err := c.Mail(from, &c.mailOptions); err != nil {
		return nil, err
	}

	// RCPT TO:
	for _, addr := range rcpt {
		if err := c.Rcpt(addr, &smtp.RcptOptions{}); err != nil {
			return nil, err
		}
	}

	// DATA
	w, err := c.Data()
	if err != nil {
		return nil, err
	}
	return w, nil
}

// SendMail will use an existing connection to send an email from
// address from, to addresses to, with message r.
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
func (c *Client) SendMail(from string, rcpt []string, in io.Reader) (code int, msg string, err error) {
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

// SetXOORG set xoorg support
func (c *Client) SetXOORG(xoorg *string) {
	c.mailOptions.XOORG = xoorg
}

// Send implements enmime.Sender interface.
func (c *Client) Send(from string, rcpt []string, msg []byte) error {
	_, _, err := c.SendMail(from, rcpt, bytes.NewBuffer(msg))
	return err
}
