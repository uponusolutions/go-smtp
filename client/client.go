// Package client implements the client of the Simple Mail Transfer Protocol as defined in RFC 5321.
// It provides basic functionality you need to use smtp.
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
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

// Client implements a SMTP Client with .
type Client struct {
	cfg Config

	// The buffer is used if chunking/bdat is used with buffering.
	// It is created on first use and it's size is chunkingMaxSize.
	chunkingBuffer []byte

	ext map[string]string // supported extensions of the server after ehlo

	// keep a reference to the connection so it can be used to create a TLS
	// connection later
	conn        net.Conn
	connAddress string // Format address:port.
	connName    string // server greet name
}

// New returns a new smtp client.
func New(opts ...Option) *Client {
	cfg := DefaultConfig()

	for _, o := range opts {
		o(&cfg)
	}

	return NewFromConfig(cfg)
}

// NewFromConfig returns a new smtp client from existing config.
func NewFromConfig(cfg Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

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

// Dial returns a connection to an SMTP server at addr. The addr must
// include a port, as in "mail.example.com:smtp".
func (c *Client) Dial(ctx context.Context, addr string) error {
	dialer := net.Dialer{Timeout: c.cfg.dialTimeout}
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return err
	}
	c.connAddress = addr

	return c.initConn(conn, true)
}

// DialTLS returns a connection to an SMTP server at addr via TLS.
// The addr must include a port, as in "mail.example.com:smtps".
//
// A nil tlsConfig is equivalent to a zero tls.Config.
func (c *Client) DialTLS(ctx context.Context, config *tls.Config, addr string) error {
	tlsDialer := tls.Dialer{
		NetDialer: &net.Dialer{Timeout: c.cfg.dialTimeout},
		Config:    config,
	}
	conn, err := tlsDialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return err
	}
	c.connAddress = addr

	return c.initConn(conn, true)
}

// initConn sets the underlying network connection for the client,
// expect greeting from the server if enabled and
// calls hello to finish basic setup
func (c *Client) initConn(conn net.Conn, expectGreet bool) error {
	c.setConn(conn)

	if expectGreet {
		if err := c.greet(); err != nil {
			return err
		}
	}

	return c.Hello()
}

// setConn sets the underlying network connection for the client.
func (c *Client) setConn(conn net.Conn) {
	c.conn = conn

	if c.cfg.debug != nil {
		c.cfg.text = textsmtp.NewTextproto(struct {
			io.Reader
			io.Writer
			io.Closer
		}{
			io.TeeReader(c.conn, c.cfg.debug),
			io.MultiWriter(c.conn, c.cfg.debug),
			c.conn,
		}, c.cfg.readerSize, c.cfg.writerSize, c.cfg.maxLineLength)
	}
	if c.cfg.text != nil {
		c.cfg.text.Replace(conn)
	} else {
		c.cfg.text = textsmtp.NewTextproto(conn, c.cfg.readerSize, c.cfg.writerSize, c.cfg.maxLineLength)
	}
}

// Close closes the connection.
func (c *Client) Close() error {
	if c.conn == nil {
		return nil
	}

	err := c.cfg.text.Close()
	c.conn = nil
	return err
}

// Greet reads the greeting of the server
// if an error occurred the connection is closed
func (c *Client) greet() error {
	// Initial greeting timeout. RFC 5321 recommends 5 minutes.
	timeout := smtp.Timeout(c.conn, c.cfg.commandTimeout)
	defer timeout()

	_, msg, err := c.readResponse(220)
	if err != nil {
		_ = c.Close()
	}

	if idx := strings.IndexRune(msg, ' '); idx >= 0 {
		msg = msg[:idx]
	}

	c.connName = msg

	return err
}

// Hello runs a hello exchange
// if an error occurred the connection is closed
func (c *Client) Hello() error {
	// verify if local name is valid
	if strings.ContainsAny(c.cfg.localName, "\n\r") {
		return errors.New("smtp: the local name must not contain CR or LF")
	}

	err := c.ehlo()

	var smtp *smtp.Status
	if err != nil && errors.As(err, &smtp) && (smtp.Code == 500 || smtp.Code == 502) {
		// The server doesn't support EHLO, fallback to HELO
		err = c.helo()
	}

	if err != nil {
		_ = c.Close()
	}

	return err
}

func (c *Client) readResponse(expectCode int) (int, string, error) {
	code, msg, err := c.cfg.text.ReadResponse(expectCode)
	if protoErr, ok := err.(*textproto.Error); ok {
		err = toSMTPErr(protoErr)
	}
	return code, msg, err
}

