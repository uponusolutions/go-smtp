package server

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/emersion/go-sasl"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/parse"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

type Conn struct {
	conn net.Conn

	text   *textsmtp.Conn
	server *Server
	helo   string

	session    Session
	locker     sync.Mutex
	binarymime bool

	from       *string
	recipients []string
	didAuth    bool
}

func newConn(c net.Conn, s *Server) *Conn {
	sc := &Conn{
		server: s,
		conn:   c,
	}

	sc.init()
	return sc
}

func (c *Conn) init() {
	if c.server.Debug != nil {
		c.text = textsmtp.NewConn(struct {
			io.Reader
			io.Writer
			io.Closer
		}{
			io.TeeReader(c.conn, c.server.Debug),
			io.MultiWriter(c.conn, c.server.Debug),
			c.conn,
		}, c.server.ReaderSize, c.server.WriterSize, c.server.MaxLineLength)
	}

	c.text = textsmtp.NewConn(c.conn, c.server.ReaderSize, c.server.WriterSize, c.server.MaxLineLength)
}

func (c *Conn) nextCommand() (string, string, error) {
	line, err := c.readLine()
	if err != nil {
		return "", "", err
	}
	return parse.Cmd(line)
}

// Commands are dispatched to the appropriate handler functions.
func (c *Conn) handle(cmd string, arg string) error {
	if cmd == "" {
		return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 2}, "Error: bad syntax")
	}

	cmd = strings.ToUpper(cmd)
	switch cmd {
	case "SEND", "SOML", "SAML", "EXPN", "HELP", "TURN":
		// These commands are not implemented in any state
		c.writeResponse(502, smtp.EnhancedCode{5, 5, 1}, fmt.Sprintf("%v command not implemented", cmd))
	case "HELO", "EHLO", "LHLO":
		lmtp := cmd == "LHLO"
		enhanced := lmtp || cmd == "EHLO"
		if lmtp {
			return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 1}, "This is not a LMTP server")
		}
		return c.handleGreet(enhanced, arg)
	case "MAIL":
		return c.handleMail(arg)
	case "RCPT":
		return c.handleRcpt(arg)
	case "VRFY":
		c.writeResponse(252, smtp.EnhancedCode{2, 5, 0}, "Cannot VRFY user, but will accept message")
	case "NOOP":
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "I have successfully done nothing")
	case "RSET": // Reset session
		c.reset()
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "Session reset")
	case "BDAT":
		return c.handleBdat(arg)
	case "DATA":
		return c.handleData(arg)
	case "QUIT":
		return smtp.Quit
	case "AUTH":
		return c.handleAuth(arg)
	case "STARTTLS":
		return c.handleStartTLS()
	default:
		msg := fmt.Sprintf("Syntax errors, %v command unrecognized", cmd)
		return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 2}, msg)
	}
	return nil
}

func (c *Conn) Server() *Server {
	return c.server
}

func (c *Conn) Session() Session {
	c.locker.Lock()
	defer c.locker.Unlock()
	return c.session
}

func (c *Conn) setSession(session Session) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.session = session
}

func (c *Conn) Close() error {
	c.locker.Lock()
	defer c.locker.Unlock()

	if c.session != nil {
		c.session.Logout()
		c.session = nil
	}

	return c.conn.Close()
}

// TLSConnectionState returns the connection's TLS connection state.
// Zero values are returned if the connection doesn't use TLS.
func (c *Conn) TLSConnectionState() (state tls.ConnectionState, ok bool) {
	tc, ok := c.conn.(*tls.Conn)
	if !ok {
		return
	}
	return tc.ConnectionState(), true
}

func (c *Conn) Hostname() string {
	return c.helo
}

func (c *Conn) Conn() net.Conn {
	return c.conn
}

func (c *Conn) authAllowed() bool {
	_, isTLS := c.TLSConnectionState()
	return isTLS || c.server.AllowInsecureAuth
}

