package smtp

import (
	"fmt"
)

type EnhancedCode [3]int

// Smtp specifies the error code, enhanced error code (if any) and
// message returned by the server.
type Smtp struct {
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

func (err *Smtp) Error() string {
	s := fmt.Sprintf("SMTP error %03d", err.Code)
	if err.Message != "" {
		s += ": " + err.Message
	}
	return s
}

func (err *Smtp) Temporary() bool {
	return err.Code/100 == 4
}

var (
	ErrDataTooLarge = &Smtp{
		Code:         552,
		EnhancedCode: EnhancedCode{5, 3, 4},
		Message:      "Maximum message size exceeded",
	}
	ErrAuthFailed = &Smtp{
		Code:         535,
		EnhancedCode: EnhancedCode{5, 7, 8},
		Message:      "Authentication failed",
	}
	ErrAuthRequired = &Smtp{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Please authenticate first",
	}
	ErrAuthUnsupported = &Smtp{
		Code:         502,
		EnhancedCode: EnhancedCode{5, 7, 0},
		Message:      "Authentication not supported",
	}
	ErrAuthUnknownMechanism = &Smtp{
		Code:         504,
		EnhancedCode: EnhancedCode{5, 7, 4},
		Message:      "Unsupported authentication mechanism",
	}
)
