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
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/deadline"
	"github.com/uponusolutions/go-smtp/internal/parse"
	"github.com/uponusolutions/go-smtp/internal/smtpreader"
	"github.com/uponusolutions/go-smtp/internal/smtpwriter"
)

// errMessageTooLarge is returned by the DATA and BDAT handlers if the message
// exceeded the maximum message size. The payload remaining on the wire was not
// consumed, so the connection is out of sync and must be closed.
var errMessageTooLarge = errors.New("smtp: maximum message size exceeded")

// errTooManyTransactions is returned by the MAIL handler if the client exceeded
// the transactions per session, it can only continue in a new session.
var errTooManyTransactions = errors.New("smtp: maximum transactions exceeded")

type state int32

const (
	stateInit                    state = 0
	stateUpgrade                 state = 1
	stateEnforceAuthentication   state = 2
	stateEnforceSecureConnection state = 3
	stateGreeted                 state = 4
	stateMail                    state = 5
)

// Conn is a connection inside a smtp server.
type Conn struct {
	ctx context.Context

	conn     net.Conn
	receiver *smtpreader.Receiver
	sender   *smtpwriter.Sender

	state state

	server *Server

	session Session
	client  clientState // what the client negotiated
	tx      transaction // current mail transaction

	transactions int // MAIL commands in this session (RFC 9422 MAILMAX)
}

// clientState holds what the server knows about the client. It is discarded
// after STARTTLS (RFC 3207 section 4.2).
type clientState struct {
	esmtp      bool     // EHLO was used instead of HELO
	helo       string   // domain given in HELO / EHLO
	mechanisms []string // auth mechanisms offered in EHLO
	didAuth    bool     // client is authenticated
}

// transaction holds the state of a mail transaction. It is reset by MAIL,
// RSET and after the message has been received.
type transaction struct {
	binarymime bool // BODY=BINARYMIME was given in MAIL
	smtputf8   bool // SMTPUTF8 was given in MAIL
	recipients int  // number of accepted recipients
}

// run loops until an error occurs (quit for example)
func (c *Conn) run() (err error) {
	var (
		cmd string
		arg string
	)
	c.greet()

	for {
		c.renewReadDeadline()
		cmd, arg, err = c.nextCommand()
		if err != nil {
			return err
		}

		err = c.handle(cmd, arg)
		if err != nil {
			// if error is a smtp status it isn't necessary to close the connection
			if smtpErr, ok := err.(*smtp.Status); ok {
				// Service closing transmission channel, after quit
				if smtpErr.Code == 221 {
					return smtpErr
				}

				// ToDo: close connection on repeated errors (e.g. authentication tries)

				c.writeStatus(smtpErr)
				continue
			}

			return err
		}
	}
}

// nextCommand reads the next line and parses it.
// The command is always returned upper case.
func (c *Conn) nextCommand() (cmd string, arg string, err error) {
	line, err := c.readLine()
	if err != nil {
		return "", "", err
	}
	return parse.Cmd(line)
}

// Commands are dispatched to the appropriate handler functions.
func (c *Conn) handle(cmd string, arg string) error {
	if cmd == "" {
		return smtp.ErrBadSyntax
	}
	switch c.state {
	case stateInit, stateUpgrade:
		return c.handleStateInit(cmd, arg)
	case stateEnforceSecureConnection:
		return c.handleStateEnforceSecureConnection(cmd, arg)
	case stateEnforceAuthentication:
		return c.handleStateEnforceAuthentication(cmd, arg)
	case stateGreeted:
		return c.handleStateGreeted(cmd, arg)
	case stateMail:
		return c.handleStateMail(cmd, arg)
	default:
		return fmt.Errorf("unsupported state %d, how?", c.state)
	}
}

func (c *Conn) handleStateInit(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "NOOP":
		return smtp.Noop
	case "VRFY":
		return c.handleVrfy(arg)
	case "RSET": // Reset session
		return c.handleRSET()
	case "QUIT":
		return smtp.Quit
	default:
		return c.commandUnknown(cmd)
	}
}

func (c *Conn) handleStateEnforceAuthentication(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "NOOP":
		return smtp.Noop
	case "VRFY":
		return c.handleVrfy(arg)
	case "RSET": // Reset session
		return c.handleRSET()
	case "QUIT":
		return smtp.Quit
	case "AUTH":
		// there is always a mechanism, as it is an enforce authentication precondition
		return c.handleAuth(arg)
	case "STARTTLS":
		return c.handleStartTLS()
	// RFC 5321, always consume bdat data if capabilities were sent
	case "BDAT":
		return c.handleBdatDiscard(arg)
	default:
		return smtp.ErrAuthRequired
	}
}