// GREET state -> waiting for HELO
func (c *Conn) handleGreet(enhanced bool, arg string) error {
	domain, err := parse.HelloArgument(arg)
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Domain/address argument required for HELO")
	}
	// c.helo is populated before NewSession so
	// NewSession can access it via Conn.Hostname.
	c.helo = domain

	// RFC 5321: "An EHLO command MAY be issued by a client later in the session"
	if c.session != nil {
		// RFC 5321: "... the SMTP server MUST clear all buffers
		// and reset the state exactly as if a RSET command has been issued."
		c.reset()
	} else {
		sess, err := c.server.Backend.NewSession(c)
		if err != nil {
			c.helo = ""
			return c.newStatusError(451, smtp.EnhancedCode{4, 0, 0}, err)
		}

		c.setSession(sess)
	}

	if !enhanced {
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("Hello %s", domain))
		return nil
	}

	caps := []string{
		"PIPELINING",
		"8BITMIME",
		"ENHANCEDSTATUSCODES",
		"CHUNKING",
	}
	if _, isTLS := c.TLSConnectionState(); c.server.TLSConfig != nil && !isTLS {
		caps = append(caps, "STARTTLS")
	}
	if c.authAllowed() {
		mechs := c.authMechanisms()

		authCap := "AUTH"
		for _, name := range mechs {
			authCap += " " + name
		}

		if len(mechs) > 0 {
			caps = append(caps, authCap)
		}
	}
	if c.server.EnableSMTPUTF8 {
		caps = append(caps, "SMTPUTF8")
	}
	if _, isTLS := c.TLSConnectionState(); isTLS && c.server.EnableREQUIRETLS {
		caps = append(caps, "REQUIRETLS")
	}
	if c.server.EnableBINARYMIME {
		caps = append(caps, "BINARYMIME")
	}
	if c.server.EnableDSN {
		caps = append(caps, "DSN")
	}
	if c.server.EnableXOORG {
		caps = append(caps, "XOORG")
	}
	if c.server.MaxMessageBytes > 0 {
		caps = append(caps, fmt.Sprintf("SIZE %v", c.server.MaxMessageBytes))
	} else {
		caps = append(caps, "SIZE")
	}
	if c.server.MaxRecipients > 0 {
		caps = append(caps, fmt.Sprintf("LIMITS RCPTMAX=%v", c.server.MaxRecipients))
	}

	args := []string{"Hello " + domain}
	args = append(args, caps...)
	c.writeResponse(250, smtp.NoEnhancedCode, args...)
	return nil
}

func (c *Conn) handleError(err error) {
	if err == io.EOF || errors.Is(err, net.ErrClosed) {
		return
	}

	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
		c.writeResponse(421, smtp.EnhancedCode{4, 4, 2}, "Idle timeout, bye bye")
		return
	}

	if smtpErr, ok := err.(*smtp.SMTPStatus); ok {
		c.writeResponse(smtpErr.Code, smtpErr.EnhancedCode, smtpErr.Message)
		return
	}

	if err == textsmtp.ErrTooLongLine {
		c.writeResponse(500, smtp.EnhancedCode{5, 4, 0}, "Too long line, closing connection")
		return
	}

	c.writeStatus(smtp.ErrConnection)
}

