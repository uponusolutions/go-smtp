package smtp

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

// EnhancedCode is the SMTP enhanced code
type EnhancedCode [3]int

// ToResponsePart returns the part of the response defined by enhanced code.
func (enhCode EnhancedCode) ToResponsePart(code int) string {
	if enhCode == NoEnhancedCode {
		return ""
	}

	// All responses must include an enhanced code, if it is missing - use
	// a generic code X.0.0.
	if enhCode == EnhancedCodeNotSet {
		cat := code / 100
		switch cat {
		case 2, 4, 5:
			return strconv.FormatInt(int64(cat), 10) + ".0.0 "
		default:
			return ""
		}
	}
	return strconv.FormatInt(int64(enhCode[0]), 10) + "." +
		strconv.FormatInt(int64(enhCode[1]), 10) + "." +
		strconv.FormatInt(int64(enhCode[2]), 10) + " "
}

// statusBase specifies the error code, enhanced error code (if any)
type statusBase struct {
	Code         int
	EnhancedCode EnhancedCode
}

// Status specifies the error code, enhanced error code (if any) and
// message returned by the server.
type Status struct {
	statusBase
	Message string
}

// StatusMultiline specifies the error code, enhanced error code (if any) and
// message as stream returned by the server.
type StatusMultiline struct {
	statusBase
	Message iter.Seq2[string, bool]
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

// NewStatus creates a new status.
func NewStatus(code int, enhCode EnhancedCode, msg string) *Status {
	return &Status{
		statusBase: statusBase{
			Code:         code,
			EnhancedCode: enhCode,
		},
		Message: msg,
	}
}

// NewStatusMultiline creates a new status multiline.
func NewStatusMultiline(code int, enhCode EnhancedCode, msg iter.Seq2[string, bool]) *StatusMultiline {
	return &StatusMultiline{
		statusBase: statusBase{
			Code:         code,
			EnhancedCode: enhCode,
		},
		Message: msg,
	}
}

// Error returns a error string.
func (err *statusBase) Error() string {
	return fmt.Sprintf("SMTP error %03d", err.Code)
}

// Error returns a error string.
func (err *Status) Error() string {
	if err.Message != "" {
		return err.statusBase.Error() + ": " + err.Message
	}
	return err.statusBase.Error()
}

// Error returns a error string.
func (err *StatusMultiline) Error() string {
	if err.Message != nil {
		sb := strings.Builder{}
		for message, hasNextLine := range err.Message {
			sb.WriteString(message)
			if hasNextLine {
				sb.WriteByte('\n')
			}
		}
		message := sb.String()
		if message != "" {
			return err.statusBase.Error() + ": " + message
		}
	}
	return err.statusBase.Error()
}

// Positive returns true if the status code is 2xx.
func (err *statusBase) Positive() bool {
	return err.Code/100 == 2
}

// Temporary returns true if the status code is 4xx.
func (err *statusBase) Temporary() bool {
	return err.Code/100 == 4
}

// Permanent returns true if the status code is 5xx.
func (err *statusBase) Permanent() bool {
	return err.Code/100 == 5
}

var (
	// Reset is returned by Reader passed to Data function if client does not
	// send another BDAT command and instead issues RSET command.
	Reset = &Status{
		statusBase: statusBase{
			Code:         250,
			EnhancedCode: EnhancedCode{2, 0, 0},
		},
		Message: "Session reset",
	}
	// VRFY default return.
	VRFY = &Status{
		statusBase: statusBase{
			Code:         252,
			EnhancedCode: EnhancedCode{2, 5, 0},
		},
		Message: "Cannot VRFY user, but will accept message",
	}
	// Noop default return.
	Noop = &Status{
		statusBase: statusBase{
			Code:         250,
			EnhancedCode: EnhancedCode{2, 0, 0},
		},
		Message: "I have successfully done nothing",
	}
	// Quit is returned by Reader passed to Data function if client does not
	// send another BDAT command and instead issues QUIT command.
	Quit = &Status{
		statusBase: statusBase{
			Code:         221,
			EnhancedCode: EnhancedCode{2, 0, 0},
		},
		Message: "Bye",
	}
	// ErrConnection is returned if a connection error occurs.
	ErrConnection = &Status{
		statusBase: statusBase{
			Code:         421,
			EnhancedCode: EnhancedCode{4, 4, 0},
		},
		Message: "Connection error, sorry",
	}
	// ErrDataTooLarge is returned if the maximum message size is exceeded.
	ErrDataTooLarge = &Status{
		statusBase: statusBase{
			Code:         552,
			EnhancedCode: EnhancedCode{5, 3, 4},
		},
		Message: "Maximum message size exceeded",
	}
	// ErrAuthFailed is returned if the authentication failed.
	ErrAuthFailed = &Status{
		statusBase: statusBase{
			Code:         535,
			EnhancedCode: EnhancedCode{5, 7, 8},
		},
		Message: "Authentication failed",
	}
	// ErrAuthRequired is returned if the authentication is required.
	ErrAuthRequired = &Status{
		statusBase: statusBase{
			Code:         502,
			EnhancedCode: EnhancedCode{5, 7, 0},
		},
		Message: "Please authenticate first",
	}
	// ErrAuthUnsupported is returned if the authentication is not supported.
	ErrAuthUnsupported = &Status{
		statusBase: statusBase{
			Code:         502,
			EnhancedCode: EnhancedCode{5, 7, 0},
		},
		Message: "Authentication not supported",
	}
	// ErrAuthUnknownMechanism is returned if the authentication unsupported.
	ErrAuthUnknownMechanism = &Status{
		statusBase: statusBase{
			Code:         504,
			EnhancedCode: EnhancedCode{5, 7, 4},
		},
		Message: "Unsupported authentication mechanism",
	}
	// ErrNoRecipients is returned if no recipients are set.
	ErrNoRecipients = &Status{
		statusBase: statusBase{
			Code:         502,
			EnhancedCode: EnhancedCode{5, 5, 1},
		},
		Message: "Missing RCPT TO command.",
	}
)