func (c *Conn) handleStateGreeted(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "MAIL":
		return c.handleMail(arg)
	case "NOOP":
		return smtp.Noop
	case "VRFY":
		return c.handleVrfy(arg)
	case "RSET": // Reset session
		return c.handleRSET()
	case "QUIT":
		return smtp.Quit
	case "AUTH":
		if len(c.client.mechanisms) > 0 {
			return c.handleAuth(arg)
		}
		return smtp.ErrAuthUnsupported
	case "STARTTLS":
		return c.handleStartTLS()
	// RFC 5321, always consume bdat data if capabilities were sent
	case "BDAT":
		return c.handleBdatDiscard(arg)
	default:
		return c.commandUnknown(cmd)
	}
}

func (c *Conn) handleStateMail(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "RCPT":
		return c.handleRcpt(arg)
	case "NOOP":
		return smtp.Noop
	case "VRFY":
		return c.handleVrfy(arg)
	case "RSET": // Reset session
		return c.handleRSET()
	case "BDAT":
		if !c.server.enableCHUNKING {
			return c.handleBdatDiscard(arg)
		}
		return c.handleBdat(arg)
	case "DATA":
		return c.handleData(arg)
	case "QUIT":
		return smtp.Quit
	case "STARTTLS":
		return c.handleStartTLS()
	default:
		return c.commandUnknown(cmd)
	}
}

func (c *Conn) handleStateEnforceSecureConnection(cmd string, arg string) error {
	switch cmd {
	case "HELO", "EHLO":
		return c.handleGreet(cmd == "EHLO", arg)
	case "NOOP":
		return smtp.Noop
	case "VRFY":
		return c.handleVrfy(arg)
	case "STARTTLS":
		return c.handleStartTLS()
	case "QUIT":
		return smtp.Quit
	// RFC 5321, always consume bdat data if capabilities were sent
	case "BDAT":
		return c.handleBdatDiscard(arg)
	default:
		return smtp.ErrSecureConnectionRequired
	}
}

func (c *Conn) commandUnknown(cmd string) *smtp.Status {
	return smtp.NewStatusS(502, smtp.EnhancedCode{5, 5, 1}, fmt.Sprintf("%s command unknown, state %d", cmd, c.state))
}

// Server returns the server this connection comes from.
func (c *Conn) Server() *Server {
	return c.server
}

// Close closes the connection.
func (c *Conn) Close(err error) {
	c.logger().DebugContext(c.ctx, "connection is closing")

	// flush any pending data from writer before closing connection
	_ = c.sender.Flush()

	closeErr := c.conn.Close()
	if closeErr != nil {
		if err == nil {
			err = closeErr
		} else {
			err = errors.Join(err, closeErr)
		}
	}

	if err != nil {
		c.logger().ErrorContext(c.ctx, "close error", slog.Any("err", err))
	}

	if c.session != nil {
		c.session.Close(c.ctx, err)
		c.session = nil
	}
}

// TLSConnectionState returns the connection's TLS connection state.
// Zero values are returned if the connection doesn't use TLS.
func (c *Conn) TLSConnectionState() (tls.ConnectionState, bool) {
	tc, ok := c.conn.(*tls.Conn)
	if !ok {
		return tls.ConnectionState{}, ok
	}
	return tc.ConnectionState(), true
}

// IsTLS returns if the connection is encrypted by tls.
func (c *Conn) IsTLS() bool {
	_, ok := c.conn.(*tls.Conn)
	return ok
}

// Hostname returns the name of the connected client.
func (c *Conn) Hostname() string {
	return c.client.helo
}

// Esmtp returns true if ESMTP was used (EHLO instead of HELO).
func (c *Conn) Esmtp() bool {
	return c.client.esmtp
}

// Mechanisms returns the allowed auth mechanism for this connection.
func (c *Conn) Mechanisms() []string {
	return c.client.mechanisms
}

