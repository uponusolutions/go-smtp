package server

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/parse"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

const (
	StateInit = iota
	StateUpgrade
	StateEnforceAuthentication
	StateEnforceSecureConnection
	StateGreeted
	StateMail
)

type Conn struct {
	ctx context.Context

	conn net.Conn

	state int

	text   *textsmtp.Conn
	server *Server

	session    Session
	binarymime bool

	helo       string // set in helo / ehlo
	recipients int    // count recipients
	didAuth    bool
}

func (c *Conn) run() {
	c.greet()

	for {
		cmd, arg, err := c.nextCommand()
		if err != nil {
			c.handleError(err)
			return
		}

		err = c.handle(cmd, arg)
		if err != nil {
			c.handleError(err)
			return
		}
	}
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

	switch c.state {
	case StateInit:
		fallthrough
	case StateUpgrade:
		return c.handleStateInit(cmd, arg)
	case StateEnforceSecureConnection:
		return c.handleStateEnforceSecureConnection(cmd, arg)
	case StateEnforceAuthentication:
		return c.handleStateEnforceAuthentication(cmd, arg)
	case StateGreeted:
		return c.handleStateGreeted(cmd, arg)
	case StateMail:
		return c.handleStateMail(cmd, arg)
	}

	return fmt.Errorf("unsupported state %d, how?", c.state)
}

func (c *Conn) handleStateInit(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "NOOP":
		c.writeStatus(smtp.Noop)
	case "VRFY":
		c.writeStatus(smtp.VRFY)
	case "RSET": // Reset session
		return c.handleRSET()
	case "QUIT":
		return smtp.Quit
	default:
		c.writeCommandUnknown(cmd)
	}
	return nil
}

func (c *Conn) handleStateEnforceAuthentication(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "NOOP":
		c.writeStatus(smtp.Noop)
	case "VRFY":
		c.writeStatus(smtp.VRFY)
	case "RSET": // Reset session
		return c.handleRSET()
	case "QUIT":
		return smtp.Quit
	case "AUTH":
		return c.handleAuth(arg)
	default:
		c.writeResponse(530, smtp.EnhancedCode{5, 7, 0}, "Authentication required")
	}
	return nil
}

func (c *Conn) handleStateGreeted(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "MAIL":
		return c.handleMail(arg)
	case "NOOP":
		c.writeStatus(smtp.Noop)
	case "VRFY":
		c.writeStatus(smtp.VRFY)
	case "RSET": // Reset session
		return c.handleRSET()
	case "QUIT":
		return smtp.Quit
	case "AUTH":
		return c.handleAuth(arg)
	case "STARTTLS":
		return c.handleStartTLS()
	default:
		c.writeCommandUnknown(cmd)
	}
	return nil
}

func (c *Conn) handleStateMail(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "RCPT":
		return c.handleRcpt(arg)
	case "NOOP":
		c.writeStatus(smtp.Noop)
	case "VRFY":
		c.writeStatus(smtp.VRFY)
	case "RSET": // Reset session
		return c.handleRSET()
	case "BDAT":
		if !c.server.enableCHUNKING {
			return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "CHUNKING is not implemented")
		}
		return c.handleBdat(arg)
	case "DATA":
		return c.handleData(arg)
	case "QUIT":
		return smtp.Quit
	case "STARTTLS":
		return c.handleStartTLS()
	default:
		c.writeCommandUnknown(cmd)
	}
	return nil
}

func (c *Conn) handleStateEnforceSecureConnection(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "NOOP":
		c.writeStatus(smtp.Noop)
	case "VRFY":
		c.writeStatus(smtp.VRFY)
	case "STARTTLS":
		return c.handleStartTLS()
	case "QUIT":
		return smtp.Quit
	default:
		c.writeResponse(530, smtp.EnhancedCode{5, 7, 0}, "Must issue a STARTTLS command first")
	}
	return nil
}

func (c *Conn) writeCommandUnknown(cmd string) {
	c.writeStatus(smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, fmt.Sprintf("%s command unknown, state %d", cmd, c.state)))
}

func (c *Conn) Server() *Server {
	return c.server
}

