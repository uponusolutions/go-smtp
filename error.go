package smtp

import (
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
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
// and 5xx responses. This constant is exported for use by extensions, you
// should probably use EnhancedCodeNotSet instead.
var NoEnhancedCode = EnhancedCode{-1, -1, -1}

// EnhancedCodeNotSet is a nil value of EnhancedCode field in smtp, used
// to indicate that backend failed to provide enhanced status code. X.0.0 will
// be used (X is derived from error code).
var EnhancedCodeNotSet = EnhancedCode{0, 0, 0}

// ToPart returns the part of the string after the code
// which is defined by the enhanced code with the trailing whitespace.
// E.g. "5.1.1 "
func (enhCode EnhancedCode) ToPart(code int) string {
	if enhCode == NoEnhancedCode {
		return ""
	}

	// All responses must include an enhanced code, if it is missing - use
	// a generic code X.0.0.
	if enhCode == EnhancedCodeNotSet {
		cat := code / 100
		switch cat {
		case 2, 4, 5:
			return strconv.Itoa(cat) + ".0.0 "
		default:
			return ""
		}
	}
	return strconv.Itoa(enhCode[0]) + "." +
		strconv.Itoa(enhCode[1]) + "." +
		strconv.Itoa(enhCode[2]) + " "
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
	base := fmt.Sprintf("SMTP error %03d", s.Code)
	if s.EnhancedCode != NoEnhancedCode && s.EnhancedCode != EnhancedCodeNotSet {
		base += fmt.Sprintf(" %d.%d.%d", s.EnhancedCode[0], s.EnhancedCode[1], s.EnhancedCode[2])
	}
	if len(s.Lines) > 0 {
		return base + ": " + s.Text()
	}
	return base
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

// writeLine writes a single reply line. last selects the terminating form.
func writeLine(w io.Writer, code []byte, enhCode []byte, message string, last bool) (i int, err error) {
	var p int

	i, err = w.Write(code)
	if err != nil {
		return i, err
	}

	if last {
		// RFC 5321 permits omitting the space when there is no text, but
		// RFC 4954 requires it for the 334 challenge and net/textproto
		// rejects any reply shorter than four bytes. Always emit it.
		p, err = w.Write([]byte{' '})
	} else {
		p, err = w.Write([]byte{'-'})
	}
	i += p
	if err != nil {
		return i, err
	}

	p, err = w.Write(enhCode)
	i += p
	if err != nil {
		return i, err
	}

	p, err = w.Write([]byte(strings.TrimSuffix(message, "\r")))
	i += p
	if err != nil {
		return i, err
	}

	p, err = w.Write([]byte{'\r', '\n'})
	i += p

	return i, err
}

// WriteTo writes the smtp status reply.
func (s *Status) WriteTo(w io.Writer) (int64, error) {
	codeString := []byte(strconv.Itoa(s.Code))
	enhCodeString := []byte(s.EnhancedCode.ToPart(s.Code))

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

var (
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
	// ErrConnection is returned if a connection error occurs.
	ErrConnection = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 4, 0},
		Lines:        []string{"Connection error, sorry"},
	}
	// ErrDataTooLarge is returned if the maximum message size is exceeded.
	ErrDataTooLarge = &Status{
		Code:         552,
		EnhancedCode: EnhancedCode{5, 3, 4},
		Lines:        []string{"Maximum message size exceeded"},
	}
	// ErrAuthFailed is returned if the authentication failed.
	ErrAuthFailed = &Status{
		Code:         535,
		EnhancedCode: EnhancedCode{5, 7, 8},
		Lines:        []string{"Authentication failed"},
	}
	// ErrAuthRequired is returned if the authentication is required.
	ErrAuthRequired = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Lines:        []string{"Please authenticate first"},
	}
	// ErrAuthUnsupported is returned if the authentication is not supported.
	ErrAuthUnsupported = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Lines:        []string{"Authentication not supported"},
	}
	// ErrAuthUnknownMechanism is returned if the authentication unsupported.
	ErrAuthUnknownMechanism = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 7, 4},
		Lines:        []string{"Unsupported authentication mechanism"},
	}
	// ErrNoRecipients is returned if no recipients are set.
	ErrNoRecipients = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 5, 1},
		Lines:        []string{"Missing RCPT TO command."},
	}
)