// Conn returns the connection.
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
func (c *Conn) handleGreet(esmtp bool, arg string) error {
	domain, err := parse.HelloArgument(arg)
	if err != nil {
		return smtp.ErrHeloArgRequired
	}

	// c.client.esmtp is populated before NewSession so
	// NewSession can access it via Conn.Esmtp.
	c.client.esmtp = esmtp

	// c.client.helo is populated before NewSession so
	// NewSession can access it via Conn.Hostname.
	c.client.helo = domain

	// RFC 5321: "An EHLO command MAY be issued by a client later in the session"
	// RFC 5321: "... the SMTP server MUST clear all buffers
	// and reset the state exactly as if a RSET command has been issued."
	if c.state != stateInit && c.state != stateEnforceSecureConnection && c.state != stateEnforceAuthentication {
		err := c.reset()
		if err != nil {
			return err
		}
	}

	if c.server.enforceSecureConnection && !c.IsTLS() {
		c.state = stateEnforceSecureConnection
	} else if c.server.enforceAuthentication && !c.client.didAuth {
		c.state = stateEnforceAuthentication
	} else {
		c.state = stateGreeted
	}

	if !esmtp {
		return smtp.NewStatusS(250, smtp.NoEnhancedCode, c.server.hostname+" greets "+c.client.helo)
	}

	c.client.mechanisms = c.session.AuthMechanisms(c.ctx)

	if len(c.client.mechanisms) == 0 && c.server.enforceAuthentication {
		// without any auth mechanism, no authentication can happen => deadlock
		return c.newStatusError(smtp.ErrNoAuthMechanism, err)
	}

	return c.handleGreetResponse()
}

func (c *Conn) handleGreetResponse() *smtp.Status {
	isTLS := c.IsTLS()

	// 17 lines is the current maximum
	lines := make([]string, 0, 17)

	lines = append(lines, c.server.hostname+" greets "+c.client.helo)
	lines = append(lines, "PIPELINING")
	lines = append(lines, "8BITMIME")
	lines = append(lines, "ENHANCEDSTATUSCODES")

	if c.server.enableCHUNKING {
		lines = append(lines, "CHUNKING")
	}
	if !isTLS && c.server.tlsConfig != nil {
		lines = append(lines, "STARTTLS")
	}
	if len(c.client.mechanisms) > 0 {
		lines = append(lines, "AUTH "+strings.Join(c.client.mechanisms, " "))
	}
	if c.server.enableSMTPUTF8 {
		lines = append(lines, "SMTPUTF8")
	}
	// We explicit allow require tls to be set if the connection is unencrypted
	// because if the server runs in a fully trusted zone
	// then the encryption is only relevant when it leaves this zone.
	if c.server.enableREQUIRETLS {
		lines = append(lines, "REQUIRETLS")
	}
	if c.server.enableBINARYMIME {
		lines = append(lines, "BINARYMIME")
	}
	if c.server.enableDSN {
		lines = append(lines, "DSN")
	}
	if c.server.enableXOORG {
		lines = append(lines, "XOORG")
	}
	if c.server.maxMessageBytes > 0 {
		lines = append(lines, "SIZE "+strconv.FormatInt(c.server.maxMessageBytes, 10))
	} else {
		lines = append(lines, "SIZE")
	}
	if limits := c.limits(); limits != "" {
		lines = append(lines, "LIMITS"+limits)
	}
	if c.server.enableRRVS {
		lines = append(lines, "RRVS")
	}
	if c.server.enableDELIVERBY {
		// maximum of nine digits
		if c.server.minimumDeliverInSeconds > 0 && c.server.minimumDeliverInSeconds < 999999999 {
			lines = append(lines, "DELIVERBY "+strconv.Itoa(c.server.minimumDeliverInSeconds))
		} else {
			lines = append(lines, "DELIVERBY")
		}
	}
	if c.server.enableMTPRIORITY {
		if c.server.mtPriorityProfile != smtp.PriorityUnspecified {
			lines = append(lines, "MT-PRIORITY "+string(c.server.mtPriorityProfile))
		} else {
			lines = append(lines, "MT-PRIORITY")
		}
	}

	return smtp.NewStatusM(250, smtp.NoEnhancedCode, lines)
}

// limits returns the RFC 9422 LIMITS parameters with a leading space, or ""
// if there are none. Values are only valid from 1 to 999999, a larger limit
// isn't advertised (but still enforced).
func (c *Conn) limits() string {
	var sb strings.Builder
	for _, l := range []struct {
		name  string
		value int
	}{
		{"MAILMAX", c.server.maxTransactions},
		{"RCPTMAX", c.server.maxRecipients},
	} {
		if l.value > 0 && l.value <= 999999 {
			sb.WriteString(" " + l.name + "=" + strconv.Itoa(l.value))
		}
	}
	return sb.String()
}