func (c *Conn) Close(ctx context.Context) error {
	var sessionErr error
	if c.session != nil {
		sessionErr = c.session.Close(ctx)
		c.session = nil
	}
	return errors.Join(sessionErr, c.conn.Close())
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

// IsTLS returns if the connection is encrypted by tls.
func (c *Conn) IsTLS() bool {
	_, ok := c.conn.(*tls.Conn)
	return ok
}

func (c *Conn) Hostname() string {
	return c.helo
}

func (c *Conn) Conn() net.Conn {
	return c.conn
}

func (c *Conn) handleRSET() error {
	err := c.reset()
	if err != nil {
		return err
	}
	c.writeStatus(smtp.Reset)
	return nil
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
	// RFC 5321: "... the SMTP server MUST clear all buffers
	// and reset the state exactly as if a RSET command has been issued."
	if c.state != StateInit && c.state != StateEnforceSecureConnection && c.state != StateEnforceAuthentication {
		err := c.reset()
		if err != nil {
			return err
		}
	}

	if c.server.enforceSecureConnection && !c.IsTLS() {
		c.state = StateEnforceSecureConnection
	} else if c.server.enforceAuthentication {
		c.state = StateEnforceAuthentication
	} else {
		c.state = StateGreeted
	}

	if !enhanced {
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("Hello %s", domain))
		return nil
	}

	caps := []string{
		"PIPELINING",
		"8BITMIME",
		"ENHANCEDSTATUSCODES",
	}

	if c.server.enableCHUNKING {
		caps = append(caps, "CHUNKING")
	}

	isTLS := c.IsTLS()

	if !isTLS && c.server.tlsConfig != nil {
		caps = append(caps, "STARTTLS")
	}

	mechs := c.session.AuthMechanisms(c.ctx)
	if len(mechs) > 0 {
		authCap := "AUTH"
		for _, name := range mechs {
			authCap += " " + name
		}

		caps = append(caps, authCap)
	}

	if c.server.enableSMTPUTF8 {
		caps = append(caps, "SMTPUTF8")
	}
	if isTLS && c.server.enableREQUIRETLS {
		caps = append(caps, "REQUIRETLS")
	}
	if c.server.enableBINARYMIME {
		caps = append(caps, "BINARYMIME")
	}
	if c.server.enableDSN {
		caps = append(caps, "DSN")
	}
	if c.server.enableXOORG {
		caps = append(caps, "XOORG")
	}
	if c.server.maxMessageBytes > 0 {
		caps = append(caps, fmt.Sprintf("SIZE %v", c.server.maxMessageBytes))
	} else {
		caps = append(caps, "SIZE")
	}
	if c.server.maxRecipients > 0 {
		caps = append(caps, fmt.Sprintf("LIMITS RCPTMAX=%v", c.server.maxRecipients))
	}

	args := []string{"Hello " + domain}
	args = append(args, caps...)
	c.writeResponse(250, smtp.NoEnhancedCode, args...)
	return nil
}

func (c *Conn) handleError(err error) {
	if err == io.EOF || errors.Is(err, net.ErrClosed) {
		c.logger().ErrorContext(c.ctx, "connection closed unexpectedly", slog.Any("err", err))
		return
	}

	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
		c.logger().ErrorContext(c.ctx, "idle timeout", slog.Any("err", err))
		c.writeResponse(421, smtp.EnhancedCode{4, 4, 2}, "Idle timeout, bye bye")
		return
	}

	if smtpErr, ok := err.(*smtp.SMTPStatus); ok {
		if smtpErr.Code != 221 {
			c.logger().ErrorContext(c.ctx, "smtp error", slog.Any("err", err))
		}
		c.writeResponse(smtpErr.Code, smtpErr.EnhancedCode, smtpErr.Message)
		return
	}

	if err == textsmtp.ErrTooLongLine {
		c.logger().ErrorContext(c.ctx, "line too long")
		c.writeResponse(500, smtp.EnhancedCode{5, 4, 0}, "Too long line, closing connection")
		return
	}

	c.logger().ErrorContext(c.ctx, "line too long", slog.Any("err", err))
	c.writeStatus(smtp.ErrConnection)
}

func (c *Conn) logger() *slog.Logger {
	// Fallback if the connection couldn't be created or is already closed.
	if c.session == nil {
		return slog.Default()
	}
	logger := c.session.Logger(c.ctx)
	if logger == nil {
		return slog.Default()
	}
	return logger
}

