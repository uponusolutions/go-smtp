package parse

import (
	"fmt"
	"strings"
)

// CutPrefixFold is a version of strings.CutPrefix which is case-insensitive.
func CutPrefixFold(s, prefix string) (string, bool) {
	if len(s) < len(prefix) || !strings.EqualFold(s[:len(prefix)], prefix) {
		return "", false
	}
	return s[len(prefix):], true
}

func IsPrintableASCII(val string) bool {
	for _, ch := range val {
		if ch < ' ' || '~' < ch {
			return false
		}
	}
	return true
}

func Cmd(line string) (cmd string, arg string, err error) {
	line = strings.TrimRight(line, "\r\n")

	l := len(line)
	switch {
	case strings.HasPrefix(strings.ToUpper(line), "STARTTLS"):
		return "STARTTLS", "", nil
	case l == 0:
		return "", "", nil
	case l < 4:
		return "", "", fmt.Errorf("command too short: %q", line)
	case l == 4:
		return strings.ToUpper(line), "", nil
	case l == 5:
		// Too long to be only command, too short to have args
		return "", "", fmt.Errorf("mangled command: %q", line)
	}

	// If we made it here, command is long enough to have args
	if line[4] != ' ' {
		// There wasn't a space after the command?
		return "", "", fmt.Errorf("mangled command: %q", line)
	}

	return strings.ToUpper(line[0:4]), strings.TrimSpace(line[5:]), nil
}

// Takes the arguments proceeding a command and files them
// into a map[string]string after uppercasing each key.  Sample arg
// string:
//
//	" BODY=8BITMIME SIZE=1024 SMTPUTF8"
//
// The leading space is mandatory.
func Args(s string) (map[string]string, error) {
	argMap := map[string]string{}
	for _, arg := range strings.Fields(s) {
		m := strings.Split(arg, "=")
		switch len(m) {
		case 2:
			argMap[strings.ToUpper(m[0])] = m[1]
		case 1:
			argMap[strings.ToUpper(m[0])] = ""
		default:
			return nil, fmt.Errorf("failed to parse arg string: %q", arg)
		}
	}
	return argMap, nil
}

func HelloArgument(arg string) (string, error) {
	domain := arg
	if idx := strings.IndexRune(arg, ' '); idx >= 0 {
		domain = arg[:idx]
	}
	if domain == "" {
		return "", fmt.Errorf("invalid domain")
	}
	return domain, nil
}

// Parser parses command arguments defined in RFC 5321 section 4.1.2.
type Parser struct {
	S string
}

func (p *Parser) peekByte() (byte, bool) {
	if len(p.S) == 0 {
		return 0, false
	}
	return p.S[0], true
}

func (p *Parser) readByte() (byte, bool) {
	ch, ok := p.peekByte()
	if ok {
		p.S = p.S[1:]
	}
	return ch, ok
}

func (p *Parser) acceptByte(ch byte) bool {
	got, ok := p.peekByte()
	if !ok || got != ch {
		return false
	}
	p.readByte()
	return true
}

func (p *Parser) expectByte(ch byte) error {
	if !p.acceptByte(ch) {
		if len(p.S) == 0 {
			return fmt.Errorf("expected '%v', got EOF", string(ch))
		} else {
			return fmt.Errorf("expected '%v', got '%v'", string(ch), string(p.S[0]))
		}
	}
	return nil
}

func (p *Parser) ReversePath() (string, error) {
	if strings.HasPrefix(p.S, "<>") {
		p.S = strings.TrimPrefix(p.S, "<>")
		return "", nil
	}
	return p.Path()
}

func (p *Parser) Path() (string, error) {
	hasBracket := p.acceptByte('<')
	if p.acceptByte('@') {
		i := strings.IndexByte(p.S, ':')
		if i < 0 {
			return "", fmt.Errorf("malformed a-d-l")
		}
		p.S = p.S[i+1:]
	}
	mbox, err := p.Mailbox()
	if err != nil {
		return "", fmt.Errorf("in mailbox: %v", err)
	}
	if hasBracket {
		if err := p.expectByte('>'); err != nil {
			return "", err
		}
	}
	return mbox, nil
}

func (p *Parser) Mailbox() (string, error) {
	localPart, err := p.localPart()
	if err != nil {
		return "", fmt.Errorf("in local-part: %v", err)
	} else if localPart == "" {
		return "", fmt.Errorf("local-part is empty")
	}

	if err := p.expectByte('@'); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(localPart)
	sb.WriteByte('@')

	for {
		ch, ok := p.peekByte()
		if !ok {
			break
		}
		if ch == ' ' || ch == '\t' || ch == '>' {
			break
		}
		p.readByte()
		sb.WriteByte(ch)
	}

	if strings.HasSuffix(sb.String(), "@") {
		return "", fmt.Errorf("domain is empty")
	}

	return sb.String(), nil
}

func (p *Parser) localPart() (string, error) {
	var sb strings.Builder

	if p.acceptByte('"') { // quoted-string
		for {
			ch, ok := p.readByte()
			switch ch {
			case '\\':
				ch, ok = p.readByte()
			case '"':
				return sb.String(), nil
			}
			if !ok {
				return "", fmt.Errorf("malformed quoted-string")
			}
			sb.WriteByte(ch)
		}
	} else { // dot-string
		for {
			ch, ok := p.peekByte()
			if !ok {
				return sb.String(), nil
			}
			switch ch {
			case '@':
				return sb.String(), nil
			case '(', ')', '<', '>', '[', ']', ':', ';', '\\', ',', '"', ' ', '\t':
				return "", fmt.Errorf("malformed dot-string")
			}
			p.readByte()
			sb.WriteByte(ch)
		}
	}
}
