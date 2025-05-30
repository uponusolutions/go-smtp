// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

// A Client represents a client connection to an SMTP server.
type Client struct {
	// keep a reference to the connection so it can be used to create a TLS
	// connection later
	conn       net.Conn
	text       *textsmtp.Conn
	serverName string
	lmtp       bool
	ext        map[string]string // supported extensions
	localName  string            // the name to use in HELO/EHLO/LHLO
	didGreet   bool              // whether we've received greeting from server
	greetError error             // the error from the greeting
	didHello   bool              // whether we've said HELO/EHLO/LHLO
	helloError error             // the error from the hello
	rcpts      []string          // recipients accumulated for the current session

	// Time to wait for command responses (this includes 3xx reply to DATA).
	TLSHandshakeTimeout time.Duration

	// Time to wait for command responses (this includes 3xx reply to DATA).
	CommandTimeout time.Duration
	// Time to wait for responses after final dot.
	SubmissionTimeout time.Duration

	// Logger for all network activity.
	Debug io.Writer
}

// 30 seconds was chosen as it's the same duration as http.DefaultTransport's
// timeout.
var defaultDialer = net.Dialer{Timeout: 30 * time.Second}

// Dial returns a new Client connected to an SMTP server at addr. The addr must
// include a port, as in "mail.example.com:smtp".
//
// This function returns a plaintext connection. To enable TLS, use
// DialStartTLS.
func Dial(addr string) (*Client, error) {
	conn, err := defaultDialer.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	client := NewClient(conn)
	client.serverName, _, _ = net.SplitHostPort(addr)
	return client, nil
}

// DialTLS returns a new Client connected to an SMTP server via TLS at addr.
// The addr must include a port, as in "mail.example.com:smtps".
//
// A nil tlsConfig is equivalent to a zero tls.Config.
func DialTLS(addr string, tlsConfig *tls.Config) (*Client, error) {
	tlsDialer := tls.Dialer{
		NetDialer: &defaultDialer,
		Config:    tlsConfig,
	}
	conn, err := tlsDialer.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	client := NewClient(conn)
	client.serverName, _, _ = net.SplitHostPort(addr)
	return client, nil
}

// DialStartTLS retruns a new Client connected to an SMTP server via STARTTLS
// at addr. The addr must include a port, as in "mail.example.com:smtp".
//
// A nil tlsConfig is equivalent to a zero tls.Config.
func DialStartTLS(addr string, tlsConfig *tls.Config) (*Client, error) {
	c, err := Dial(addr)
	if err != nil {
		return nil, err
	}
	if err := initStartTLS(c, tlsConfig); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}

// NewClient returns a new Client using an existing connection and host as a
// server name to be used when authenticating.
func NewClient(conn net.Conn) *Client {
	c := &Client{
		localName: "localhost",
		// As recommended by RFC 5321. For DATA command reply (3xx one) RFC
		// recommends a slightly shorter timeout but we do not bother
		// differentiating these.
		CommandTimeout: 5 * time.Minute,
		// 10 minutes + 2 minute buffer in case the server is doing transparent
		// forwarding and also follows recommended timeouts.
		SubmissionTimeout: 12 * time.Minute,
		// 30 seconds, very generous
		TLSHandshakeTimeout: 30 * time.Second,
	}

	c.setConn(conn)

	return c
}

// NewClientStartTLS creates a new Client and performs a STARTTLS command.
func NewClientStartTLS(conn net.Conn, tlsConfig *tls.Config) (*Client, error) {
	c := NewClient(conn)
	if err := initStartTLS(c, tlsConfig); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}

func initStartTLS(c *Client, tlsConfig *tls.Config) error {
	if err := c.hello(); err != nil {
		return err
	}
	if ok, _ := c.Extension("STARTTLS"); !ok {
		return errors.New("smtp: server doesn't support STARTTLS")
	}
	if err := c.startTLS(tlsConfig); err != nil {
		return err
	}
	return nil
}