// handleError handles error and closes the connection afterwards.
func (c *Conn) handleError(err error) {
	if err == io.EOF || errors.Is(err, net.ErrClosed) {
		c.Close(fmt.Errorf("connection closed unexpectedly: %w", err))
		return
	}

	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
		c.writeStatus(smtp.ErrIdleTimeout)
		c.Close(fmt.Errorf("idle timeout: %w", err))
		return
	}

	if smtpErr, ok := err.(*smtp.Status); ok {
		c.writeStatus(smtpErr)

		if smtpErr.Code != 221 {
			c.Close(fmt.Errorf("smtp error: %w", err))
		} else {
			c.Close(nil)
		}

		return
	}

	if errors.Is(err, errTooManyTransactions) {
		c.writeStatus(smtp.ErrTooManyTransactions)
		c.Close(err)
		return
	}

	if errors.Is(err, errMessageTooLarge) {
		c.writeStatus(smtp.ErrDataTooLarge)
		c.Close(err)
		return
	}

	if err == smtp.ErrTooLongLine {
		c.writeStatus(smtp.ErrCommandLineTooLong)
		c.Close(errors.New("line too long"))
		return
	}

	c.writeStatus(smtp.ErrConnection)
	c.Close(fmt.Errorf("unknown error: %w", err))
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
// nolint: revive
func (c *Conn) handleMail(arg string) error {
	// RFC 9422: every MAIL command counts, regardless of whether it succeeds.
	c.transactions++
	if c.server.maxTransactions > 0 && c.transactions > c.server.maxTransactions {
		return errTooManyTransactions
	}

	arg, ok := parse.CutPrefixFold(arg, "FROM:")
	if !ok {
		return smtp.ErrMailSyntax
	}

	p := parse.Parser{S: strings.TrimSpace(arg)}
	from, err := p.ReversePath()
	if err != nil {
		return c.newStatusError(smtp.ErrMailSyntax, err)
	}
	args, err := parse.Args(p.S)
	if err != nil {
		return c.newStatusError(smtp.ErrMailParams, err)
	}

	opts := &smtp.MailOptions{}

	c.tx = transaction{}
	// This is where the Conn may put BODY=8BITMIME, but we already
	// read the DATA as bytes, so it does not effect our processing.
	for _, arg := range args {
		switch arg.Key {
		case "SIZE":
			size, err := strconv.ParseUint(arg.Value, 10, 32)
			if err != nil {
				return smtp.ErrMalformedSize
			}

			if c.server.maxMessageBytes > 0 && int64(size) > c.server.maxMessageBytes {
				return smtp.ErrDataTooLarge
			}

			opts.Size = int64(size)
		case "XOORG":
			value, err := decodeXtext(arg.Value)
			if err != nil || value == "" {
				return smtp.ErrMalformedXOORG
			}
			if !c.server.enableXOORG {
				return smtp.ErrXOORGNotImplemented
			}
			opts.XOORG = &value
		case "SMTPUTF8":
			if !c.server.enableSMTPUTF8 {
				return smtp.ErrSMTPUTF8NotImplemented
			}
			opts.UTF8 = true
		case "REQUIRETLS":
			if !c.server.enableREQUIRETLS {
				return smtp.ErrREQUIRETLSNotImplemented
			}
			opts.RequireTLS = true
		case "BODY":
			value := strings.ToUpper(arg.Value)
			switch smtp.BodyType(value) {
			case smtp.BodyBinaryMIME:
				if !c.server.enableBINARYMIME {
					return smtp.ErrBINARYMIMENotImplemented
				}
				c.tx.binarymime = true
			case smtp.Body7Bit, smtp.Body8BitMIME:
				// This space is intentionally left blank
			default:
				return smtp.ErrUnknownBody
			}
			opts.Body = smtp.BodyType(value)
		case "RET":
			if !c.server.enableDSN {
				return smtp.ErrRETNotImplemented
			}
			value := strings.ToUpper(arg.Value)
			switch smtp.DSNReturn(value) {
			case smtp.DSNReturnFull, smtp.DSNReturnHeaders:
				// This space is intentionally left blank
			default:
				return smtp.ErrUnknownRET
			}
			opts.Return = smtp.DSNReturn(value)
		case "ENVID":
			if !c.server.enableDSN {
				return smtp.ErrENVIDNotImplemented
			}
			value, err := decodeXtext(arg.Value)
			if err != nil || value == "" || !parse.IsPrintableASCII(value) {
				return smtp.ErrMalformedENVID
			}
			opts.EnvelopeID = value
		case "AUTH":
			value, err := decodeXtext(arg.Value)
			if err != nil || value == "" {
				return smtp.ErrMalformedAUTH
			}
			if value == "<>" {
				value = ""
			} else {
				p := parse.Parser{S: value}
				value, err = p.Mailbox()
				if err != nil || p.S != "" {
					return smtp.ErrMalformedAUTHMailbox
				}
			}
			opts.Auth = &value
		case "BY":
			if err := handleMailBY(c.server, opts, arg.Value); err != nil {
				return err
			}
		case "MT-PRIORITY":
			if err := handleMailMTPRIORITY(c.server, opts, arg.Value); err != nil {
				return err
			}
		default:
			return smtp.ErrUnknownMailArg
		}
	}

	if !opts.UTF8 && !parse.IsASCII(from) {
		return smtp.ErrNonASCIISender
	}

	// The null reverse-path is not an address.
	if from != "" {
		if err := c.validateAddress(from, smtp.ErrSenderNotAccepted); err != nil {
			return err
		}
	}

	c.tx.smtputf8 = opts.UTF8

	if err := c.session.Mail(c.ctx, from, opts); err != nil {
		if smtpErr, ok := err.(*smtp.Status); ok {
			// a positive response also counts as a success
			if smtpErr.Positive() {
				c.state = stateMail
			}
			return smtpErr
		}
		return c.newStatusError(smtp.ErrMailNotAccepted, err)
	}

	c.state = stateMail
	return smtp.NewStatusS(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("Roger, accepting mail from <%v>", from))
}