// cmd is a convenience function that sends a command and returns the response
// textproto.Error returned by c.text.ReadResponse is converted into smtp.
func (c *Client) cmd(expectCode int, format string, args ...any) (int, string, error) {
	timeout := smtp.Timeout(c.conn, c.cfg.commandTimeout)
	defer timeout()

	id, err := c.cfg.text.Cmd(format, args...)
	if err != nil {
		return 0, "", err
	}
	c.cfg.text.StartResponse(id)
	defer c.cfg.text.EndResponse(id)

	return c.readResponse(expectCode)
}

// helo sends the HELO greeting to the server. It should be used only when the
// server does not support ehlo.
func (c *Client) helo() error {
	c.ext = nil
	_, _, err := c.cmd(250, "HELO %s", c.cfg.localName)
	return err
}

// ehlo sends the EHLO (extended hello) greeting to the server. It
// should be the preferred greeting for servers that support it.
func (c *Client) ehlo() error {
	cmd := "EHLO"

	_, msg, err := c.cmd(250, "%s %s", cmd, c.cfg.localName)
	if err != nil {
		return err
	}
	ext := make(map[string]string)
	extList := strings.Split(msg, "\n")
	if len(extList) > 1 {
		extList = extList[1:]
		for _, line := range extList {
			i := strings.IndexByte(line, ' ')
			if i < 0 {
				ext[line] = ""
			} else {
				ext[line[:i]] = line[i+1:]
			}
		}
	}
	c.ext = ext
	return err
}

// StartTLS sends the STARTTLS command and encrypts all further communication.
// Only servers that advertise the STARTTLS extension support this function.
//
// A nil config is equivalent to a zero tls.Config.
//
// If server returns an error, it will be of type *smtp.
// if an error occurred the connection is closed
func (c *Client) StartTLS(config *tls.Config, serverName string) error {
	_, _, err := c.cmd(220, "STARTTLS")
	if err != nil {
		_ = c.Quit()
		return err
	}

	if config == nil {
		config = &tls.Config{
			ServerName: serverName,
		}
	} else if config.ServerName == "" && serverName != "" {
		// Make a copy to avoid polluting argument
		config = config.Clone()
		config.ServerName = serverName
	}

	conn := tls.Client(c.conn, config)

	timeout := smtp.Timeout(conn, c.cfg.tlsHandshakeTimeout)
	defer timeout()

	err = conn.Handshake()
	if err != nil {
		_ = c.Close()
		return err
	}

	err = c.initConn(conn, false)
	if err != nil {
		return err
	}

	return nil
}

// TLSConnectionState returns the client's TLS connection state.
// The return values are their zero values if STARTTLS did
// not succeed.
func (c *Client) TLSConnectionState() (tls.ConnectionState, bool) {
	tc, ok := c.conn.(*tls.Conn)
	if !ok {
		return tls.ConnectionState{}, ok
	}
	return tc.ConnectionState(), true
}

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Verify(addr string, opts *VrfyOptions) error {
	if err := validateLine(addr); err != nil {
		return err
	}

	var sb strings.Builder

	sb.Grow(2048)
	fmt.Fprintf(&sb, "VRFY %s", addr)

	// By default utf8 is preferred
	if opts == nil || opts.UTF8 != UTF8Disabled {
		if _, ok := c.ext["SMTPUTF8"]; ok {
			sb.WriteString(" SMTPUTF8")
		} else if opts != nil && opts.UTF8 == UTF8Force {
			return errors.New("smtp: server does not support SMTPUTF8")
		}
	}

	_, _, err := c.cmd(250, "%s", sb.String())
	return err
}

