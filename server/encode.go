package server

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
)

func decodeSASLResponse(s string) ([]byte, error) {
	if s == "=" {
		return []byte{}, nil
	}
	return base64.StdEncoding.DecodeString(s)
}

// This regexp matches 'hexchar' token defined in
// https://tools.ietf.org/html/rfc4954#section-8 however it is intentionally
// relaxed by requiring only '+' to be present.  It allows us to detect
// malformed values such as +A or +HH and report them appropriately.
var hexcharRe = regexp.MustCompile(`\+[0-9A-F]?[0-9A-F]?`)

func decodeXtext(val string) (string, error) {
	if !strings.Contains(val, "+") {
		return val, nil
	}

	var replaceErr error
	decoded := hexcharRe.ReplaceAllStringFunc(val, func(match string) string {
		if len(match) != 3 {
			replaceErr = errors.New("incomplete hexchar")
			return ""
		}
		char, err := strconv.ParseInt(match, 16, 8)
		if err != nil {
			replaceErr = err
			return ""
		}

		return string(rune(char))
	})
	if replaceErr != nil {
		return "", replaceErr
	}

	return decoded, nil
}

// This regexp matches 'EmbeddedUnicodeChar' token defined in
// https://datatracker.ietf.org/doc/html/rfc6533.html#section-3
// however it is intentionally relaxed by requiring only '\x{HEX}' to be
// present.  It also matches disallowed characters in QCHAR and QUCHAR defined
// in above.
// So it allows us to detect malformed values and report them appropriately.
var eUOrDCharRe = regexp.MustCompile(`\\x[{][0-9A-F]+[}]|[[:cntrl:] \\+=]`)

// Decodes the utf-8-addr-xtext or the utf-8-addr-unitext form.
func decodeUTF8AddrXtext(val string) (string, error) {
	var replaceErr error
	decoded := eUOrDCharRe.ReplaceAllStringFunc(val, func(match string) string {
		if len(match) == 1 {
			replaceErr = errors.New("disallowed character:" + match)
			return ""
		}

		hexpoint := match[3 : len(match)-1]
		char, err := strconv.ParseUint(hexpoint, 16, 21)
		if err != nil {
			replaceErr = err
			return ""
		}
		switch len(hexpoint) {
		case 2:
			switch {
			// all xtext-specials
			case 0x01 <= char && char <= 0x09 ||
				0x11 <= char && char <= 0x19 ||
				char == 0x10 || char == 0x20 ||
				char == 0x2B || char == 0x3D || char == 0x7F:
			// 2-digit forms
			case char == 0x5C || 0x80 <= char && char <= 0xFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 3-digit forms
		case 3:
			switch {
			case 0x100 <= char && char <= 0xFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 4-digit forms excluding surrogate
		case 4:
			switch {
			case 0x1000 <= char && char <= 0xD7FF:
			case 0xE000 <= char && char <= 0xFFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 5-digit forms
		case 5:
			switch {
			case 0x1_0000 <= char && char <= 0xF_FFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// 6-digit forms
		case 6:
			switch {
			case 0x10_0000 <= char && char <= 0x10_FFFF:
				// This space is intentionally left blank
			default:
				replaceErr = errors.New("illegal hexpoint:" + hexpoint)
				return ""
			}
		// the other invalid forms
		default:
			replaceErr = errors.New("illegal hexpoint:" + hexpoint)
			return ""
		}

		return string(rune(char))
	})
	if replaceErr != nil {
		return "", replaceErr
	}

	return decoded, nil
}

func decodeTypedAddress(val string) (smtp.DSNAddressType, string, error) {
	tv := strings.SplitN(val, ";", 2)
	if len(tv) != 2 || tv[0] == "" || tv[1] == "" {
		return "", "", errors.New("bad address")
	}
	aType, aAddr := strings.ToUpper(tv[0]), tv[1]

	var err error
	switch smtp.DSNAddressType(aType) {
	case smtp.DSNAddressTypeRFC822:
		aAddr, err = decodeXtext(aAddr)
		if err == nil && !textsmtp.IsPrintableASCII(aAddr) {
			err = errors.New("illegal address:" + aAddr)
		}
	case smtp.DSNAddressTypeUTF8:
		aAddr, err = decodeUTF8AddrXtext(aAddr)
	default:
		err = errors.New("unknown address type:" + aType)
	}
	if err != nil {
		return "", "", err
	}

	return smtp.DSNAddressType(aType), aAddr, nil
}