func handleMailBY(server *Server, opts *smtp.MailOptions, value string) error {
	if !server.enableDELIVERBY {
		return smtp.ErrDELIVERBYNotImplemented
	}
	deliverBy := parseDeliverByArgument(value)
	if deliverBy == nil {
		return smtp.ErrMalformedBY
	}
	if server.minimumDeliverInSeconds != 0 &&
		deliverBy.Mode == smtp.DeliverByReturn &&
		deliverBy.Seconds < server.minimumDeliverInSeconds {
		return smtp.ErrBYBelowMinimum
	}
	opts.DeliverBy = deliverBy
	return nil
}

// validateAddress calls the address validator of the server, if set.
// If it returns an error that is not a negative *smtp.Status, the reply is
// rejected, a 553 (mailbox name not allowed, RFC 5321 section 4.2.2).
func (c *Conn) validateAddress(addr string, rejected *smtp.Status) error {
	if c.server.addressValidator == nil {
		return nil
	}

	err := c.server.addressValidator(addr)
	if err == nil {
		return nil
	}

	if status, ok := err.(*smtp.Status); ok && status.Positive() {
		// A positive reply would end the command without a result.
		err = errors.New("address validator returned a positive reply")
	}

	return c.newStatusError(rejected, err)
}

func handleMailMTPRIORITY(server *Server, opts *smtp.MailOptions, value string) error {
	if !server.enableMTPRIORITY {
		return smtp.ErrMTPRIORITYNotImplemented
	}
	mtPriority, err := strconv.Atoi(value)
	if err != nil {
		return smtp.ErrMalformedMTPRIORITY
	}
	if mtPriority < -9 || mtPriority > 9 {
		return smtp.ErrMTPRIORITYOutOfRange
	}
	opts.MTPriority = &mtPriority
	return nil
}

// MAIL state -> waiting for RCPTs followed by DATA
func (c *Conn) handleRcpt(arg string) error {
	arg, ok := parse.CutPrefixFold(arg, "TO:")
	if !ok {
		return smtp.ErrRcptSyntax
	}

	p := parse.Parser{S: strings.TrimSpace(arg)}
	recipient, err := p.Path()
	if err != nil {
		return smtp.ErrRcptSyntax
	}
	if !c.tx.smtputf8 && !parse.IsASCII(recipient) {
		return smtp.ErrNonASCIIRecipient
	}

	if err := c.validateAddress(recipient, smtp.ErrRcptAddressNotAccepted); err != nil {
		return err
	}

	if c.server.maxRecipients > 0 && c.tx.recipients >= c.server.maxRecipients {
		return smtp.NewStatusS(452, smtp.EnhancedCode{4, 5, 3},
			fmt.Sprintf("Maximum limit of %v recipients reached", c.server.maxRecipients),
		)
	}

	args, err := parse.Args(p.S)
	if err != nil {
		return c.newStatusError(smtp.ErrRcptParams, err)
	}

	opts := &smtp.RcptOptions{}

	for _, arg := range args {
		switch arg.Key {
		case "NOTIFY":
			if err := handleRcptNotify(c.server, opts, arg.Value); err != nil {
				return err
			}
		case "ORCPT":
			if err := handleRcptORCPT(c.server, opts, arg.Value); err != nil {
				return err
			}
		case "RRVS":
			if err := handleRcptRRVS(c.server, opts, arg.Value); err != nil {
				return err
			}
		default:
			return smtp.ErrUnknownRcptArg
		}
	}

	if err := c.session.Rcpt(c.ctx, recipient, opts); err != nil {
		if smtpErr, ok := err.(*smtp.Status); ok {
			// a positive response also counts as a success
			if smtpErr.Positive() {
				c.tx.recipients++
			}
			return smtpErr
		}
		return c.newStatusError(smtp.ErrRcptNotAccepted, err)
	}

	c.tx.recipients++
	return smtp.NewStatusS(250, smtp.EnhancedCode{2, 0, 0}, fmt.Sprintf("I'll make sure <%v> gets this", recipient))
}