// NewClientLMTP returns a new LMTP Client (as defined in RFC 2033) using an
// existing connection and host as a server name to be used when authenticating.
func NewClientLMTP(conn net.Conn) *Client {
	c := NewClient(conn)
	c.lmtp = true
	return c
}

// setConn sets the underlying network connection for the client.
func (c *Client) setConn(conn net.Conn) {
	// Doubled maximum line length per RFC 5321 (Section 4.5.3.1.6)
	maxLineLength := 2000
	readerSize := 4096
	writerSize := 4096

	c.conn = conn

	if c.Debug != nil {
		c.text = textsmtp.NewConn(struct {
			io.Reader
			io.Writer
			io.Closer
		}{
			io.TeeReader(c.conn, c.Debug),
			io.MultiWriter(c.conn, c.Debug),
			c.conn,
		}, readerSize, writerSize, maxLineLength)
	}
	if c.text != nil {
		c.text.Replace(conn)
	} else {
		c.text = textsmtp.NewConn(conn, readerSize, writerSize, maxLineLength)
	}
}

// Close closes the connection.
func (c *Client) Close() error {
	return c.text.Close()
}

func (c *Client) greet() error {
	if c.didGreet {
		return c.greetError
	}

	// Initial greeting timeout. RFC 5321 recommends 5 minutes.
	c.conn.SetDeadline(time.Now().Add(c.CommandTimeout))
	defer c.conn.SetDeadline(time.Time{})

	c.didGreet = true
	_, _, err := c.readResponse(220)
	if err != nil {
		c.greetError = err
		c.text.Close()
	}

	return c.greetError
}

// hello runs a hello exchange if needed.
func (c *Client) hello() error {
	if c.didHello {
		return c.helloError
	}

	if err := c.greet(); err != nil {
		return err
	}

	c.didHello = true
	if err := c.ehlo(); err != nil {
		var smtp *smtp.SMTPStatus
		if errors.As(err, &smtp) && (smtp.Code == 500 || smtp.Code == 502) {
			// The server doesn't support EHLO, fallback to HELO
			c.helloError = c.helo()
		} else {
			c.helloError = err
		}
	}
	return c.helloError
}

// Hello sends a HELO or EHLO to the server as the given host name.
// Calling this method is only necessary if the client needs control
// over the host name used. The client will introduce itself as "localhost"
// automatically otherwise. If Hello is called, it must be called before
// any of the other methods.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Hello(localName string) error {
	if err := validateLine(localName); err != nil {
		return err
	}
	if c.didHello {
		return errors.New("smtp: Hello called after other methods")
	}
	c.localName = localName
	return c.hello()
}

func (c *Client) readResponse(expectCode int) (int, string, error) {
	code, msg, err := c.text.ReadResponse(expectCode)
	if protoErr, ok := err.(*textproto.Error); ok {
		err = toSMTPErr(protoErr)
	}
	return code, msg, err
}

// cmd is a convenience function that sends a command and returns the response
// textproto.Error returned by c.text.ReadResponse is converted into smtp.
func (c *Client) cmd(expectCode int, format string, args ...any) (int, string, error) {
	c.conn.SetDeadline(time.Now().Add(c.CommandTimeout))
	defer c.conn.SetDeadline(time.Time{})

	id, err := c.text.Cmd(format, args...)
	if err != nil {
		return 0, "", err
	}
	c.text.StartResponse(id)
	defer c.text.EndResponse(id)

	return c.readResponse(expectCode)
}

// helo sends the HELO greeting to the server. It should be used only when the
// server does not support ehlo.
func (c *Client) helo() error {
	c.ext = nil
	_, _, err := c.cmd(250, "HELO %s", c.localName)
	return err
}

// ehlo sends the EHLO (extended hello) greeting to the server. It
// should be the preferred greeting for servers that support it.
func (c *Client) ehlo() error {
	cmd := "EHLO"
	if c.lmtp {
		cmd = "LHLO"
	}

	_, msg, err := c.cmd(250, "%s %s", cmd, c.localName)
	if err != nil {
		return err
	}
	ext := make(map[string]string)
	extList := strings.Split(msg, "\n")
	if len(extList) > 1 {
		extList = extList[1:]
		for _, line := range extList {
			args := strings.SplitN(line, " ", 2)
			if len(args) > 1 {
				ext[args[0]] = args[1]
			} else {
				ext[args[0]] = ""
			}
		}
	}
	c.ext = ext
	return err
}

