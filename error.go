package smtp

import (
	"fmt"
)

type EnhancedCode [3]int

// SMTPStatus specifies the error code, enhanced error code (if any) and
// message returned by the server.
type SMTPStatus struct {
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

func NewStatus(code int, enhCode EnhancedCode, msg string) *SMTPStatus {
	return &SMTPStatus{
		Code:         code,
		EnhancedCode: enhCode,
		Message:      msg,
	}
}

func (err *SMTPStatus) Error() string {
	s := fmt.Sprintf("SMTP error %03d", err.Code)
	if err.Message != "" {
		s += ": " + err.Message
	}
	return s
}

func (err *SMTPStatus) Temporary() bool {
	return err.Code/100 == 4
}

var (
	ErrBadCommand = &SMTPStatus{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Bad command",
	}
	ErrBadSyntax = &SMTPStatus{
		Code:         500,
		EnhancedCode: EnhancedCode{5, 5, 2},
		Message:      "Bad syntax",
	}
	ErrDataTooLarge = &SMTPStatus{
		Code:         552,
		EnhancedCode: EnhancedCode{5, 3, 4},
		Message:      "Maximum message size exceeded",
	}
	ErrAuthFailed = &SMTPStatus{
		Code:         535,
		EnhancedCode: EnhancedCode{5, 7, 8},
		Message:      "Authentication failed",
	}
	ErrAuthRequired = &SMTPStatus{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Please authenticate first",
	}
	ErrAuthUnsupported = &SMTPStatus{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Authentication not supported",
	}
	ErrAuthUnknownMechanism = &SMTPStatus{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 7, 4},
		Message:      "Unsupported authentication mechanism",
	}
)