// READY state -> waiting for MAIL
func (c *Conn) handleMail(arg string) error {
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

			if c.server.maxMessageBytes > 0 && int64(size) > c.server.maxMessageBytes {
				return smtp.NewStatus(552, smtp.EnhancedCode{5, 3, 4}, "Max message size exceeded")
			}

			opts.Size = int64(size)
		case "XOORG":
			value, err := decodeXtext(value)
			if err != nil || value == "" {
				return smtp.NewStatus(500, smtp.EnhancedCode{5, 5, 4}, "Malformed XOORG parameter value")
			}
			if !c.server.enableXOORG {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "EnableXOORG is not implemented")
			}
			opts.XOORG = &value
		case "SMTPUTF8":
			if !c.server.enableSMTPUTF8 {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "SMTPUTF8 is not implemented")
			}
			opts.UTF8 = true
		case "REQUIRETLS":
			if !c.server.enableREQUIRETLS {
				return smtp.NewStatus(504, smtp.EnhancedCode{5, 5, 4}, "REQUIRETLS is not implemented")
			}
			opts.RequireTLS = true
		case "BODY":
			value = strings.ToUpper(value)
			switch smtp.BodyType(value) {
			case smtp.BodyBinaryMIME:
				if !c.server.enableBINARYMIME {
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
			if !c.server.enableDSN {
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
			if !c.server.enableDSN {
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

	if err := c.session.Mail(c.ctx, from, opts); err != nil {
		return c.newStatusError(451, smtp.EnhancedCode{4, 0, 0}, "Mail not accepted", err)
	}

	c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("Roger, accepting mail from <%v>", from))
	c.state = StateMail
	return nil
}

// MAIL state -> waiting for RCPTs followed by DATA
func (c *Conn) handleRcpt(arg string) error {
	arg, ok := parse.CutPrefixFold(arg, "TO:")
	if !ok {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Was expecting RCPT arg syntax of TO:<address>")
	}

	p := parse.Parser{S: strings.TrimSpace(arg)}
	recipient, err := p.Path()
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 2}, "Was expecting RCPT arg syntax of TO:<address>")
	}

	if c.server.maxRecipients > 0 && c.recipients >= c.server.maxRecipients {
		return smtp.NewStatus(452, smtp.EnhancedCode{4, 5, 3}, fmt.Sprintf("Maximum limit of %v recipients reached", c.server.maxRecipients))
	}

	args, err := parse.Args(p.S)
	if err != nil {
		return smtp.NewStatus(501, smtp.EnhancedCode{5, 5, 4}, "Unable to parse RCPT ESMTP parameters")
	}

	opts := &smtp.RcptOptions{}

	for key, value := range args {
		switch key {
		case "NOTIFY":
			if !c.server.enableDSN {
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
			if !c.server.enableDSN {
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

	if err := c.session.Rcpt(c.ctx, recipient, opts); err != nil {
		return c.newStatusError(451, smtp.EnhancedCode{4, 0, 0}, "Recipient not accepted", err)
	}
	c.recipients++
	c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("I'll make sure <%v> gets this", recipient))
	return nil
}

func (c *Conn) handleAuth(arg string) error {
	if c.didAuth {
		return smtp.NewStatus(503, smtp.EnhancedCode{5, 5, 1}, "Already authenticated")
	}
	parts := strings.Fields(arg)
	if len(parts) == 0 {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 4}, "Missing parameter")
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

	sasl, err := c.session.Auth(c.ctx, mechanism)
	if err != nil {
		return c.newStatusError(454, smtp.EnhancedCode{4, 7, 0}, "Authentication failed", err)
	}

	response := ir
	for {
		challenge, done, err := sasl.Next(response)
		if err != nil {
			return c.newStatusError(454, smtp.EnhancedCode{4, 7, 0}, "Authentication failed", err)
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

	if c.state == StateEnforceAuthentication {
		c.state = StateGreeted
	}

	return nil
}

func (c *Conn) handleStartTLS() error {
	if _, isTLS := c.TLSConnectionState(); isTLS {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "Already running in TLS")
	}

	if c.server.tlsConfig == nil {
		return smtp.NewStatus(502, smtp.EnhancedCode{5, 5, 1}, "TLS not supported")
	}

	c.writeResponse(220, smtp.EnhancedCode{2, 0, 0}, "Ready to start TLS")

	// Upgrade to TLS
	tlsConn := tls.Server(c.conn, c.server.tlsConfig)

	if err := tlsConn.HandshakeContext(c.ctx); err != nil {
		c.logger().ErrorContext(c.ctx, "handleStartTLS", slog.Any("err", err))
		return smtp.NewStatus(550, smtp.EnhancedCode{5, 0, 0}, "Handshake error")
	}

	c.conn = tlsConn
	c.text.Replace(tlsConn)
	c.state = StateUpgrade // same as StateInit but calls logout/reset on ehlo/helo

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

	var r io.Reader

	rstart := func() io.Reader {
		if r != nil {
			return r
		}
		// We have recipients, go to accept data
		c.writeResponse(354, smtp.NoEnhancedCode, "Go ahead. End your data with <CR><LF>.<CR><LF>")

		r := textsmtp.NewDotReader(c.text.R, c.server.maxMessageBytes)
		return r
	}

	uuid, err := c.session.Data(c.ctx, rstart)
	if err != nil {
		return err
	}

	c.accepted(uuid)
	return c.reset()
}

func (c *Conn) handleBdat(arg string) error {
	size, last, err := bdatArg(arg)
	if err != nil {
		return err
	}

	data := &bdat{
		maxMessageBytes: c.server.maxMessageBytes,
		size:            size,
		last:            last,
		bytesReceived:   0,
		input:           c.text.R,
		nextCommand: func() (string, string, error) {
			c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "Continue")
			return c.nextCommand()
		},
	}

	uuid, err := c.session.Data(c.ctx, func() io.Reader { return data })

	if err == smtp.Reset {
		c.reset()
		c.writeStatus(smtp.Reset)
		return nil
	}

	if err != nil {
		return err
	}

	c.accepted(uuid)
	return c.reset()
}

func (c *Conn) accepted(uuid string) {
	if uuid != "" {
		if len(uuid) > 977 {
			uuid = uuid[:974] + "..."
		}
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "OK: queued as "+uuid)
	} else {
		c.writeResponse(250, smtp.EnhancedCode{2, 0, 0}, "OK: queued")
	}
}

