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

const defaultChunkingMaxSize = 1048576 * 2

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

// MailOptions contains parameters for the MAIL command.
type MailOptions struct {
	// Size of the body. Can be 0 if not specified by client.
	Size int64

	// TLS is required for the message transmission.
	//
	// The message should be rejected if it can't be transmitted
	// with TLS.
	RequireTLS bool

	// The message envelope or message header contains UTF-8-encoded strings.
	// This flag is set by SMTPUTF8-aware (RFC 6531) client.
	UTF8 UTF8

	// Value of RET= argument, FULL or HDRS.
	Return smtp.DSNReturn

	// Envelope identifier set by the client.
	EnvelopeID string

	// Accepted Domain from Exchange Online, e.g. from OutgoingConnector
	XOORG *string

	// The authorization identity asserted by the message sender in decoded
	// form with angle brackets stripped.
	//
	// nil value indicates missing AUTH, non-nil empty string indicates
	// AUTH=<>.
	//
	// Defined in RFC 4954.
	Auth *string
}

// VrfyOptions contains parameters for the VRFY command.
type VrfyOptions struct {
	// The message envelope or message header contains UTF-8-encoded strings.
	// This flag is set by SMTPUTF8-aware (RFC 6531) client.
	UTF8 UTF8
}

// Client is an SMTP client.
// It sends one or more mails to a SMTP server over a single connection.
// TODO: Add context support.
type Client struct {
	serverAddresses    [][]string // Format address:port.
	serverAddressIndex int        // first server address to try
	tlsConfig          *tls.Config
	saslClient         sasl.Client

	// keep a reference to the connection so it can be used to create a TLS
	// connection later
	conn        net.Conn
	connAddress string // Format address:port.
	connName    string // server greet name

	text      *textsmtp.Textproto
	ext       map[string]string // supported extensions
	localName string            // the name to use in HELO/EHLO/LHLO

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

	// The buffer is used if chunking/bdat is used with buffering.
	// It is created on first use and it's size is chunkingMaxSize.
	chunkingBuffer []byte
}

// New returns a new SMTP client.
// When not set via options the address 127.0.0.1:25 is used.
// When not set via options a default tls.Config is used.
func New(opts ...Option) *Client {
	c := &Client{
		serverAddresses: [][]string{{"127.0.0.1:25"}},

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

		// Default chunking max size, 2 MiB
		chunkingMaxSize: defaultChunkingMaxSize,

		// chunking buffer enabled by default
		chunkingBufferEnabled: true,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// Option defines a client option.
type Option func(c *Client)

// WithServerAddresses sets the SMTP servers address.
func WithServerAddresses(addrs ...string) Option {
	return func(c *Client) {
		c.serverAddresses = [][]string{addrs}
	}
}

// WithServerAddressesPrio sets the SMTP servers address.
func WithServerAddressesPrio(addrs ...[]string) Option {
	return func(c *Client) {
		c.serverAddresses = addrs
	}
}

// WithServerAddressIndex sets the SMTP server index.
func WithServerAddressIndex(index int) Option {
	return func(c *Client) {
		c.serverAddressIndex = index
	}
}

// WithMailOptions sets the mail options.
func WithMailOptions(mailOptions MailOptions) Option {
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

// WithChunkingMaxSize sets the chunking max size.
// A zero value disables chunk size limitation.
// A negative value disables chunking from the client.
func WithChunkingMaxSize(chunkingMaxSize int) Option {
	return func(c *Client) {
		c.chunkingMaxSize = chunkingMaxSize
	}
}

// WithChunkingBuffer sets if the chunking buffer is used when necessary
// If no size is available and chunkingMaxSize > 4096 then
// the buffer is automatically used.
// If you guarantee that you reader has large enough chunks,
// you can disable the chunking buffer here.
// It is enabled by default.
func WithChunkingBuffer(enabled bool) Option {
	return func(c *Client) {
		c.chunkingBufferEnabled = enabled
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

// ServerAddresses returns the server address.
func (c *Client) ServerAddresses() [][]string {
	return c.serverAddresses
}

// ServerAddress returns the current server address.
func (c *Client) ServerAddress() string {
	return c.connAddress
}

// ServerName returns the current server name.
func (c *Client) ServerName() string {
	return c.connName
}

// Connect connects to one of the available SMTP server.
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// If an error occures, the connection is closed if open.
func (c *Client) Connect(ctx context.Context) error {
	// verify if local name is valid
	if strings.ContainsAny(c.localName, "\n\r") {
		return errors.New("smtp: the local name must not contain CR or LF")
	}

	var err error

	for i := 0; i < len(c.serverAddresses); i++ {
		for p := 0; p < len(c.serverAddresses[i]); p++ {
			// use c.serverAddressIndex
			address := c.serverAddresses[i][(p+c.serverAddressIndex)%len(c.serverAddresses[i])]
			err = c.connectAddress(ctx, address)
			if err == nil {
				c.connAddress = address
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
	var conn net.Conn

	switch c.security {
	case SecurityPlain:
		fallthrough
	case SecurityStartTLS:
		fallthrough
	case SecurityPreferStartTLS:
		conn, err = c.dial(ctx, addr)
	case SecurityTLS:
		conn, err = c.dialTLS(ctx, addr)
	}

	if err != nil {
		return err
	}

	c.setConn(conn)

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
			serverName, _, _ := net.SplitHostPort(addr)

			err = c.startTLS(serverName)
			if err != nil {
				return err
			}
		}
	}

	return c.auth()
}

func (c *Client) auth() error {
	// Authenticate if authentication is possible and sasl client available.
	if ok, _ := c.Extension("AUTH"); ok && c.saslClient != nil {
		if err := c.Auth(c.saslClient); err != nil {
			_ = c.Quit()
			return err
		}
	}
	return nil
}

// Len defines the Len method existing in some structs to get the length of the internal []byte (e.g. bytes.Buffer)
type Len interface {
	Len() int
}

func (c *Client) prepare(ctx context.Context, from string, rcpt []string, size int) (IDataCloser, error) {
	if c.conn == nil {
		err := c.Connect(ctx)
		if err != nil {
			return nil, err
		}
	}

	if len(rcpt) < 1 {
		return nil, errors.New("no recipients")
	}

	mailOptions := c.mailOptions
	if size > 0 && c.mailOptions.Size == 0 {
		mailOptions.Size = int64(size)
	}

	// MAIL FROM:
	if err := c.Mail(from, &mailOptions); err != nil {
		return nil, err
	}

	// RCPT TO:
	for _, addr := range rcpt {
		if err := c.Rcpt(addr, &smtp.RcptOptions{}); err != nil {
			return nil, err
		}
	}

	// DATA
	w, err := c.Content(size)
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

	_, err = io.Copy(w, in)
	if err != nil {
		// if err isn't smtp.Status we are in an unknown state, close connection
		if _, ok := err.(*smtp.Status); !ok {
			err = errors.Join(err, c.Close())
		}
		return 0, "", err
	}

	code, msg, err = w.CloseWithResponse()

	// if err isn't smtp.Status we are in an unknown state, close connection
	if _, ok := err.(*smtp.Status); err != nil && !ok {
		err = errors.Join(err, c.Close())
	}

	return code, msg, err
}
