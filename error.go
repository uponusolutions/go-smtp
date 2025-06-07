package smtp

import (
	"fmt"
	"strings"
)

// EnhancedCode is the SMTP enhanced code
type EnhancedCode [3]int

// Status specifies the error code, enhanced error code (if any) and
// message returned by the server.
type Status struct {
	Code         int
	EnhancedCode EnhancedCode
	Message      string
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
func NewStatus(code int, enhCode EnhancedCode, msg ...string) *Status {
	return &Status{
		Code:         code,
		EnhancedCode: enhCode,
		Message:      strings.Join(msg, "\n"),
	}
}

// Error returns a error string.
func (err *Status) Error() string {
	s := fmt.Sprintf("SMTP error %03d", err.Code)
	if err.Message != "" {
		s += ": " + err.Message
	}
	return s
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
	// Reset is returned by Reader pased to Data function if client does not
	// send another BDAT command and instead issues RSET command.
	Reset = &Status{
		Code:         250,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Message:      "Session reset",
	}
	// VRFY default return.
	VRFY = &Status{
		Code:         252,
		EnhancedCode: EnhancedCode{2, 5, 0},
		Message:      "Cannot VRFY user, but will accept message",
	}
	// Noop default return.
	Noop = &Status{
		Code:         250,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Message:      "I have successfully done nothing",
	}
	// Quit is returned by Reader pased to Data function if client does not
	// send another BDAT command and instead issues QUIT command.
	Quit = &Status{
		Code:         221,
		EnhancedCode: EnhancedCode{2, 0, 0},
		Message:      "Bye",
	}
	// ErrConnection is returned if a connection error occurs.
	ErrConnection = &Status{
		Code:         421,
		EnhancedCode: EnhancedCode{4, 4, 0},
		Message:      "Connection error, sorry",
	}
	// ErrDataTooLarge is returned if the maximum message size is exceeded.
	ErrDataTooLarge = &Status{
		Code:         552,
		EnhancedCode: EnhancedCode{5, 3, 4},
		Message:      "Maximum message size exceeded",
	}
	// ErrAuthFailed is returned if the authentication failed.
	ErrAuthFailed = &Status{
		Code:         535,
		EnhancedCode: EnhancedCode{5, 7, 8},
		Message:      "Authentication failed",
	}
	// ErrAuthRequired is returned if the authentication is required.
	ErrAuthRequired = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Please authenticate first",
	}
	// ErrAuthUnsupported is returned if the authentication is not supported.
	ErrAuthUnsupported = &Status{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Authentication not supported",
	}
	// ErrAuthUnknownMechanism is returned if the authentication unsupported..
	ErrAuthUnknownMechanism = &Status{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 7, 4},
		Message:      "Unsupported authentication mechanism",
	}
)