// READY state -> waiting for MAIL
func (c *Conn) handleMail(arg string) error {
	if c.helo == "" {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Please introduce yourself first.")
	}
	if c.from != nil {
		return smtp.NewStatus(503, smtp.EnhancedCode{5, 5, 1}, "Already received FROM:")
	}

	arg, ok := parse.CutPrefixFold(arg, "FROM:")
	if !ok {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Was expecting MAIL arg syntax of FROM:<address>")
	}

	p := parse.Parser{S: strings.TrimSpace(arg)}
	from, err := p.ReversePath()
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Was expecting MAIL arg syntax of FROM:<address>")
	}
	args, err := parse.Args(p.S)
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unable to parse MAIL ESMTP parameters")
	}

	opts := &smtp.MailOptions{}

	c.binarymime = false
	// This is where the Conn may put BODY=8BITMIME, but we already
	// read the DATA as bytes, so it does not effect our processing.
	for key, value := range args {
		switch key {
		case "SIZE":
			size, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unable to parse SIZE as an integer")
			}

			if c.server.MaxMessageBytes > 0 && int64(size) > c.server.MaxMessageBytes {
				return smtp.NewStatus(552, smtp.EnhancedCode{5, 3, 4}, "Max message size exceeded")
			}

			opts.Size = int64(size)
		case "XOORG":
			if !c.server.EnableXOORG {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "EnableXOORG is not implemented")
			}
			opts.XOORG = value
		case "SMTPUTF8":
			if !c.server.EnableSMTPUTF8 {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "SMTPUTF8 is not implemented")
			}
			opts.UTF8 = true
		case "REQUIRETLS":
			if !c.server.EnableREQUIRETLS {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "REQUIRETLS is not implemented")
			}
			opts.RequireTLS = true
		case "BODY":
			value = strings.ToUpper(value)
			switch smtp.BodyType(value) {
			case smtp.BodyBinaryMIME:
				if !c.server.EnableBINARYMIME {
					return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "BINARYMIME is not implemented")
				}
				c.binarymime = true
			case smtp.Body7Bit, smtp.Body8BitMIME:
				// This space is intentionally left blank
			default:
				return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unknown BODY value")
			}
			opts.Body = smtp.BodyType(value)
		case "RET":
			if !c.server.EnableDSN {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "RET is not implemented")
			}
			value = strings.ToUpper(value)
			switch smtp.DSNReturn(value) {
			case smtp.DSNReturnFull, smtp.DSNReturnHeaders:
				// This space is intentionally left blank
			default:
				return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unknown RET value")
			}
			opts.Return = smtp.DSNReturn(value)
		case "ENVID":
			if !c.server.EnableDSN {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "ENVID is not implemented")
			}
			value, err := decodeXtext(value)
			if err != nil || value == "" || !textsmtp.IsPrintableASCII(value) {
				return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Malformed ENVID parameter value")
			}
			opts.EnvelopeID = value
		case "AUTH":
			value, err := decodeXtext(value)
			if err != nil || value == "" {
				return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 4}, "Malformed AUTH parameter value")
			}
			if value == "<>" {
				value = ""
			} else {
				p := parse.Parser{S: value}
				value, err = p.Mailbox()
				if err != nil || p.S != "" {
					return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 4}, "Malformed AUTH parameter mailbox")
				}
			}
			opts.Auth = &value
		default:
			return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 4}, "Unknown MAIL FROM argument")
		}
	}

	if err := c.Session().Mail(from, opts); err != nil {
		return c.newStatusError(451, smtp.EnhancedCode{4, 0, 0}, err)
	}

	c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("Roger, accepting mail from <%v>", from))
	c.from = &from
	return nil
}

// This regexp matches 'hexchar' token defined in
// https://tools.ietf.org/html/rfc4954#section-8 however it is intentionally
// relaxed by requiring only '+' to be present.  It allows us to detect
// malformed values such as +A or +HH and report them appropriately.
var hexcharRe = regexp.MustCompile(`\+[0-9A-F]?[0-9A-F]?`)

func decodeXtext(val string) (string, error) {
	if !strings.Contains(val, "+") {
		return val, nil
	}

	var replaceErr error
	decoded := hexcharRe.ReplaceAllStringFunc(val, func(match string) string {
		if len(match) != 3 {
			replaceErr = errors.New("incomplete hexchar")
			return ""
		}
		char, err := strconv.ParseInt(match, 16, 8)
		if err != nil {
			replaceErr = err
			return ""
		}

		return string(rune(char))
	})
	if replaceErr != nil {
		return "", replaceErr
	}

	return decoded, nil
}

// This regexp matches 'EmbeddedUnicodeChar' token defined in
// https://datatracker.ietf.org/doc/html/rfc6533.html#section-3
// however it is intentionally relaxed by requiring only '\x{HEX}' to be
// present.  It also matches disallowed characters in QCHAR and QUCHAR defined
// in above.
// So it allows us to detect malformed values and report them appropriately.
var eUOrDCharRe = regexp.MustCompile(`\\x[{][0-9A-F]+[}]|[[:cntrl:] \\+=]`)