func (c *Conn) Reject(ctx context.Context) {
	c.writeResponse(421, smtp.EnhancedCode{4, 4, 5}, "Too busy. Try again later.")
	c.Close(ctx)
}

func (c *Conn) greet() {
	protocol := "ESMTP"
	c.writeResponse(220, smtp.NoEnhancedCode, fmt.Sprintf("%v %s Service Ready", c.server.hostname, protocol))
}

func (c *Conn) writeStatus(status *smtp.SMTPStatus) {
	c.writeResponse(status.Code, status.EnhancedCode, status.Message)
}

func (c *Conn) writeResponse(code int, enhCode smtp.EnhancedCode, text ...string) {
	c.logger().DebugContext(c.ctx, "write", slog.Int("code", code), slog.Any("enhCode", enhCode), slog.Any("text", text))

	// TODO: error handling
	if c.server.writeTimeout != 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.server.writeTimeout))
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

func (c *Conn) newStatusError(code int, enhCode smtp.EnhancedCode, msg string, err error) *smtp.SMTPStatus {
	if smtpErr, ok := err.(*smtp.SMTPStatus); ok {
		return smtpErr
	} else {
		c.logger().ErrorContext(c.ctx, msg, slog.Any("err", err))
		return smtp.NewStatus(code, enhCode, msg)
	}
}

// Reads a line of input
func (c *Conn) readLine() (string, error) {
	if c.server.readTimeout != 0 {
		if err := c.conn.SetReadDeadline(time.Now().Add(c.server.readTimeout)); err != nil {
			return "", err
		}
	}
	line, err := c.text.ReadLine()
	if err == nil {
		c.logger().DebugContext(c.ctx, "read", slog.String("line", line))
	}
	return line, err
}

func (c *Conn) reset() error {
	// Reset state to Greeted
	if c.state == StateMail {
		c.state = StateGreeted
	}

	c.recipients = 0

	upgrade := c.state == StateUpgrade

	// Authentication is only revoked if starttls is used.
	if upgrade {
		c.didAuth = false
	}
	ctx, err := c.session.Reset(c.ctx, upgrade)
	c.ctx = ctx
	return err
}
