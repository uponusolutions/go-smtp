package smtp

import (
	"fmt"
	"iter"
	"strings"
)

// EnhancedCode is the SMTP enhanced code
type EnhancedCode [3]int

// Status specifies the error code, enhanced error code (if any)
type Status struct {
	Code         int
	EnhancedCode EnhancedCode
}

// StatusSingle specifies the error code, enhanced error code (if any) and
// message returned by the server.
type StatusSingle struct {
	Status
	Message string
}

// StatusMulti specifies the error code, enhanced error code (if any) and
// message as stream returned by the server.
type StatusMulti struct {
	Status
	Message iter.Seq[string]
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
func NewStatus(code int, enhCode EnhancedCode, msg string) *StatusSingle {
	return &StatusSingle{
		Status: Status{
			Code:         code,
			EnhancedCode: enhCode,
		},
		Message: msg,
	}
}

// NewStatusMultiline creates a new status multiline.
// You should only use this, if you are return more then one line and you must set Message.
func NewStatusMultiline(code int, enhCode EnhancedCode, msg iter.Seq[string]) *StatusMulti {
	return &StatusMulti{
		Status: Status{
			Code:         code,
			EnhancedCode: enhCode,
		},
		Message: msg,
	}
}

// Error returns a error string.
func (err *Status) Error() string {
	return fmt.Sprintf("SMTP error %03d", err.Code)
}

// Error returns a error string.
func (err *StatusSingle) Error() string {
	if err.Message != "" {
		return err.Status.Error() + ": " + err.Message
	}
	return err.Status.Error()
}

// Error returns a error string.
func (err *StatusMulti) Error() string {
	if err.Message != nil {
		sb := strings.Builder{}
		first := true
		for message := range err.Message {
			if first {
				first = false
			} else {
				sb.WriteByte('\n')
			}
			sb.WriteString(message)
		}
		message := sb.String()
		if message != "" {
			return err.Status.Error() + ": " + message
		}
	}
	return err.Status.Error()
}

// Positive returns true if the status code is 2xx.
func (err *Status) Positive() bool {
	return err.Code/100 == 2
}

// Temporary returns true if the status code is 4xx.
func (err *Status) Temporary() bool {
	return err.Code/100 == 4
}

// Permanent returns true if the status code is 5xx.
func (err *Status) Permanent() bool {
	return err.Code/100 == 5
}

var (
	// Reset is returned by Reader passed to Data function if client does not
	// send another BDAT command and instead issues RSET command.
	Reset = &StatusSingle{
		Status: Status{
			Code:         250,
			EnhancedCode: EnhancedCode{2, 0, 0},
		},
		Message: "Session reset",
	}
	// VRFY default return.
	VRFY = &StatusSingle{
		Status: Status{
			Code:         252,
			EnhancedCode: EnhancedCode{2, 5, 0},
		},
		Message: "Cannot VRFY user, but will accept message",
	}
	// Noop default return.
	Noop = &StatusSingle{
		Status: Status{
			Code:         250,
			EnhancedCode: EnhancedCode{2, 0, 0},
		},
		Message: "I have successfully done nothing",
	}
	// Quit is returned by Reader passed to Data function if client does not
	// send another BDAT command and instead issues QUIT command.
	Quit = &StatusSingle{
		Status: Status{
			Code:         221,
			EnhancedCode: EnhancedCode{2, 0, 0},
		},
		Message: "Bye",
	}
	// ErrConnection is returned if a connection error occurs.
	ErrConnection = &StatusSingle{
		Status: Status{
			Code:         421,
			EnhancedCode: EnhancedCode{4, 4, 0},
		},
		Message: "Connection error, sorry",
	}
	// ErrDataTooLarge is returned if the maximum message size is exceeded.
	ErrDataTooLarge = &StatusSingle{
		Status: Status{
			Code:         552,
			EnhancedCode: EnhancedCode{5, 3, 4},
		},
		Message: "Maximum message size exceeded",
	}
	// ErrAuthFailed is returned if the authentication failed.
	ErrAuthFailed = &StatusSingle{
		Status: Status{
			Code:         535,
			EnhancedCode: EnhancedCode{5, 7, 8},
		},
		Message: "Authentication failed",
	}
	// ErrAuthRequired is returned if the authentication is required.
	ErrAuthRequired = &StatusSingle{
		Status: Status{
			Code:         502,
			EnhancedCode: EnhancedCode{5, 7, 0},
		},
		Message: "Please authenticate first",
	}
	// ErrAuthUnsupported is returned if the authentication is not supported.
	ErrAuthUnsupported = &StatusSingle{
		Status: Status{
			Code:         502,
			EnhancedCode: EnhancedCode{5, 7, 0},
		},
		Message: "Authentication not supported",
	}
	// ErrAuthUnknownMechanism is returned if the authentication unsupported.
	ErrAuthUnknownMechanism = &StatusSingle{
		Status: Status{
			Code:         504,
			EnhancedCode: EnhancedCode{5, 7, 4},
		},
		Message: "Unsupported authentication mechanism",
	}
	// ErrNoRecipients is returned if no recipients are set.
	ErrNoRecipients = &StatusSingle{
		Status: Status{
			Code:         502,
			EnhancedCode: EnhancedCode{5, 5, 1},
		},
		Message: "Missing RCPT TO command.",
	}
)