// Decodes the utf-8-addr-xtext or the utf-8-addr-unitext form.
func decodeUTF8AddrXtext(val string) (string, error) {
	var replaceErr error
	decoded := eUOrDCharRe.ReplaceAllStringFunc(val, func(match string) string {
		if len(match) == 1 {
			replaceErr = errors.New("disallowed character:" + match)
			return ""
		}

		hexpoint := match[3 : len(match)-1]
		char, err := strconv.ParseUint(hexpoint, 16, 21)
		if err != nil {
			replaceErr = err
			return ""
		}
		switch len(hexpoint) {
		case 2:
			switch {
			// all xtext-specials
			case 0x01 <= char && char <= 0x09 ||
				0x11 <= char && char <= 0x19 ||
				char == 0x10 || char == 0x20 ||
				char == 0x2B || char == 0x3D || char == 0x7F:
			// 2-digit forms
			case char == 0x5C || 0x80 <= char && char <= 0xFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 3-digit forms
		case 3:
			switch {
			case 0x100 <= char && char <= 0xFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 4-digit forms excluding surrogate
		case 4:
			switch {
			case 0x1000 <= char && char <= 0xD7FF:
			case 0xE000 <= char && char <= 0xFFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 5-digit forms
		case 5:
			switch {
			case 0x1_0000 <= char && char <= 0xF_FFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 6-digit forms
		case 6:
			switch {
			case 0x10_0000 <= char && char <= 0x10_FFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// the other invalid forms
		default:
			replaceErr = errors.New("illegal hexpoint:" + hexpoint)
			return ""
		}

		return string(rune(char))
	})
	if replaceErr != nil {
		return "", replaceErr
	}

	return decoded, nil
}

func decodeTypedAddress(val string) (smtp.DSNAddressType, string, error) {
	tv := strings.SplitN(val, ";", 2)
	if len(tv) != 2 || tv[0] == "" || tv[1] == "" {
		return "", "", errors.New("bad address")
	}
	aType, aAddr := strings.ToUpper(tv[0]), tv[1]

	var err error
	switch smtp.DSNAddressType(aType) {
	case smtp.DSNAddressTypeRFC822:
		aAddr, err = decodeXtext(aAddr)
		if err == nil && !textsmtp.IsPrintableASCII(aAddr) {
			err = errors.New("illegal address:" + aAddr)
		}
	case smtp.DSNAddressTypeUTF8:
		aAddr, err = decodeUTF8AddrXtext(aAddr)
	default:
		err = errors.New("unknown address type:" + aType)
	}
	if err != nil {
		return "", "", err
	}

	return smtp.DSNAddressType(aType), aAddr, nil
}

// MAIL state -> waiting for RCPTs followed by DATA
func (c *Conn) handleRcpt(arg string) error {
	if c.from == nil {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Missing MAIL FROM command.")
	}

	arg, ok := parse.CutPrefixFold(arg, "TO:")
	if !ok {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Was expecting RCPT arg syntax of TO:<address>")
	}

	p := parse.Parser{S: strings.TrimSpace(arg)}
	recipient, err := p.Path()
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Was expecting RCPT arg syntax of TO:<address>")
	}

	if c.server.MaxRecipients > 0 && len(c.recipients) >= c.server.MaxRecipients {
		return smtp.NewStatus(452, smtp.EnhancedCode{4, 5, 3}, fmt.Sprintf("Maximum limit of %v recipients reached", c.server.MaxRecipients))
	}

	args, err := parse.Args(p.S)
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unable to parse RCPT ESMTP parameters")
	}

	opts := &smtp.RcptOptions{}

	for key, value := range args {
		switch key {
		case "NOTIFY":
			if !c.server.EnableDSN {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "NOTIFY is not implemented")
			}
			notify := []smtp.DSNNotify{}
			for _, val := range strings.Split(value, ",") {
				notify = append(notify, smtp.DSNNotify(strings.ToUpper(val)))
			}
			if err := textsmtp.CheckNotifySet(notify); err != nil {
				return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Malformed NOTIFY parameter value")
			}
			opts.Notify = notify
		case "ORCPT":
			if !c.server.EnableDSN {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "ORCPT is not implemented")
			}
			aType, aAddr, err := decodeTypedAddress(value)
			if err != nil || aAddr == "" {
				return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Malformed ORCPT parameter value")
			}
			opts.OriginalRecipientType = aType
			opts.OriginalRecipient = aAddr
		default:
			return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 4}, "Unknown RCPT TO argument")
		}
	}

	if err := c.Session().Rcpt(recipient, opts); err != nil {
		return c.newStatusError(451, smtp.EnhancedCode{4, 0, 0}, err)
	}
	c.recipients = append(c.recipients, recipient)
	c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("I'll make sure <%v> gets this", recipient))
	return nil
}