func handleRcptNotify(server *Server, opts *smtp.RcptOptions, value string) error {
	if !server.enableDSN {
		return smtp.ErrNOTIFYNotImplemented
	}
	notify := []smtp.DSNNotify{}
	for val := range strings.SplitSeq(value, ",") {
		notify = append(notify, smtp.DSNNotify(strings.ToUpper(val)))
	}
	if err := parse.CheckNotifySet(notify); err != nil {
		return smtp.ErrMalformedNOTIFY
	}
	opts.Notify = notify
	return nil
}

func handleRcptORCPT(server *Server, opts *smtp.RcptOptions, value string) error {
	if !server.enableDSN {
		return smtp.ErrORCPTNotImplemented
	}
	aType, aAddr, err := decodeTypedAddress(value)
	if err != nil || aAddr == "" {
		return smtp.ErrMalformedORCPT
	}
	opts.OriginalRecipientType = aType
	opts.OriginalRecipient = aAddr
	return nil
}

func handleRcptRRVS(server *Server, opts *smtp.RcptOptions, value string) error {
	if !server.enableRRVS {
		return smtp.ErrRRVSNotImplemented
	}
	value, _, _ = strings.Cut(value, ";") // discard the no-support action
	rrvsTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return smtp.ErrMalformedRRVS
	}
	opts.RequireRecipientValidSince = rrvsTime
	return nil
}

func (c *Conn) handleVrfy(arg string) error {
	p := parse.Parser{S: strings.TrimSpace(arg)}
	vrfy, err := p.Path()
	if err != nil {
		return smtp.ErrVrfySyntax
	}
	args, err := parse.Args(p.S)
	if err != nil {
		return c.newStatusError(smtp.ErrVrfyParams, err)
	}

	opts := &smtp.VrfyOptions{}

	for _, arg := range args {
		if arg.Key == "SMTPUTF8" {
			if !c.server.enableSMTPUTF8 {
				return smtp.ErrSMTPUTF8NotImplemented
			}
			opts.UTF8 = true
		}
	}
	if !opts.UTF8 && !parse.IsASCII(vrfy) {
		return smtp.ErrUTF8ReplyRequired
	}

	if err := c.validateAddress(vrfy, smtp.ErrVrfyAddressNotAccepted); err != nil {
		return err
	}

	res := c.session.Verify(c.ctx, vrfy, opts)

	if res == nil {
		return smtp.VRFY
	}

	return res
}

func (c *Conn) handleAuth(arg string) error {
	if c.client.didAuth {
		return smtp.ErrAlreadyAuthenticated
	}
	parts := strings.Fields(arg)
	if len(parts) == 0 {
		return smtp.ErrAuthMissingParam
	}

	mechanism := strings.ToUpper(parts[0])

	// Is mechanism allowed?
	if !slices.Contains(c.client.mechanisms, mechanism) {
		return smtp.ErrAuthUnknownMechanism
	}

	// Parse client initial response if there is one
	var ir []byte
	if len(parts) > 1 {
		var err error
		ir, err = decodeSASLResponse(parts[1])
		if err != nil {
			return smtp.ErrInvalidBase64
		}
	}

	sasl, err := c.session.Auth(c.ctx, mechanism)
	if err != nil {
		return c.newStatusError(smtp.ErrAuthTemporaryFailure, err)
	}

	if sasl == nil {
		return c.newStatusError(smtp.ErrNoAuthHandler, err)
	}

	response := ir
	for {
		challenge, done, err := sasl.Next(response)
		if err != nil {
			return c.newStatusError(smtp.ErrAuthTemporaryFailure, err)
		}

		if done {
			break
		}

		encoded := ""
		if len(challenge) > 0 {
			encoded = base64.StdEncoding.EncodeToString(challenge)
		}
		c.writeStatus(smtp.NewStatusS(334, smtp.NoEnhancedCode, encoded))

		c.renewReadDeadline()
		encoded, err = c.readLine()
		if err != nil {
			return err
		}

		if encoded == "*" {
			// https://tools.ietf.org/html/rfc4954#page-4
			return smtp.ErrAuthCancelled
		}

		response, err = decodeSASLResponse(encoded)
		if err != nil {
			return smtp.ErrInvalidBase64
		}
	}

	c.client.didAuth = true
	if c.state == stateEnforceAuthentication {
		c.state = stateGreeted
	}

	return smtp.AuthSucceeded
}

