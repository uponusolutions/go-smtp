package smtp

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// ErrTooLongLine occurs if the smtp line is too long.
	ErrTooLongLine = errors.New("smtp: too long a line in input stream")

	// Crnl \r\n
	Crnl = []byte{'\r', '\n'}
	// Dotcrnl .\r\n
	Dotcrnl = []byte{'.', '\r', '\n'}
	// Crlfdot \r\n.
	Crlfdot = []byte{'\r', '\n', '.'}
)

// EnhancedCode is the SMTP enhanced code
type EnhancedCode [3]int

// Status specifies the error code, enhanced error code (if any) and
// message returned by the server.
type Status struct {
	Code         int
	EnhancedCode EnhancedCode
	Lines        []string
}

// NoEnhancedCode is used to indicate that enhanced error code should not be
// included in response.
//
// Note that RFC 2034 requires an enhanced code to be included in all 2xx, 4xx
// and 5xx responses.
var NoEnhancedCode = EnhancedCode{0, 0, 0}

// ToPart returns the part of the string after the code
// which is defined by the enhanced code with the trailing whitespace.
// E.g. "5.1.1 " or "5.7.26 " (RFC 3463: subject and detail have up to 3 digits).
func (enhCode EnhancedCode) ToPart() []byte {
	if enhCode == NoEnhancedCode {
		return nil
	}
	// 10 is the longest valid code with the trailing space, e.g. "5.123.123 ".
	b := make([]byte, 0, 10)
	b = strconv.AppendInt(b, int64(enhCode[0]), 10)
	b = append(b, '.')
	b = strconv.AppendInt(b, int64(enhCode[1]), 10)
	b = append(b, '.')
	b = strconv.AppendInt(b, int64(enhCode[2]), 10)
	return append(b, ' ')
}

// NewStatusM creates a new status with multiple message lines.
func NewStatusM(code int, enhCode EnhancedCode, msg []string) *Status {
	return &Status{
		Code:         code,
		EnhancedCode: enhCode,
		Lines:        msg,
	}
}

// NewStatusS creates a new status with a single message line.
func NewStatusS(code int, enhCode EnhancedCode, msg string) *Status {
	return &Status{
		Code:         code,
		EnhancedCode: enhCode,
		Lines:        []string{msg},
	}
}

// Error returns a error string.
func (s *Status) Error() string {
	return s.String()
}

// Positive returns true if the status code is 2xx.
func (s *Status) Positive() bool {
	return s.Code/100 == 2
}

// Temporary returns true if the status code is 4xx.
func (s *Status) Temporary() bool {
	return s.Code/100 == 4
}

// Permanent returns true if the status code is 5xx.
func (s *Status) Permanent() bool {
	return s.Code/100 == 5
}

// Text returns all lines joined by \n in a single string.
func (s *Status) Text() string {
	return strings.Join(s.Lines, "\n")
}

// String returns the status as a formatted message.
// Important: Should not be used for sending a response.
func (s *Status) String() string {
	base := fmt.Sprintf("%03d", s.Code)
	if s.EnhancedCode != NoEnhancedCode {
		base += fmt.Sprintf(" %d.%d.%d", s.EnhancedCode[0], s.EnhancedCode[1], s.EnhancedCode[2])
	}
	if len(s.Lines) > 0 {
		return base + " " + s.Text()
	}
	return base
}

// StatusWriter is an interface which needs to be implemented by a writer to be used by Status.WriteTo.
type StatusWriter interface {
	io.ByteWriter
	io.StringWriter
	io.Writer
}

// writeLine writes a single reply line. last selects the terminating form.
func writeLine(w StatusWriter, code string, enhCode []byte, message string, last bool) (i int, err error) {
	var p int

	i, err = w.WriteString(code)
	if err != nil {
		return i, err
	}

	if last {
		// RFC 5321 permits omitting the space when there is no text, but
		// RFC 4954 requires it for the 334 challenge and net/textproto
		// rejects any reply shorter than four bytes. Always emit it.
		err = w.WriteByte(' ')
	} else {
		err = w.WriteByte('-')
	}
	if err != nil {
		return i, err
	}
	i++ // single byte added

	p, err = w.Write(enhCode)
	i += p
	if err != nil {
		return i, err
	}

	p, err = w.WriteString(message)
	i += p
	if err != nil {
		return i, err
	}

	p, err = w.Write(Crnl)
	i += p

	return i, err
}