// Auth authenticates a client using the provided authentication mechanism.
// Only servers that advertise the AUTH extension support this function.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Auth(saslClient sasl.Client) error {
	if saslClient == nil {
		return errors.New("smtp: SASL client is missing")
	}

	encoding := base64.StdEncoding
	mech, resp, err := saslClient.Start()
	if err != nil {
		return err
	}
	var resp64 []byte
	if len(resp) > 0 {
		resp64 = make([]byte, encoding.EncodedLen(len(resp)))
		encoding.Encode(resp64, resp)
	} else if resp != nil {
		resp64 = []byte{'='}
	}
	code, msg64, err := c.cmd(0, "%s", strings.TrimSpace(fmt.Sprintf("AUTH %s %s", mech, resp64)))
	for err == nil {
		var msg []byte
		switch code {
		case 334:
			msg, err = encoding.DecodeString(msg64)
		case 235:
			// the last message isn't base64 because it isn't a challenge
			msg = []byte(msg64)
		default:
			err = toSMTPErr(&textproto.Error{Code: code, Msg: msg64})
		}
		if err == nil {
			if code == 334 {
				resp, err = saslClient.Next(msg)
			} else {
				resp = nil
			}
		}
		if err != nil {
			// abort the AUTH
			_, _, _ = c.cmd(501, "*")
			break
		}
		if resp == nil {
			break
		}
		resp64 = make([]byte, encoding.EncodedLen(len(resp)))
		encoding.Encode(resp64, resp)
		code, msg64, err = c.cmd(0, "%s", string(resp64))
	}
	return err
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
func (c *Client) Mail(from string, opts *MailOptions) error {
	if err := validateLine(from); err != nil {
		return err
	}

	var sb strings.Builder
	// A high enough power of 2 than 510+14+26+11+9+9+39+500
	sb.Grow(2048)
	fmt.Fprintf(&sb, "MAIL FROM:<%s>", from)
	if _, ok := c.ext["8BITMIME"]; ok {
		sb.WriteString(" BODY=8BITMIME")
	}
	if _, ok := c.ext["SIZE"]; ok && opts != nil && opts.Size != 0 {
		fmt.Fprintf(&sb, " SIZE=%v", opts.Size)
	}
	if opts != nil && opts.RequireTLS {
		if _, ok := c.ext["REQUIRETLS"]; !ok {
			return errors.New("smtp: server does not support REQUIRETLS")
		}
		sb.WriteString(" REQUIRETLS")
	}
	// By default utf8 is preferred
	if opts == nil || opts.UTF8 != UTF8Disabled {
		if _, ok := c.ext["SMTPUTF8"]; ok {
			sb.WriteString(" SMTPUTF8")
		} else if opts != nil && opts.UTF8 == UTF8Force {
			return errors.New("smtp: server does not support SMTPUTF8")
		}
	}
	if _, ok := c.ext["DSN"]; ok && opts != nil {
		switch opts.Return {
		case smtp.DSNReturnFull, smtp.DSNReturnHeaders:
			fmt.Fprintf(&sb, " RET=%s", string(opts.Return))
		case "":
			// This space is intentionally left blank
		default:
			return errors.New("smtp: Unknown RET parameter value")
		}
		if opts.EnvelopeID != "" {
			if !textsmtp.IsPrintableASCII(opts.EnvelopeID) {
				return errors.New("smtp: Malformed ENVID parameter value")
			}
			fmt.Fprintf(&sb, " ENVID=%s", encodeXtext(opts.EnvelopeID))
		}
	}
	if opts != nil && opts.Auth != nil {
		if _, ok := c.ext["AUTH"]; ok {
			fmt.Fprintf(&sb, " AUTH=%s", encodeXtext(*opts.Auth))
		}
		// We can safely discard parameter if server does not support AUTH.
	}

	if opts != nil && opts.XOORG != nil {
		if _, ok := c.ext["XOORG"]; ok {
			fmt.Fprintf(&sb, " XOORG=%s", encodeXtext(*opts.XOORG))
		}
		// We can safely discard parameter if server does not support AUTH.
	}

	_, _, err := c.cmd(250, "%s", sb.String())
	return err
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
	if err := validateLine(to); err != nil {
		return err
	}

	var sb strings.Builder
	// A high enough power of 2 than 510+29+501
	sb.Grow(2048)
	fmt.Fprintf(&sb, "RCPT TO:<%s>", to)
	if _, ok := c.ext["DSN"]; ok && opts != nil {
		if len(opts.Notify) != 0 {
			sb.WriteString(" NOTIFY=")
			if err := textsmtp.CheckNotifySet(opts.Notify); err != nil {
				return errors.New("smtp: Malformed NOTIFY parameter value")
			}
			for i, v := range opts.Notify {
				if i != 0 {
					sb.WriteString(",")
				}
				sb.WriteString(string(v))
			}
		}
		if opts.OriginalRecipient != "" {
			var enc string
			switch opts.OriginalRecipientType {
			case smtp.DSNAddressTypeRFC822:
				if !textsmtp.IsPrintableASCII(opts.OriginalRecipient) {
					return errors.New("smtp: Illegal address")
				}
				enc = encodeXtext(opts.OriginalRecipient)
			case smtp.DSNAddressTypeUTF8:
				if _, ok := c.ext["SMTPUTF8"]; ok {
					enc = encodeUTF8AddrUnitext(opts.OriginalRecipient)
				} else {
					enc = encodeUTF8AddrXtext(opts.OriginalRecipient)
				}
			default:
				return errors.New("smtp: Unknown address type")
			}
			fmt.Fprintf(&sb, " ORCPT=%s;%s", string(opts.OriginalRecipientType), enc)
		}
	}
	if _, _, err := c.cmd(25, "%s", sb.String()); err != nil {
		return err
	}
	return nil
}

