package textsmtp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"strconv"
	"strings"
)

// Textproto is used as a wrapper around a connection to read and write to it
type Textproto struct {
	R                  *bufio.Reader
	W                  *bufio.Writer
	conn               io.ReadWriteCloser
	maxLineLength      int
	lineLengthExceeded bool
	textproto.Pipeline
}

// NewTextproto creates a new connection wrapper.
func NewTextproto(
	conn io.ReadWriteCloser,
	readerSize int,
	writerSize int,
	maxLineLength int,
) *Textproto {
	if readerSize == 0 {
		readerSize = 4096 // default
	}

	if writerSize == 0 {
		writerSize = 4096 // default
	}

	return &Textproto{
		R:                  bufio.NewReaderSize(conn, readerSize),
		W:                  bufio.NewWriterSize(conn, writerSize),
		conn:               conn,
		maxLineLength:      maxLineLength,
		lineLengthExceeded: false,
	}
}

// ErrTooLongLine occurs if the smtp line is too long.
var ErrTooLongLine = errors.New("smtp: too long a line in input stream")

// Cmd is a convenience method that sends a command after
// waiting its turn in the pipeline. The command text is the
// result of formatting format with args and appending \r\n.
// Cmd returns the id of the command, for use with StartResponse and EndResponse.
//
// For example, a client might run a HELP command that returns a dot-body
// by using:
//
//	id, err := c.Cmd("HELP")
//	if err != nil {
//		return nil, err
//	}
//
//	c.StartResponse(id)
//	defer c.EndResponse(id)
//
//	if _, _, err = c.ReadCodeLine(110); err != nil {
//		return nil, err
//	}
//	text, err := c.ReadDotBytes()
//	if err != nil {
//		return nil, err
//	}
//	return c.ReadCodeLine(250)
func (t *Textproto) Cmd(format string, args ...any) (id uint, err error) {
	id = t.Next()
	t.StartRequest(id)
	err = t.PrintfLine(format, args...)
	t.EndRequest(id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// PrintfLine writes the formatted output followed by \r\n.
func (t *Textproto) PrintfLine(format string, args ...any) error {
	if _, err := fmt.Fprintf(t.W, format, args...); err != nil {
		return err
	}

	if _, err := t.W.Write(crnl); err != nil {
		return err
	}
	return t.W.Flush()
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
//
// If the prefix of the status does not match the digits in expectCode,
// ReadResponse returns with err set to &Error{code, message}.
// For example, if expectCode is 31, an error will be returned if
// the status is not in the range [310,319].
//
// An expectCode <= 0 disables the check of the status code.
func (t *Textproto) ReadResponse(expectCode int) (code int, message string, err error) {
	code, continued, message, err := t.readCodeLine(expectCode)
	multi := continued
	for continued {
		line, err := t.ReadLine()
		if err != nil {
			return 0, "", err
		}

		var code2 int
		var moreMessage string
		code2, continued, moreMessage, err = parseCodeLine(line, 0)
		if err != nil || code2 != code {
			message += "\n" + strings.TrimRight(line, "\r\n")
			continued = true
			continue
		}
		message += "\n" + moreMessage
	}
	if err != nil && multi && message != "" {
		// replace one line error message with all lines (full message)
		err = &textproto.Error{Code: code, Msg: message}
	}
	return code, message, err
}

// ReadCodeLine reads a code line.
func (t *Textproto) ReadCodeLine(expectCode int) (int, string, error) {
	code, continued, message, err := t.readCodeLine(expectCode)
	if err == nil && continued {
		err = textproto.ProtocolError("unexpected multi-line response: " + message)
	}
	return code, message, err
}

func (t *Textproto) readCodeLine(expectCode int) (code int, continued bool, message string, err error) {
	line, err := t.ReadLine()
	if err != nil {
		return code, continued, message, err
	}
	return parseCodeLine(line, expectCode)
}

func parseCodeLine(line string, expectCode int) (code int, continued bool, message string, err error) {
	if len(line) < 4 || line[3] != ' ' && line[3] != '-' {
		err = textproto.ProtocolError("short response: " + line)
		return code, continued, message, err
	}
	continued = line[3] == '-'
	code, err = strconv.Atoi(line[0:3])
	if err != nil || code < 100 {
		err = textproto.ProtocolError("invalid response code: " + line)
		return code, continued, message, err
	}
	message = line[4:]
	if 1 <= expectCode && expectCode < 10 && code/100 != expectCode ||
		10 <= expectCode && expectCode < 100 && code/10 != expectCode ||
		100 <= expectCode && expectCode < 1000 && code != expectCode {
		err = &textproto.Error{Code: code, Msg: message}
	}
	return code, continued, message, err
}

// ReadLine reads a single line from r,
// eliding the final \n or \r\n from the returned string.
func (t *Textproto) ReadLine() (string, error) {
	line, err := t.readLineSlice()
	return string(line), err
}

func (t *Textproto) readLineSlice() ([]byte, error) {
	// If the line limit was exceeded once, the connection shouldn't be used anymore.
	if t.lineLengthExceeded {
		return nil, ErrTooLongLine
	}

	var line []byte
	for {
		l, more, err := t.R.ReadLine()
		if err != nil {
			return nil, err
		}

		if t.maxLineLength > 0 && len(l)+len(line) > t.maxLineLength {
			t.lineLengthExceeded = true
			return nil, ErrTooLongLine
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

// Replace conn.
func (t *Textproto) Replace(conn io.ReadWriteCloser) {
	t.conn = conn
	t.R.Reset(t.conn)
	t.W.Reset(t.conn)
}

// Close closes the connection.
func (t *Textproto) Close() error {
	return t.conn.Close()
}
