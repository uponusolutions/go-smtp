package client

import (
	"errors"
	"strconv"
	"strings"
)

// validateLine checks to see if a line has CR or LF.
func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("smtp: a line must not contain CR or LF")
	}
	return nil
}

func encodeXtext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		switch {
		case ch >= '!' && ch <= '~' && ch != '+' && ch != '=':
			// printable non-space US-ASCII except '+' and '='
			out.WriteRune(ch)
		default:
			out.WriteRune('+')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
		}
	}
	return out.String()
}

// encodeUTF8AddrUnitext encodes raw string to the utf-8-addr-unitext form in RFC 6533.
func encodeUTF8AddrUnitext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		switch {
		case ch >= '!' && ch <= '~' && ch != '+' && ch != '=':
			// printable non-space US-ASCII except '+' and '='
			out.WriteRune(ch)
		case ch <= '\x7F':
			// other ASCII: CTLs, space and specials
			out.WriteRune('\\')
			out.WriteRune('x')
			out.WriteRune('{')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
			out.WriteRune('}')
		default:
			// UTF-8 non-ASCII
			out.WriteRune(ch)
		}
	}
	return out.String()
}

// encodeUTF8AddrXtext encodes raw string to the utf-8-addr-xtext form in RFC 6533.
func encodeUTF8AddrXtext(raw string) string {
	var out strings.Builder
	out.Grow(len(raw))

	for _, ch := range raw {
		switch {
		case ch >= '!' && ch <= '~' && ch != '+' && ch != '=':
			// printable non-space US-ASCII except '+' and '='
			out.WriteRune(ch)
		default:
			out.WriteRune('\\')
			out.WriteRune('x')
			out.WriteRune('{')
			out.WriteString(strings.ToUpper(strconv.FormatInt(int64(ch), 16)))
			out.WriteRune('}')
		}
	}
	return out.String()
}