func (c *Conn) handleAuth(arg string) error {
	if c.helo == "" {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Please introduce yourself first.")
	}
	if c.didAuth {
		return smtp.NewStatus(503, smtp.EnhancedCode{5, 5, 1}, "Already authenticated")
	}

	parts := strings.Fields(arg)
	if len(parts) == 0 {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 4}, "Missing parameter")
	}

	if !c.authAllowed() {
		return smtp.NewStatus(523, smtp.EnhancedCode{5, 7, 10}, "TLS is required")
	}

	mechanism := strings.ToUpper(parts[0])

	// Parse client initial response if there is one
	var ir []byte
	if len(parts) > 1 {
		var err error
		ir, err = decodeSASLResponse(parts[1])
		if err != nil {
			return smtp.NewStatus(454, smtp.EnhancedCode{4, 7, 0}, "Invalid base64 data")
		}
	}

	sasl, err := c.auth(mechanism)
	if err != nil {
		return c.newStatusError(454, smtp.EnhancedCode{4, 7, 0}, err)
	}

	response := ir
	for {
		challenge, done, err := sasl.Next(response)
		if err != nil {
			return c.newStatusError(454, smtp.EnhancedCode{4, 7, 0}, err)
		}

		if done {
			break
		}

		encoded := ""
		if len(challenge) > 0 {
			encoded = base64.StdEncoding.EncodeToString(challenge)
		}
		c.writeResponse(334, smtp.NoEnhancedCode, encoded)

		encoded, err = c.readLine()
		if err != nil {
			return err
		}

		if encoded == "*" {
			// https://tools.ietf.org/html/rfc4954#page-4
			return smtp.NewStatus(501, smtp.EnhancedCode{5, 0, 0}, "Negotiation cancelled")
		}

		response, err = decodeSASLResponse(encoded)
		if err != nil {
			return smtp.NewStatus(454, smtp.EnhancedCode{4, 7, 0}, "Invalid base64 data")
		}
	}

	c.writeResponse(235, smtp.EnhancedCode{2, 0, 0}, "Authentication succeeded")

	c.didAuth = true
	return nil
}

func decodeSASLResponse(s string) ([]byte, error) {
	if s == "=" {
		return []byte{}, nil
	}
	return base64.StdEncoding.DecodeString(s)
}

func (c *Conn) authMechanisms() []string {
	if authSession, ok := c.Session().(AuthSession); ok {
		return authSession.AuthMechanisms()
	}
	return nil
}

func (c *Conn) auth(mech string) (sasl.Server, error) {
	if authSession, ok := c.Session().(AuthSession); ok {
		return authSession.Auth(mech)
	}
	return nil, smtp.ErrAuthUnknownMechanism
}

func (c *Conn) handleStartTLS() error {
	if _, isTLS := c.TLSConnectionState(); isTLS {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Already running in TLS")
	}

	if c.server.TLSConfig == nil {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "TLS not supported")
	}

	c.writeResponse(220, smtp.EnhancedCode{2, 0, 0}, "Ready to start TLS")

	// Upgrade to TLS
	tlsConn := tls.Server(c.conn, c.server.TLSConfig)

	if err := tlsConn.Handshake(); err != nil {
		return smtp.NewStatus(550, smtp.EnhancedCode{5, 0, 0}, "Handshake error")
	}

	c.conn = tlsConn
	c.init()

	// Reset all state and close the previous Session.
	// This is different from just calling reset() since we want the Backend to
	// be able to see the information about TLS connection in the
	// ConnectionState object passed to it.
	if session := c.Session(); session != nil {
		session.Logout()
		c.setSession(nil)
	}
	c.helo = ""
	c.reset()
	return nil
}

