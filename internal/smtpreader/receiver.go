package smtpreader

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"strconv"
	"strings"

	"github.com/uponusolutions/go-smtp"
)

// Receiver is used as a wrapper around a connection to read from it
type Receiver struct {
	*bufio.Reader
	maxLineLength      int
	lineLengthExceeded bool
}

// NewReceiver creates a new connection wrapper.
func NewReceiver(
	conn io.Reader,
	readerSize int,
	maxLineLength int,
) *Receiver {
	if readerSize == 0 {
		readerSize = 4096 // default
	}

	return &Receiver{
		Reader:             bufio.NewReaderSize(conn, readerSize),
		maxLineLength:      maxLineLength,
		lineLengthExceeded: false,
	}
}

// ReadFullLine reads a single line from r,
// eliding the final \n or \r\n from the returned string.
func (t *Receiver) ReadFullLine() (string, error) {
	line, err := t.readLineSlice()
	return string(line), err
}

// ReadResponse reads a multi-line response of the form:
//
//	code-message line 1
//	code-message line 2
//	...
//	code message line n
//
// where code is a three-digit status code. The first line starts with the
// code and a hyphen. The response is terminated by a line that starts
// with the same code followed by a space. Each line in message is
// separated by a newline (\n).
//
// See page 36 of RFC 959 (https://www.ietf.org/rfc/rfc959.txt) for
// details of another form of response accepted:
//
//	code-message line 1
//	message line 2
//	...
//	code message line n
func (t *Receiver) ReadResponse() (*smtp.Status, error) {
	status, continued, err := t.readFirstCodeLine()
	if err != nil {
		return nil, err
	}

	if err = t.readResponseExtra(status, continued, true); err != nil {
		return nil, err
	}

	return status, nil
}

// ReadResponseValid returns an error if the code does not match expectation.
func (t *Receiver) ReadResponseValid(expectCode int) error {
	status, continued, err := t.readFirstCodeLine()
	if err != nil {
		return err
	}

	unexpected := IsCodeUnexpected(status.Code, expectCode)

	err = t.readResponseExtra(status, continued, unexpected)
	if err != nil {
		return err
	}

	if unexpected {
		return status
	}

	return nil
}

func (t *Receiver) readResponseExtra(status *smtp.Status, continued bool, appendMessage bool) error {
	var message string
	var err error

	encCodePart := status.EnhancedCode.ToPart()
	for continued {
		continued, message, err = t.readExtraCodeLine(strconv.Itoa(status.Code), encCodePart)
		if err != nil {
			return err
		}
		if appendMessage {
			status.Lines = append(status.Lines, message)
		}
	}
	return nil
}

func (t *Receiver) readFirstCodeLine() (*smtp.Status, bool, error) {
	line, err := t.ReadFullLine()
	if err != nil {
		return nil, false, err
	}
	return parseFirstCodeLine(line)
}

func parseFirstCodeLine(line string) (*smtp.Status, bool, error) {
	if len(line) < 4 || line[3] != ' ' && line[3] != '-' {
		return nil, false, textproto.ProtocolError(fmt.Sprintf("short response: %q", line))
	}
	continued := line[3] == '-'
	code, err := strconv.Atoi(line[0:3])
	if err != nil || code < 100 {
		return nil, false, textproto.ProtocolError(fmt.Sprintf("invalid response code: %q", line))
	}
	message := line[4:]

	// all 2xx, 4xx, and 5xx response lines
	if code >= 300 && code < 400 {
		return smtp.NewStatusS(code, smtp.NoEnhancedCode, message), continued, nil
	}

	index := strings.Index(message, " ")
	if index == -1 {
		return smtp.NewStatusS(code, smtp.NoEnhancedCode, message), continued, nil
	}
	enhCode, err := parseEnhancedCode(message[:index])
	if err != nil {
		return smtp.NewStatusS(code, smtp.NoEnhancedCode, message), continued, nil
	}

	message = message[index+1:]

	return smtp.NewStatusS(code, enhCode, message), continued, nil
}

func (t *Receiver) readExtraCodeLine(codeString string, enhCodePart []byte) (continued bool, message string, err error) {
	line, err := t.ReadFullLine()
	if err != nil {
		return false, "", err
	}
	return parseExtraCodeLine(line, codeString, enhCodePart)
}

// parseExtraCodeLine does strict verification like described in RFC 5321
func parseExtraCodeLine(line string, codeString string, enhCodePart []byte) (bool, string, error) {
	if len(line) < 4+len(enhCodePart) ||
		(line[3] != ' ' && line[3] != '-') ||
		line[0:3] != codeString ||
		line[4:(4+len(enhCodePart))] != string(enhCodePart) {
		return false, "", textproto.ProtocolError(fmt.Sprintf("invalid response: %q", line))
	}
	return line[3] == '-', line[4+len(enhCodePart):], nil
}

func parseEnhancedCode(s string) (smtp.EnhancedCode, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return smtp.NoEnhancedCode, errors.New("wrong amount of enhanced code parts")
	}

	code := smtp.NoEnhancedCode
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return smtp.NoEnhancedCode, err
		}
		code[i] = num
	}
	return code, nil
}

// IsCodeUnexpected validates if the code is unexpected.
// If the prefix of the status does not match the digits in expectCode,
// ReadResponse returns with err set to &Error{code, message}.
// For example, if expectCode is 31, an error will be returned if
// the status is not in the range [310,319].
func IsCodeUnexpected(code int, expectCode int) bool {
	return 1 <= expectCode && expectCode < 10 && code/100 != expectCode ||
		10 <= expectCode && expectCode < 100 && code/10 != expectCode ||
		100 <= expectCode && expectCode < 1000 && code != expectCode
}

func (t *Receiver) readLineSlice() ([]byte, error) {
	// If the line limit was exceeded once, the connection shouldn't be used anymore.
	if t.lineLengthExceeded {
		return nil, smtp.ErrTooLongLine
	}

	var line []byte
	for {
		l, more, err := t.ReadLine()
		if err != nil {
			return nil, err
		}

		if t.maxLineLength > 0 && len(l)+len(line) > t.maxLineLength {
			t.lineLengthExceeded = true
			return nil, smtp.ErrTooLongLine
		}

		// Avoid the copy if the first call produced a full line.
		if line == nil && !more {
			return l, nil
		}
		line = append(line, l...)
		if !more {
			break
		}
	}
	return line, nil
}
