package textsmtp

import (
	"bufio"
	"fmt"
	"io"
	"net/textproto"
	"strconv"
	"strings"
)

type Conn struct {
	R    *bufio.Reader
	W    *bufio.Writer
	conn io.ReadWriteCloser
	textproto.Pipeline
}

func NewConn(conn io.ReadWriteCloser) *Conn {
	return &Conn{
		R:    bufio.NewReader(conn),
		W:    bufio.NewWriter(conn),
		conn: conn,
	}
}

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
func (c *Conn) Cmd(format string, args ...any) (id uint, err error) {
	id = c.Next()
	c.StartRequest(id)
	err = c.PrintfLine(format, args...)
	c.EndRequest(id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// PrintfLine writes the formatted output followed by \r\n.
func (t *Conn) PrintfLine(format string, args ...any) error {
	fmt.Fprintf(t.W, format, args...)
	t.W.Write(crnl)
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
func (t *Conn) ReadResponse(expectCode int) (code int, message string, err error) {
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
	return
}

func (t *Conn) ReadCodeLine(expectCode int) (code int, message string, err error) {
	code, continued, message, err := t.readCodeLine(expectCode)
	if err == nil && continued {
		err = textproto.ProtocolError("unexpected multi-line response: " + message)
	}
	return
}

func (t *Conn) readCodeLine(expectCode int) (code int, continued bool, message string, err error) {
	line, err := t.ReadLine()
	if err != nil {
		return
	}
	return parseCodeLine(line, expectCode)
}

func parseCodeLine(line string, expectCode int) (code int, continued bool, message string, err error) {
	if len(line) < 4 || line[3] != ' ' && line[3] != '-' {
		err = textproto.ProtocolError("short response: " + line)
		return
	}
	continued = line[3] == '-'
	code, err = strconv.Atoi(line[0:3])
	if err != nil || code < 100 {
		err = textproto.ProtocolError("invalid response code: " + line)
		return
	}
	message = line[4:]
	if 1 <= expectCode && expectCode < 10 && code/100 != expectCode ||
		10 <= expectCode && expectCode < 100 && code/10 != expectCode ||
		100 <= expectCode && expectCode < 1000 && code != expectCode {
		err = &textproto.Error{Code: code, Msg: message}
	}
	return
}

// ReadLine reads a single line from r,
// eliding the final \n or \r\n from the returned string.
func (t *Conn) ReadLine() (string, error) {
	line, err := t.readLineSlice()
	return string(line), err
}

// ReadLineBytes is like ReadLine but returns a []byte instead of a string.
func (t *Conn) ReadLineBytes() ([]byte, error) {
	line, err := t.readLineSlice()
	if line != nil {
		buf := make([]byte, len(line))
		copy(buf, line)
		line = buf
	}
	return line, err
}

func (t *Conn) readLineSlice() ([]byte, error) {
	var line []byte
	for {
		l, more, err := t.R.ReadLine()
		if err != nil {
			return nil, err
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

// Close closes the connection.
func (c *Conn) Close() error {
	return c.conn.Close()
}
