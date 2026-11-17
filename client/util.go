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

// limit returns a numeric limit advertised by the LIMITS extension (RFC 9422).
//
// RFC 9422 section 3.7: if the LIMITS parameters are syntactically invalid,
// they are ignored entirely; if only the value of a limit is invalid, only that
// limit is ignored. Numeric limits are 1 to 999999 without leading zeros.
func limit(ext map[string]string, name string) (int, bool) {
	param, ok := ext["LIMITS"]
	if !ok {
		return 0, false
	}

	var value string
	found := false
	// limits-param = limit-name-value 0*[SP limit-name-value]
	for _, nv := range strings.Split(param, " ") {
		n, v, hasValue := strings.Cut(nv, "=")
		if !validLimitName(n) || (hasValue && !validLimitValue(v)) {
			return 0, false
		}
		if !found && strings.EqualFold(n, name) {
			value, found = v, true
		}
	}

	if !found || len(value) < 1 || len(value) > 6 || value[0] < '1' || value[0] > '9' {
		return 0, false
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return n, true
}

// validLimitName checks limit-name = 1*(ALPHA / DIGIT / "-" / "_").
func validLimitName(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		alpha := r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
		digit := r >= '0' && r <= '9'
		if !alpha && !digit && r != '-' && r != '_' {
			return false
		}
	}
	return true
}

// validLimitValue checks limit-value = 1*(%x21-3A / %x3C-7E), any VCHAR except ";".
func validLimitValue(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 0x21 || s[i] > 0x7e || s[i] == ';' {
			return false
		}
	}
	return true
}