// Content issues a DATA or BDAT (prefer BDAT if available) command to
// the server and returns a writer that
// can be used to write the mail headers and body. The caller should
// close the writer before calling any more methods on c. A call to
// Data must be preceded by one or more calls to Rcpt.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Content(size int) (*DataCloser, error) {
	if _, ok := c.ext["CHUNKING"]; c.cfg.chunkingMaxSize >= 0 && ok {
		return c.Bdat(size)
	}
	return c.Data()
}

// Data issues a DATA command to the server and returns a writer that
// can be used to write the mail headers and body. The caller should
// close the writer before calling any more methods on c. A call to
// Data must be preceded by one or more calls to Rcpt.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Data() (*DataCloser, error) {
	_, _, err := c.cmd(354, "DATA")
	if err != nil {
		return nil, err
	}
	return &DataCloser{c: c, writer: textsmtp.NewDotWriter(c.cfg.text.W)}, nil
}

// Bdat issues a BDAT command to the server and returns a writer that
// can be used to write the mail headers and body. The caller should
// close the writer before calling any more methods on c. A call to
// Data must be preceded by one or more calls to Rcpt.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Bdat(size int) (*DataCloser, error) {
	if c.cfg.chunkingMaxSize < 0 {
		return nil, errors.New("smtp: chunking is disabled on the client by negative chunking max size)")
	}
	if _, ok := c.ext["CHUNKING"]; !ok {
		return nil, errors.New("smtp: server doesn't support chunking")
	}

	// if chunking max size is active but smaller than a typically []byte write call, the buffer is just overhead
	if c.cfg.chunkingBufferEnabled && size == 0 && (c.cfg.chunkingMaxSize == 0 || c.cfg.chunkingMaxSize > 4096) {
		// c.bdatBuffer is init on first use and always reuse it
		bufferSize := defaultChunkingMaxSize
		if c.cfg.chunkingMaxSize > 0 {
			bufferSize = c.cfg.chunkingMaxSize
		}
		if len(c.chunkingBuffer) < bufferSize {
			c.chunkingBuffer = make([]byte, bufferSize)
		}

		return &DataCloser{c: c, writer: textsmtp.NewBdatWriterBuffered(c.cfg.chunkingMaxSize, c.cfg.text.W, func() error {
			_, _, err := c.cfg.text.ReadResponse(250)
			return err
		}, size, c.chunkingBuffer[:bufferSize])}, nil
	}

	return &DataCloser{c: c, writer: textsmtp.NewBdatWriter(c.cfg.chunkingMaxSize, c.cfg.text.W, func() error {
		_, _, err := c.cfg.text.ReadResponse(250)
		return err
	}, size)}, nil
}

// Extension reports whether an extension is support by the server.
// The extension name is case-insensitive. If the extension is supported,
// Extension also returns a string that contains any parameters the
// server specifies for the extension.
func (c *Client) Extension(ext string) (bool, string) {
	ext = strings.ToUpper(ext)
	param, ok := c.ext[ext]
	return ok, param
}

// SupportsAuth checks whether an authentication mechanism is supported.
func (c *Client) SupportsAuth(mech string) bool {
	mechs, ok := c.ext["AUTH"]
	if !ok {
		return false
	}
	for m := range strings.SplitSeq(mechs, " ") {
		if strings.EqualFold(m, mech) {
			return true
		}
	}
	return false
}

// MaxMessageSize returns the maximum message size accepted by the server.
// 0 means unlimited.
//
// If the server doesn't convey this information, ok = false is returned.
func (c *Client) MaxMessageSize() (size int, ok bool) {
	v := c.ext["SIZE"]
	if v == "" {
		return 0, false
	}
	size, err := strconv.Atoi(v)
	if err != nil || size < 0 {
		return 0, false
	}
	return size, true
}

// Reset sends the RSET command to the server, aborting the current mail
// transaction.
func (c *Client) Reset() error {
	if _, _, err := c.cmd(250, "RSET"); err != nil {
		return err
	}
	return nil
}

// Noop sends the NOOP command to the server. It does nothing but check
// that the connection to the server is okay.
func (c *Client) Noop() error {
	_, _, err := c.cmd(250, "NOOP")
	return err
}

// Quit sends the QUIT command and closes the connection to the server.
// If Quit fails the connection will still be closed.
func (c *Client) Quit() error {
	if c.conn == nil {
		return nil
	}
	_, _, err := c.cmd(221, "QUIT")
	if err != nil {
		_ = c.Close()
		return err
	}
	return c.Close()
}

// Connected returns the current server name.
func (c *Client) Connected() bool {
	return c.conn != nil
}

// ServerAddress returns the current server address.
func (c *Client) ServerAddress() string {
	return c.connAddress
}

// ServerName returns the current server name.
func (c *Client) ServerName() string {
	return c.connName
}