func (c *Conn) handleStartTLS() error {
	if _, isTLS := c.TLSConnectionState(); isTLS {
		return smtp.ErrAlreadyTLS
	}

	if c.server.tlsConfig == nil {
		return smtp.ErrTLSNotSupported
	}

	// allow the session to change tlsConfig
	tlsConfig, err := c.session.STARTTLS(c.ctx, c.server.tlsConfig)
	if err != nil {
		return c.newStatusError(smtp.ErrTLSConfig, err)
	}
	if tlsConfig == nil {
		return smtp.ErrTLSConfigNil
	}

	c.writeStatus(smtp.StartTLS)

	// Upgrade to TLS
	tlsConn := tls.Server(c.conn, tlsConfig)

	c.renewReadDeadline()

	if err := tlsConn.HandshakeContext(c.ctx); err != nil {
		// err is not a smtp error and will terminate the connection
		return err
	}

	c.conn = tlsConn
	c.receiver.Reset(tlsConn)
	c.sender.Reset(tlsConn)
	// RFC 3207 section 4.2: the server MUST discard any knowledge obtained
	// from the client, such as the argument to the EHLO command, which was
	// not obtained from the TLS negotiation itself. This also revokes the
	// authentication.
	c.client = clientState{}
	c.state = stateUpgrade // same as StateInit but calls logout/reset on ehlo/helo

	return nil
}

// DATA
func (c *Conn) handleData(arg string) error {
	// at least a single recipient needs to be set
	if c.tx.recipients == 0 {
		return smtp.ErrNoRecipients
	}

	if arg != "" {
		return smtp.ErrDataArgs
	}
	if c.tx.binarymime {
		return smtp.ErrDataBINARYMIME
	}

	var r io.Reader

	rstart := func() io.Reader {
		if r != nil {
			return r
		}
		// We have recipients, go to accept data
		c.writeStatus(smtp.StartMailInput)
		c.setDataDeadline()

		// r gets exposed to be able to discard the rest of the message
		r = smtpreader.NewDot(c.receiver.Reader, c.server.maxMessageBytes)
		return r
	}

	uuid, err := c.session.Data(c.ctx, rstart)

	// an error which isn't a smtp status error will always terminate the connection
	if _, ok := err.(*smtp.Status); err != nil && !ok {
		return err
	}

	// Make sure all the data has been consumed, regardless of what the backend returned.
	// The reader keeps returning ErrDataTooLarge once the limit is hit, so this also
	// detects a breach the backend swallowed. In that case the rest of the message is
	// still on the wire and the connection can't be resynchronised.
	if r != nil {
		if _, discardErr := io.Copy(io.Discard, r); errors.Is(discardErr, smtp.ErrDataTooLarge) {
			return errMessageTooLarge
		}
	}

	// Without a 354 the client sent no message, a rejected DATA command leaves the transaction open.
	if r == nil && err != nil {
		return err
	}

	// RFC 5321 4.1.1.4: the end of mail data ends the transaction, whether the message
	// was accepted or rejected. Keeping it would leave sender and recipients behind,
	// a following MAIL is refused and a pipelined RCPT/DATA reuses the old envelope.
	if resetErr := c.reset(); resetErr != nil {
		return resetErr
	}
	if err != nil {
		return err
	}
	return c.accepted(uuid)
}

// discardBdat discards a BDAT chunk that isn't processed. Like a processed
// chunk, it is limited to the maximum message size. A larger chunk isn't read,
// it stays on the wire and the connection must be closed.
func (c *Conn) discardBdat(size int64) error {
	if c.server.maxMessageBytes > 0 && size > c.server.maxMessageBytes {
		return errMessageTooLarge
	}
	_, err := c.receiver.Discard(int(size))
	return err
}

func (c *Conn) handleBdatDiscard(arg string) error {
	size, _, err := smtpreader.BdatArg(arg)
	if err != nil {
		return err
	}
	c.setDataDeadline()
	if err = c.discardBdat(size); err != nil {
		return err
	}
	if !c.server.enableCHUNKING {
		return smtp.ErrCHUNKINGNotImplemented
	}
	return smtp.ErrBadSequence
}

