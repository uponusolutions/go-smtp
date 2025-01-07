package textsmtp

import (
	"errors"
	"strconv"
	"strings"

	"github.com/uponusolutions/go-smtp"
)

func IsPrintableASCII(val string) bool {
	for _, ch := range val {
		if ch < ' ' || '~' < ch {
			return false
		}
	}
	return true
}

func EncodeXtext(raw string) string {
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

// EncodeUTF8AddrUnitext encodes raw string to the utf-8-addr-unitext form in RFC 6533.
func EncodeUTF8AddrUnitext(raw string) string {
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

// EncodeUTF8AddrXtext encodes raw string to the utf-8-addr-xtext form in RFC 6533.
func EncodeUTF8AddrXtext(raw string) string {
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

func CheckNotifySet(values []smtp.DSNNotify) error {
	if len(values) == 0 {
		return errors.New("malformed NOTIFY parameter value")
	}

	seen := map[smtp.DSNNotify]struct{}{}
	for _, val := range values {
		switch val {
		case smtp.DSNNotifyNever, smtp.DSNNotifyDelayed, smtp.DSNNotifyFailure, smtp.DSNNotifySuccess:
			if _, ok := seen[val]; ok {
				return errors.New("malformed NOTIFY parameter value")
			}
		default:
			return errors.New("malformed NOTIFY parameter value")
		}
		seen[val] = struct{}{}
	}
	if _, ok := seen[smtp.DSNNotifyNever]; ok && len(seen) > 1 {
		return errors.New("malformed NOTIFY parameter value")
	}

	return nil
}