// WriteTo writes the smtp status reply.
func (s *Status) WriteTo(w StatusWriter) (int64, error) {
	codeString := strconv.Itoa(s.Code)
	enhCodeString := s.EnhancedCode.ToPart()

	// no message set
	if len(s.Lines) == 0 {
		p, err := writeLine(w, codeString, enhCodeString, "", true)
		return int64(p), err
	}

	var i int
	for q, m := range s.Lines {
		p, err := writeLine(w, codeString, enhCodeString, m, q+1 == len(s.Lines))
		i += p
		if err != nil {
			return int64(i), err
		}
	}

	return int64(i), nil
}

// LogValue implements slog.LogValuer so that formatting a Status is deferred
// to the handler and skipped entirely when the record is dropped.
func (s *Status) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("code", s.Code),
		slog.String("enhCode", fmt.Sprintf("%d.%d.%d", s.EnhancedCode[0], s.EnhancedCode[1], s.EnhancedCode[2])),
		slog.String("text", s.Text()),
	)
}

// Static statuses, shared by all connections. They must not be modified.
var (
	// Success and intermediate replies.

	// Noop default return.
	Noop = &Status{
		Code:         250,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"I have successfully done nothing"},
	}
	// Quit is returned by Reader passed to Data function if client does not
	// send another BDAT command and instead issues QUIT command.
	Quit = &Status{
		Code:         221,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"Bye"},
	}
	// Reset is returned by Reader passed to Data function if client does not
	// send another BDAT command and instead issues RSET command.
	Reset = &Status{
		Code:         250,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"Session reset"},
	}
	// VRFY default return.
	VRFY = &Status{
		Code:         252,
		EnhancedCode: EnhancedCode{2, 5, 0},
		Lines:        []string{"Cannot VRFY user, but will accept message"},
	}
	// AuthSucceeded is returned if the authentication succeeded.
	AuthSucceeded = &Status{
		Code:         235,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"Authentication succeeded"},
	}
	// StartTLS is returned before the TLS handshake of STARTTLS.
	StartTLS = &Status{
		Code:         220,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"Ready to start TLS"},
	}
	// BDATContinue is returned after a BDAT chunk that isn't the last one.
	BDATContinue = &Status{
		Code:         250,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"Continue"},
	}
	// Queued is returned if a message was accepted without a queue id.
	Queued = &Status{
		Code:         250,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Lines:        []string{"OK: queued"},
	}
	// StartMailInput is returned if the server is ready to receive the DATA content.
	StartMailInput = &Status{
		Code:         354,
		EnhancedCode: NoEnhancedCode,
		Lines:        []string{"Go ahead. End your data with <CR><LF>.<CR><LF>"},
	}

	// Connection errors.

	// ErrConnection is returned if a connection error occurs.
	ErrConnection = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 4, 0},
		Lines:        []string{"Connection error, sorry"},
	}
	// ErrInternal is returned if the server recovered from an internal error (panic).
	ErrInternal = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"Internal server error"},
	}
	// ErrTooManyConnections is returned if the server's concurrent connection limit is exceeded.
	ErrTooManyConnections = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 3, 2},
		Lines:        []string{"Too many concurrent connections, try again later"},
	}
	// ErrTooManyTransactions is returned if the client exceeded the
	// transactions per session (RFC 9422 MAILMAX). The connection is closed.
	ErrTooManyTransactions = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 7, 0},
		Lines:        []string{"Too many transactions in this session"},
	}
	// ErrIdleTimeout is returned if the client exceeded the read timeout.
	ErrIdleTimeout = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 4, 2},
		Lines:        []string{"Idle timeout, bye bye"},
	}
	// ErrCommandLineTooLong is returned if a command line exceeds the maximum line length.
	ErrCommandLineTooLong = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 4, 0},
		Lines:        []string{"Too long line"},
	}

	// Command and sequence errors.

	// ErrBadSyntax is returned if a command line can't be parsed.
	ErrBadSyntax = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 2},
		Lines:        []string{"Error: bad syntax"},
	}
	// ErrBadSequence is returned if a command is invalid at this point.
	ErrBadSequence = &Status{
		Code:         503,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"Bad sequence of commands"},
	}
	// ErrHeloArgRequired is returned if HELO or EHLO is sent without an argument.
	ErrHeloArgRequired = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 2},
		Lines:        []string{"Domain/address argument required for HELO"},
	}
	// ErrSecureConnectionRequired is returned if a secure connection is enforced and the client didn't use TLS.
	ErrSecureConnectionRequired = &Status{
		Code:         530,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Lines:        []string{"Must issue a STARTTLS command first"},
	}
	// ErrNoAuthMechanism is returned if authentication is enforced but the session offers no mechanism.
	ErrNoAuthMechanism = &Status{
		Code:         451,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"No auth mechanism available but authentication enforced"},
	}

	// MAIL errors.

	// ErrMailSyntax is returned if the MAIL command can't be parsed.
	ErrMailSyntax = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 2},
		Lines:        []string{"Was expecting MAIL arg syntax of FROM:<address>"},
	}
	// ErrMailParams is returned if the MAIL ESMTP parameters can't be parsed.
	ErrMailParams = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unable to parse MAIL ESMTP parameters"},
	}
	// ErrUnknownMailArg is returned if MAIL contains an unknown parameter.
	ErrUnknownMailArg = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unknown MAIL FROM argument"},
	}
	// ErrNonASCIISender is returned if the MAIL address contains non-ASCII
	// characters but SMTPUTF8 was not requested (RFC 6531 section 3.5).
	ErrNonASCIISender = &Status{
		Code:         550,
		EnhancedCode: EnhancedCode{5, 6, 7},
		Lines:        []string{"Non-ASCII addresses not permitted for that sender"},
	}
	// ErrMailNotAccepted is returned if the backend rejected MAIL with a non-SMTP error.
	ErrMailNotAccepted = &Status{
		Code:         451,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"Mail not accepted"},
	}
	// ErrMalformedSize is returned if the SIZE parameter isn't a valid integer.
	ErrMalformedSize = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unable to parse SIZE as an integer"},
	}
	// ErrMalformedXOORG is returned if the XOORG parameter is malformed.
	ErrMalformedXOORG = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed XOORG parameter value"},
	}
	// ErrXOORGNotImplemented is returned if XOORG is used but not enabled.
	ErrXOORGNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"EnableXOORG is not implemented"},
	}
	// ErrSMTPUTF8NotImplemented is returned if SMTPUTF8 is used but not enabled.
	ErrSMTPUTF8NotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"SMTPUTF8 is not implemented"},
	}
	// ErrREQUIRETLSNotImplemented is returned if REQUIRETLS is used but not enabled.
	ErrREQUIRETLSNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"REQUIRETLS is not implemented"},
	}
	// ErrBINARYMIMENotImplemented is returned if BODY=BINARYMIME is used but not enabled.
	ErrBINARYMIMENotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"BINARYMIME is not implemented"},
	}
	// ErrUnknownBody is returned if the BODY parameter has an unknown value.
	ErrUnknownBody = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unknown BODY value"},
	}
	// ErrRETNotImplemented is returned if RET is used but DSN isn't enabled.
	ErrRETNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"RET is not implemented"},
	}
	// ErrUnknownRET is returned if the RET parameter has an unknown value.
	ErrUnknownRET = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unknown RET value"},
	}
	// ErrENVIDNotImplemented is returned if ENVID is used but DSN isn't enabled.
	ErrENVIDNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"ENVID is not implemented"},
	}
	// ErrMalformedENVID is returned if the ENVID parameter is malformed.
	ErrMalformedENVID = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed ENVID parameter value"},
	}
	// ErrMalformedAUTH is returned if the MAIL AUTH parameter is malformed.
	ErrMalformedAUTH = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed AUTH parameter value"},
	}
	// ErrMalformedAUTHMailbox is returned if the MAIL AUTH parameter contains an invalid mailbox.
	ErrMalformedAUTHMailbox = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed AUTH parameter mailbox"},
	}
	// ErrDELIVERBYNotImplemented is returned if BY is used but DELIVERBY isn't enabled.
	ErrDELIVERBYNotImplemented = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"DELIVERBY is not implemented"},
	}
	// ErrMalformedBY is returned if the BY parameter is malformed.
	ErrMalformedBY = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed BY parameter value"},
	}
	// ErrBYBelowMinimum is returned if the BY parameter is below the configured minimum.
	ErrBYBelowMinimum = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"BY parameter is below server minimum"},
	}
	// ErrMTPRIORITYNotImplemented is returned if MT-PRIORITY is used but not enabled.
	ErrMTPRIORITYNotImplemented = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"MT-PRIORITY is not implemented"},
	}
	// ErrMalformedMTPRIORITY is returned if the MT-PRIORITY parameter is malformed.
	ErrMalformedMTPRIORITY = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed MT-PRIORITY parameter value"},
	}
	// ErrMTPRIORITYOutOfRange is returned if the MT-PRIORITY parameter is outside the valid range.
	ErrMTPRIORITYOutOfRange = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"MT-PRIORITY is outside valid range"},
	}

	// RCPT errors.

	// ErrRcptSyntax is returned if the RCPT command can't be parsed.
	ErrRcptSyntax = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 2},
		Lines:        []string{"Was expecting RCPT arg syntax of TO:<address>"},
	}
	// ErrRcptParams is returned if the RCPT ESMTP parameters can't be parsed.
	ErrRcptParams = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unable to parse RCPT ESMTP parameters"},
	}
	// ErrUnknownRcptArg is returned if RCPT contains an unknown parameter.
	ErrUnknownRcptArg = &Status{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unknown RCPT TO argument"},
	}
	// ErrNonASCIIRecipient is returned if a RCPT address contains non-ASCII
	// characters but SMTPUTF8 was not requested (RFC 6531 section 3.5).
	ErrNonASCIIRecipient = &Status{
		Code:         553,
		EnhancedCode: EnhancedCode{5, 6, 7},
		Lines:        []string{"Non-ASCII addresses not permitted for that recipient"},
	}
	// ErrRcptNotAccepted is returned if the backend rejected RCPT with a non-SMTP error.
	ErrRcptNotAccepted = &Status{
		Code:         451,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"Recipient not accepted"},
	}
	// ErrNOTIFYNotImplemented is returned if NOTIFY is used but DSN isn't enabled.
	ErrNOTIFYNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"NOTIFY is not implemented"},
	}
	// ErrMalformedNOTIFY is returned if the NOTIFY parameter is malformed.
	ErrMalformedNOTIFY = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed NOTIFY parameter value"},
	}
	// ErrORCPTNotImplemented is returned if ORCPT is used but DSN isn't enabled.
	ErrORCPTNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"ORCPT is not implemented"},
	}
	// ErrMalformedORCPT is returned if the ORCPT parameter is malformed.
	ErrMalformedORCPT = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed ORCPT parameter value"},
	}
	// ErrRRVSNotImplemented is returned if RRVS is used but not enabled.
	ErrRRVSNotImplemented = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"RRVS is not implemented"},
	}
	// ErrMalformedRRVS is returned if the RRVS parameter is malformed.
	ErrMalformedRRVS = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed RRVS parameter value"},
	}

	// VRFY errors.

	// ErrVrfySyntax is returned if the VRFY command can't be parsed.
	ErrVrfySyntax = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 2},
		Lines:        []string{"Was expecting <address>"},
	}
	// ErrVrfyParams is returned if the VRFY ESMTP parameters can't be parsed.
	ErrVrfyParams = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unable to parse VRFY ESMTP parameters"},
	}
	// ErrUTF8ReplyRequired is returned if the reply to VRFY requires UTF-8 but
	// SMTPUTF8 was not requested (RFC 6531 section 3.7.4.2).
	ErrUTF8ReplyRequired = &Status{
		Code:         550,
		EnhancedCode: EnhancedCode{5, 6, 8},
		Lines:        []string{"UTF-8 string reply is required, but not permitted by the SMTP client"},
	}

	// AUTH errors.

	// ErrAuthRequired is returned if authentication is required before the command (RFC 4954).
	ErrAuthRequired = &Status{
		Code:         530,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Lines:        []string{"Authentication required"},
	}
	// ErrAuthFailed is returned if the authentication failed.
	ErrAuthFailed = &Status{
		Code:         535,
		EnhancedCode: EnhancedCode{5, 7, 8},
		Lines:        []string{"Authentication failed"},
	}
	// ErrAuthTemporaryFailure is returned if authentication failed with a non-SMTP error (RFC 4954).
	ErrAuthTemporaryFailure = &Status{
		Code:         454,
		EnhancedCode: EnhancedCode{4, 7, 0},
		Lines:        []string{"Temporary authentication failure"},
	}
	// ErrAuthUnsupported is returned if the authentication is not supported.
	ErrAuthUnsupported = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Lines:        []string{"Authentication not supported"},
	}
	// ErrAuthUnknownMechanism is returned if AUTH uses a mechanism that isn't offered (RFC 4954).
	ErrAuthUnknownMechanism = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unrecognized authentication type"},
	}
	// ErrAlreadyAuthenticated is returned if AUTH is sent after a successful authentication.
	ErrAlreadyAuthenticated = &Status{
		Code:         503,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"Already authenticated"},
	}
	// ErrAuthMissingParam is returned if AUTH is sent without a mechanism.
	ErrAuthMissingParam = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Missing parameter"},
	}
	// ErrInvalidBase64 is returned if an AUTH response isn't valid base64.
	ErrInvalidBase64 = &Status{
		Code:         454,
		EnhancedCode: EnhancedCode{4, 7, 0},
		Lines:        []string{"Invalid base64 data"},
	}
	// ErrNoAuthHandler is returned if the session returned no handler for an offered mechanism.
	ErrNoAuthHandler = &Status{
		Code:         451,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"No auth handler received, but mechanism seems valid."},
	}
	// ErrAuthCancelled is returned if the client cancelled the AUTH exchange.
	ErrAuthCancelled = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 0, 0},
		Lines:        []string{"Negotiation cancelled"},
	}

	// STARTTLS errors.

	// ErrAlreadyTLS is returned if STARTTLS is sent on an encrypted connection.
	ErrAlreadyTLS = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"Already running in TLS"},
	}
	// ErrTLSNotSupported is returned if STARTTLS is sent but no TLS config is set.
	ErrTLSNotSupported = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"TLS not supported"},
	}
	// ErrTLSConfig is returned if the TLS config couldn't be retrieved.
	ErrTLSConfig = &Status{
		Code:         451,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"TLS config retrieval failed"},
	}
	// ErrTLSConfigNil is returned if the TLS config retrieval returned nil.
	ErrTLSConfigNil = &Status{
		Code:         451,
		EnhancedCode: EnhancedCode{4, 0, 0},
		Lines:        []string{"TLS config retrieval nil returned"},
	}

	// DATA and BDAT errors.

	// ErrNoRecipients is returned if no recipients are set.
	ErrNoRecipients = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"Missing RCPT TO command."},
	}
	// ErrDataTooLarge is returned if the maximum message size is exceeded.
	ErrDataTooLarge = &Status{
		Code:         552,
		EnhancedCode: EnhancedCode{5, 3, 4},
		Lines:        []string{"Maximum message size exceeded"},
	}
	// ErrDataArgs is returned if DATA is sent with arguments.
	ErrDataArgs = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"DATA command should not have any arguments"},
	}
	// ErrDataBINARYMIME is returned if DATA is used for a BINARYMIME message.
	ErrDataBINARYMIME = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"DATA not allowed for BINARYMIME messages"},
	}
	// ErrCHUNKINGNotImplemented is returned if BDAT is used but CHUNKING isn't enabled.
	ErrCHUNKINGNotImplemented = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"CHUNKING is not implemented"},
	}
	// ErrBDATMissingSize is returned if BDAT is sent without a chunk size.
	ErrBDATMissingSize = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Missing chunk size argument"},
	}
	// ErrBDATTooManyArgs is returned if BDAT is sent with too many arguments.
	ErrBDATTooManyArgs = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Too many arguments"},
	}
	// ErrBDATUnknownArg is returned if BDAT contains an unknown argument.
	ErrBDATUnknownArg = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Unknown BDAT argument"},
	}
	// ErrBDATMalformedSize is returned if the BDAT chunk size is malformed.
	ErrBDATMalformedSize = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"Malformed size argument"},
	}
	// ErrBDATExpected is returned if another command is sent during a BDAT transfer.
	ErrBDATExpected = &Status{
		Code:         501,
		EnhancedCode: EnhancedCode{5, 5, 4},
		Lines:        []string{"BDAT, RSET or QUIT command expected"},
	}
)