func (c *Conn) handleBdat(arg string) error {
	size, last, err := smtpreader.BdatArg(arg)
	if err != nil {
		return err
	}

	// The whole transfer, all chunks until LAST, has to finish within the data timeout.
	// The BDAT lines between chunks are read without renewing the deadline.
	c.setDataDeadline()

	// at least a single recipient needs to be set
	if c.tx.recipients == 0 {
		if err = c.discardBdat(size); err != nil {
			return err
		}
		return smtp.ErrNoRecipients
	}

	closed := false

	data := smtpreader.NewBdat(size, last, c.server.maxMessageBytes, c.receiver, func() (string, string, error) {
		// if bdat is closed (error occurred)
		if closed {
			return "", "", io.EOF
		}
		c.writeStatus(smtp.BDATContinue)
		return c.nextCommand()
	})

	queueid, err := c.session.Data(c.ctx, func() io.Reader {
		return data
	})

	smtpErr, isStatus := err.(*smtp.Status)

	// an error which isn't a SMTPStatus error will always terminate the connection
	if err != nil && !isStatus {
		return err
	}

	if err != nil {
		// read anything left to continue after this failure
		// https://www.rfc-editor.org/rfc/rfc3030.html
		// If a 5XX or 4XX code is received by the sender-SMTP in response to a BDAT
		// chunk, the transaction should be considered failed and the sender-
		// SMTP MUST NOT send any additional BDAT segments.  If the receiver-
		// SMTP has declared support for command pipelining [PIPE], the receiver
		// SMTP MUST be prepared to accept and discard additional BDAT chunks
		// already in the pipeline after the failed BDAT.
		closed = true
	}

	// Make sure all the data has been consumed, regardless of what the backend returned.
	// The reader keeps returning ErrDataTooLarge once the limit is hit, so this also
	// detects a breach the backend swallowed. In that case the chunk is still on the
	// wire and the connection can't be resynchronised.
	if _, discardErr := io.Copy(io.Discard, data); errors.Is(discardErr, smtp.ErrDataTooLarge) {
		return errMessageTooLarge
	}

	if err != nil {
		// write down error after data is discarded to prevent deadlock because of pipelining
		c.writeStatus(smtpErr)
		return c.reset()
	}

	if err = c.reset(); err != nil {
		return err
	}
	return c.accepted(queueid)
}

func (*Conn) accepted(queueid string) *smtp.Status {
	if queueid != "" {
		// limit length if queueid is too long (< 1000)
		if len(queueid) > 977 {
			queueid = queueid[:974] + "..."
		}
		return smtp.NewStatusS(250, smtp.EnhancedCode{2, 0, 0}, "OK: queued as "+queueid)
	}
	return smtp.Queued
}

func (c *Conn) greet() {
	protocol := "ESMTP"
	c.writeStatus(
		smtp.NewStatusS(220, smtp.NoEnhancedCode, fmt.Sprintf("%v %s Service Ready", c.server.hostname, protocol)),
	)
}

func (c *Conn) writeStatus(status *smtp.Status) {
	c.logger().DebugContext(
		c.ctx, "write status", slog.Any("status", status),
	)

	deadline.SetWrite(c.conn, c.server.writeTimeout)

	_, err := status.WriteTo(c.sender)

	// PIPELINE support
	// If there is something buffered in c.text.R then we can assume another command is following.
	// This means the client is doing pipelining and we don't need to respond just now.
	if err == nil && c.receiver.Buffered() == 0 {
		err = c.sender.Flush()
	}

	// A failed write (e.g. write timeout) loses the reply, the session can't go on.
	// Closing the connection makes the next read fail and ends the session.
	if err != nil {
		_ = c.conn.Close()
	}
}

// newStatusError returns err if it is a status, otherwise it logs err and returns status.
func (c *Conn) newStatusError(status *smtp.Status, err error) *smtp.Status {
	if smtpErr, ok := err.(*smtp.Status); ok {
		return smtpErr
	}
	c.logger().ErrorContext(c.ctx, status.Text(), slog.Any("err", err))
	return status
}

// renewReadDeadline sets the read deadline to now plus the read timeout,
// a zero read timeout means no deadline.
func (c *Conn) renewReadDeadline() {
	deadline.SetRead(c.conn, c.server.readTimeout)
}

// setDataDeadline sets the read deadline for a whole DATA or BDAT transfer
// to now plus the data timeout, a zero data timeout means no deadline.
func (c *Conn) setDataDeadline() {
	deadline.SetRead(c.conn, c.server.dataTimeout)
}

// readLine reads a line of input.
// It doesn't touch the read deadline, the caller has to set it.
func (c *Conn) readLine() (string, error) {
	line, err := c.receiver.ReadFullLine()
	if err == nil {
		c.logger().DebugContext(c.ctx, "read", slog.String("line", line))
	}
	return line, err
}

func (c *Conn) reset() error {
	// Reset state to Greeted
	if c.state == stateMail {
		c.state = stateGreeted
	}

	c.tx = transaction{}

	ctx, err := c.session.Reset(c.ctx, c.state == stateUpgrade)
	c.ctx = ctx
	return err
}