// startTLS sends the STARTTLS command and encrypts all further communication.
// Only servers that advertise the STARTTLS extension support this function.
//
// A nil config is equivalent to a zero tls.Config.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) startTLS(config *tls.Config) error {
	if err := c.hello(); err != nil {
		return err
	}
	_, _, err := c.cmd(220, "STARTTLS")
	if err != nil {
		return err
	}
	if config == nil {
		config = &tls.Config{}
	}
	if config.ServerName == "" && c.serverName != "" {
		// Make a copy to avoid polluting argument
		config = config.Clone()
		config.ServerName = c.serverName
	}
	if testHookStartTLS != nil {
		testHookStartTLS(config)
	}

	conn := tls.Client(c.conn, config)

	conn.SetDeadline(time.Now().Add(c.TLSHandshakeTimeout))
	defer conn.SetDeadline(time.Time{})

	err = conn.Handshake()
	if err != nil {
		return err
	}

	c.setConn(conn)
	c.didHello = false
	return nil
}

// TLSConnectionState returns the client's TLS connection state.
// The return values are their zero values if STARTTLS did
// not succeed.
func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool) {
	tc, ok := c.conn.(*tls.Conn)
	if !ok {
		return
	}
	return tc.ConnectionState(), true
}

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
//
// If server returns an error, it will be of type *smtp.
func (c *Client) Verify(addr string, opts *smtp.VrfyOptions) error {
	if err := validateLine(addr); err != nil {
		return err
	}
	if err := c.hello(); err != nil {
		return err
	}

	var sb strings.Builder

	sb.Grow(2048)
	fmt.Fprintf(&sb, "VRFY %s", addr)

	if opts != nil && opts.UTF8 {
		if _, ok := c.ext["SMTPUTF8"]; ok {
			sb.WriteString(" SMTPUTF8")
		} else {
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
func (c *Client) Auth(a sasl.Client) error {
	if err := c.hello(); err != nil {
		return err
	}
	encoding := base64.StdEncoding
	mech, resp, err := a.Start()
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
				resp, err = a.Next(msg)
			} else {
				resp = nil
			}
		}
		if err != nil {
			// abort the AUTH
			c.cmd(501, "*")
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
func (c *Client) Mail(from string, opts *smtp.MailOptions) error {
	if err := validateLine(from); err != nil {
		return err
	}
	if err := c.hello(); err != nil {
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
		if _, ok := c.ext["REQUIRETLS"]; ok {
			sb.WriteString(" REQUIRETLS")
		} else {
			return errors.New("smtp: server does not support REQUIRETLS")
		}
	}
	if opts != nil && opts.UTF8 {
		if _, ok := c.ext["SMTPUTF8"]; ok {
			sb.WriteString(" SMTPUTF8")
		} else {
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
		if opts.Notify != nil && len(opts.Notify) != 0 {
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
	c.rcpts = append(c.rcpts, to)
	return nil
}

// DataCloser implement an io.WriteCloser with the additional
// CloseWithResponse function.
type DataCloser struct {
	c *Client
	io.WriteCloser
	statusCb func(rcpt string, status *smtp.SMTPStatus)
	closed   bool
}

func (d *DataCloser) CloseWithResponse() (code int, msg string, err error) {
	if d.closed {
		return 0, "", fmt.Errorf("smtp: data writer closed twice")
	}

	if err := d.WriteCloser.Close(); err != nil {
		return 0, "", err
	}

	d.c.conn.SetDeadline(time.Now().Add(d.c.SubmissionTimeout))
	defer d.c.conn.SetDeadline(time.Time{})

	expectedResponses := len(d.c.rcpts)
	if d.c.lmtp {
		for expectedResponses > 0 {
			rcpt := d.c.rcpts[len(d.c.rcpts)-expectedResponses]
			if _, _, err := d.c.readResponse(250); err != nil {
				if smtpErr, ok := err.(*smtp.SMTPStatus); ok {
					if d.statusCb != nil {
						d.statusCb(rcpt, smtpErr)
					}
				} else {
					return 0, "", err
				}
			} else if d.statusCb != nil {
				d.statusCb(rcpt, nil)
			}
			expectedResponses--
		}
	} else {
		code, msg, err = d.c.readResponse(250)
	}

	d.closed = true
	return code, msg, err
}

func (d *DataCloser) Close() error {
	_, _, err := d.CloseWithResponse()
	return err
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
	return &DataCloser{c: c, WriteCloser: textsmtp.NewDotWriter(c.text.W)}, nil
}

// LMTPData is the LMTP-specific version of the Data method. It accepts a callback
// that will be called for each status response received from the server.
//
// Status callback will receive a smtp argument for each negative server
// reply and nil for each positive reply. I/O errors will not be reported using
// callback and instead will be returned by the Close method of DataCloser.
// Callback will be called for each successfull Rcpt call done before in the
// same order.
func (c *Client) LMTPData(statusCb func(rcpt string, status *smtp.SMTPStatus)) (*DataCloser, error) {
	if !c.lmtp {
		return nil, errors.New("smtp: not a LMTP client")
	}

	_, _, err := c.cmd(354, "DATA")
	if err != nil {
		return nil, err
	}
	return &DataCloser{c: c, WriteCloser: textsmtp.NewDotWriter(c.text.W), statusCb: statusCb}, nil
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
func (c *Client) SendMail(from string, to []string, r io.Reader) error {
	var err error

	if err = c.Mail(from, nil); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr, nil); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}
	return w.Close()
}

var testHookStartTLS func(*tls.Config) // nil, except for tests

func sendMail(addr string, implicitTLS bool, a sasl.Client, from string, to []string, r io.Reader) error {
	if err := validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err := validateLine(recp); err != nil {
			return err
		}
	}

	var (
		c   *Client
		err error
	)
	if implicitTLS {
		c, err = DialTLS(addr, nil)
	} else {
		c, err = DialStartTLS(addr, nil)
	}
	if err != nil {
		return err
	}
	defer c.Close()

	if a != nil {
		if ok, _ := c.Extension("AUTH"); !ok {
			return errors.New("smtp: server doesn't support AUTH")
		}
		if err = c.Auth(a); err != nil {
			return err
		}
	}

	if err := c.SendMail(from, to, r); err != nil {
		return err
	}

	return c.Quit()
}

// SendMail connects to the server at addr, switches to TLS, authenticates with
// the optional SASL client, and then sends an email from address from, to
// addresses to, with message r. The addr must include a port, as in
// "mail.example.com:smtp".
//
// The addresses in the to parameter are the SMTP RCPT addresses.
//
// The r parameter should be an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of r
// should be CRLF terminated. The r headers should usually include
// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
// messages is accomplished by including an email address in the to
// parameter but not including it in the r headers.
//
// SendMail is intended to be used for very simple use-cases. If you want to
// customize SendMail's behavior, use a Client instead.
//
// The SendMail function and the go-smtp package are low-level
// mechanisms and provide no support for DKIM signing (see go-msgauth), MIME
// attachments (see the mime/multipart package or the go-message package), or
// other mail functionality.
func SendMail(addr string, a sasl.Client, from string, to []string, r io.Reader) error {
	return sendMail(addr, false, a, from, to, r)
}

// SendMailTLS works like SendMail, but with implicit TLS.
func SendMailTLS(addr string, a sasl.Client, from string, to []string, r io.Reader) error {
	return sendMail(addr, true, a, from, to, r)
}

// Extension reports whether an extension is support by the server.
// The extension name is case-insensitive. If the extension is supported,
// Extension also returns a string that contains any parameters the
// server specifies for the extension.
func (c *Client) Extension(ext string) (bool, string) {
	if err := c.hello(); err != nil {
		return false, ""
	}
	ext = strings.ToUpper(ext)
	param, ok := c.ext[ext]
	return ok, param
}

// SupportsAuth checks whether an authentication mechanism is supported.
func (c *Client) SupportsAuth(mech string) bool {
	if err := c.hello(); err != nil {
		return false
	}
	mechs, ok := c.ext["AUTH"]
	if !ok {
		return false
	}
	for _, m := range strings.Split(mechs, " ") {
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
	if err := c.hello(); err != nil {
		return 0, false
	}
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
	if err := c.hello(); err != nil {
		return err
	}
	if _, _, err := c.cmd(250, "RSET"); err != nil {
		return err
	}

	c.rcpts = nil
	return nil
}

// Noop sends the NOOP command to the server. It does nothing but check
// that the connection to the server is okay.
func (c *Client) Noop() error {
	if err := c.hello(); err != nil {
		return err
	}
	_, _, err := c.cmd(250, "NOOP")
	return err
}

// Quit sends the QUIT command and closes the connection to the server.
//
// If Quit fails the connection is not closed, Close should be used
// in this case.
func (c *Client) Quit() error {
	if err := c.hello(); err != nil {
		return err
	}
	_, _, err := c.cmd(221, "QUIT")
	if err != nil {
		return err
	}
	return c.Close()
}

func parseEnhancedCode(s string) (smtp.EnhancedCode, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return smtp.EnhancedCode{}, fmt.Errorf("wrong amount of enhanced code parts")
	}

	code := smtp.EnhancedCode{}
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return code, err
		}
		code[i] = num
	}
	return code, nil
}

// toSMTPErr converts textproto.Error into smtp, parsing
// enhanced status code if it is present.
func toSMTPErr(protoErr *textproto.Error) *smtp.SMTPStatus {
	smtpErr := &smtp.SMTPStatus{
		Code:    protoErr.Code,
		Message: protoErr.Msg,
	}

	parts := strings.SplitN(protoErr.Msg, " ", 2)
	if len(parts) != 2 {
		return smtpErr
	}

	enchCode, err := parseEnhancedCode(parts[0])
	if err != nil {
		return smtpErr
	}

	msg := parts[1]

	// Per RFC 2034, enhanced code should be prepended to each line.
	msg = strings.ReplaceAll(msg, "\n"+parts[0]+" ", "\n")

	smtpErr.EnhancedCode = enchCode
	smtpErr.Message = msg
	return smtpErr
}

// validateLine checks to see if a line has CR or LF.
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("smtp: a line must not contain CR or LF")
	}
	return nil
}