// DATA
func (c *Conn) handleData(arg string) error {
	if arg != "" {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "DATA command should not have any arguments")
	}
	if c.binarymime {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "DATA not allowed for BINARYMIME messages")
	}

	if c.from == nil || len(c.recipients) == 0 {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Missing RCPT TO command.")
	}

	var r io.Reader

	rstart := func() io.Reader {
		if r != nil {
			return r
		}
		// We have recipients, go to accept data
		c.writeResponse(354, smtp.NoEnhancedCode, "Go ahead. End your data with <CR><LF>.<CR><LF>")

		r := textsmtp.NewDotReader(c.text.R, c.server.MaxMessageBytes)
		return r
	}

	uuid, err := c.Session().Data(rstart, *c.from, c.recipients)
	if err != nil {
		return err
	}

	c.accepted(uuid)
	c.reset()
	return nil
}

func (c *Conn) handleBdat(arg string) error {
	if c.from == nil || len(c.recipients) == 0 {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Missing RCPT TO command.")
	}

	size, last, err := bdatArg(arg)
	if err != nil {
		return err
	}

	data := &bdat{
		maxMessageBytes: c.server.MaxMessageBytes,
		size:            size,
		last:            last,
		bytesReceived:   0,
		input:           c.text.R,
		nextCommand: func() (string, string, error) {
			c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "Continue")
			return c.nextCommand()
		},
	}

	uuid, err := c.Session().Data(func() io.Reader { return data }, *c.from, c.recipients)

	if err == smtp.Reset {
		c.reset()
		c.writeStatus(smtp.Reset)
		return nil
	}

	if err != nil {
		return err
	}

	c.accepted(uuid)
	c.reset()
	return nil
}

func (c *Conn) accepted(uuid string) {
	if uuid != "" {
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "OK: queued as "+uuid)
	} else {
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "OK: queued")
	}
}

func (c *Conn) Reject() {
	c.writeResponse(421, smtp.EnhancedCode{4, 4, 5}, "Too busy. Try again later.")
	c.Close()
}

func (c *Conn) greet() {
	protocol := "ESMTP"
	c.writeResponse(220, smtp.NoEnhancedCode, fmt.Sprintf("%v %s Service Ready", c.server.Domain, protocol))
}

func (c *Conn) writeStatus(status *smtp.SMTPStatus) {
	c.writeResponse(status.Code, status.EnhancedCode, status.Message)
}

func (c *Conn) writeResponse(code int, enhCode smtp.EnhancedCode, text ...string) {
	// TODO: error handling
	if c.server.WriteTimeout != 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.server.WriteTimeout))
	}

	// All responses must include an enhanced code, if it is missing - use
	// a generic code X.0.0.
	if enhCode == smtp.EnhancedCodeNotSet {
		cat := code / 100
		switch cat {
		case 2, 4, 5:
			enhCode = smtp.EnhancedCode{cat, 0, 0}
		default:
			enhCode = smtp.NoEnhancedCode
		}
	}

	for i := 0; i < len(text)-1; i++ {
		c.text.PrintfLine("%d-%v", code, text[i])
	}
	if enhCode == smtp.NoEnhancedCode {
		c.text.PrintfLine("%d %v", code, text[len(text)-1])
	} else {
		c.text.PrintfLine("%d %v.%v.%v %v", code, enhCode[0], enhCode[1], enhCode[2], text[len(text)-1])
	}
}

func (c *Conn) newStatusError(code int, enhCode smtp.EnhancedCode, err error) *smtp.SMTPStatus {
	if smtpErr, ok := err.(*smtp.SMTPStatus); ok {
		return smtpErr
	} else {
		return smtp.NewStatus(code, enhCode, err.Error())
	}
}

// Reads a line of input
func (c *Conn) readLine() (string, error) {
	if c.server.ReadTimeout != 0 {
		if err := c.conn.SetReadDeadline(time.Now().Add(c.server.ReadTimeout)); err != nil {
			return "", err
		}
	}

	return c.text.ReadLine()
}

func (c *Conn) reset() {
	c.locker.Lock()
	defer c.locker.Unlock()

	if c.session != nil {
		c.session.Reset()
	}

	c.didAuth = false
	c.from = nil
	c.recipients = nil
}
