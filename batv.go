package smtp

import (
	"errors"
	"regexp"
)

const (
	// BATVRegEx defines the default BATV regular expression.
	// See https://www.ietf.org/archive/id/draft-levine-smtp-batv-01.html#metasyn
	BATVRegEx = "(?i)[a-zA-Z0-9]*=[a-zA-Z0-9]*=(.*@.*)"
)

// CompiledBATVRx is the compiled BATVRegEx.
var CompiledBATVRx = regexp.MustCompile(BATVRegEx)

// ParseBATV parses src with regex string ex to extract a BATV address.
// On error or when BATV extration is not possible/needed src is returned.
// This is a wrapper for ParseBATVEx.
// For performance it is recommented to use ParseBATVEx,
// with a compiled regex.
func ParseBATV(ex, src string) (string, error) {
	if ex == "" {
		return src, errors.New("regular expression is empty")
	}

	re, err := regexp.Compile(ex)
	if err != nil {
		return src, err
	}

	return ParseBATVEx(re, src)
}

// ParseBATVEx parses src with regex re to extract a BATV address.
// On error or when BATV extration is not possible/needed src is returned.
func ParseBATVEx(re *regexp.Regexp, src string) (string, error) {
	if re == nil {
		return src, errors.New("regular expression is nil")
	}

	matched := re.FindStringSubmatch(src)
	if len(matched) != 2 {
		return src, nil
	}

	return matched[1], nil
}