func encodeXtext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		switch {
		case ch >= '!' && ch <= '~' && ch != '+' && ch != '=':
			// printable non-space US-ASCII except '+' and '='
			out.WriteRune(ch)
		default:
			out.WriteRune('+')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
		}
	}
	return out.String()
}

// encodeUTF8AddrUnitext encodes raw string to the utf-8-addr-unitext form in RFC 6533.
func encodeUTF8AddrUnitext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		switch {
		case ch >= '!' && ch <= '~' && ch != '+' && ch != '=':
			// printable non-space US-ASCII except '+' and '='
			out.WriteRune(ch)
		case ch <= '\x7F':
			// other ASCII: CTLs, space and specials
			out.WriteRune('\\')
			out.WriteRune('x')
			out.WriteRune('{')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
			out.WriteRune('}')
		default:
			// UTF-8 non-ASCII
			out.WriteRune(ch)
		}
	}
	return out.String()
}

// encodeUTF8AddrXtext encodes raw string to the utf-8-addr-xtext form in RFC 6533.
func encodeUTF8AddrXtext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		switch {
		case ch >= '!' && ch <= '~' && ch != '+' && ch != '=':
			// printable non-space US-ASCII except '+' and '='
			out.WriteRune(ch)
		default:
			out.WriteRune('\\')
			out.WriteRune('x')
			out.WriteRune('{')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
			out.WriteRune('}')
		}
	}
	return out.String()
}
