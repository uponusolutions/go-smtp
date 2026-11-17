package parse

import (
	"errors"
	"fmt"
	"net/netip"
	"strings"
	"unicode/utf8"

	"github.com/uponusolutions/go-smtp"
)

// CutPrefixFold is a version of strings.CutPrefix which is case-insensitive.
func CutPrefixFold(s, prefix string) (string, bool) {
	if len(s) < len(prefix) || !strings.EqualFold(s[:len(prefix)], prefix) {
		return "", false
	}
	return s[len(prefix):], true
}

// Cmd parses a line and returns the command, argument or an error.
// Command is converted to upper case.
func Cmd(line string) (cmd string, arg string, err error) {
	line = strings.TrimRight(line, "\r\n")

	l := len(line)
	switch {
	case l == 0:
		return "", "", nil
	case l < 4:
		return "", "", fmt.Errorf("command too short: %q", line)
	case l == 4:
		return strings.ToUpper(line), "", nil
	case l == 5:
		// Too long to be only command, too short to have args
		return "", "", fmt.Errorf("mangled command: %q", line)
	// Commands are always 4 characters long, except STARTTLS which is 8 characters long.
	// STARTTLS has no parameters (RFC 3207).
	case l == 8 && strings.EqualFold(line, "STARTTLS"):
		return "STARTTLS", "", nil
	}

	// If we made it here, command is long enough to have args
	if line[4] != ' ' {
		// There wasn't a space after the command?
		return "", "", fmt.Errorf("mangled command: %q", line)
	}

	return strings.ToUpper(line[0:4]), strings.TrimSpace(line[5:]), nil
}

// Arg is a single ESMTP argument (e.g. SIZE=1024) with an uppercased key.
type Arg struct {
	Key   string
	Value string
}

// Args takes the arguments proceeding a command and parses them
// into a list of key/value pairs after uppercasing each key.  Sample arg
// string:
//
//	" BODY=8BITMIME SIZE=1024 SMTPUTF8"
//
// The leading space is mandatory.
func Args(s string) ([]Arg, error) {
	if len(s) == 0 {
		return nil, nil
	}

	args := make([]Arg, 0, strings.Count(s, " "))
	// Only whitespace is an allowed separator, FieldsSec acceots all Unicode whitespace
	for arg := range strings.SplitSeq(s, " ") {
		if arg == "" {
			return nil, fmt.Errorf("empty arg in arg string: %q", s)
		}
		key, value, _ := strings.Cut(arg, "=")
		if strings.IndexByte(value, '=') >= 0 {
			return nil, fmt.Errorf("failed to parse arg string: %q", arg)
		}
		args = append(args, Arg{Key: strings.ToUpper(key), Value: value})
	}
	return args, nil
}

// HelloArgument parses helo argument
func HelloArgument(arg string) (string, error) {
	domain := arg
	if idx := strings.IndexRune(arg, ' '); idx >= 0 {
		domain = arg[:idx]
	}
	if domain == "" {
		return "", errors.New("invalid domain")
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
		}
		return fmt.Errorf("expected '%v', got '%v'", string(ch), string(p.S[0]))
	}
	return nil
}

// ReversePath parses a Reverse-path (RFC 5321 section 4.1.2):
//
//	Reverse-path = Path / "<>"
func (p *Parser) ReversePath() (string, error) {
	if after, ok := strings.CutPrefix(p.S, "<>"); ok {
		if after != "" && after[0] != ' ' {
			return "", fmt.Errorf("unexpected %q after null reverse-path", after[0])
		}
		// remove the separator after the reverse-path
		p.S, _ = strings.CutPrefix(after, " ")
		return "", nil
	}
	return p.Path()
}

// Path parses a Path (RFC 5321 section 4.1.2):
//
//	Path  = "<" [ A-d-l ":" ] Mailbox ">"
//	A-d-l = At-domain *( "," At-domain )
//
// The source route (A-d-l) is parsed and discarded as required by RFC 5321
// section 4.1.1.3 and appendix C. For compatibility with broken clients the
// angle brackets may be omitted.
//
// The returned mailbox is in canonical form, see Mailbox.
func (p *Parser) Path() (string, error) {
	hasBracket := p.acceptByte('<')
	if hasBracket && p.acceptByte('@') {
		if err := p.adl(); err != nil {
			return "", fmt.Errorf("in a-d-l: %v", err)
		}
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

	if p.S != "" && p.S[0] != ' ' {
		return "", fmt.Errorf("unexpected %q after path", p.S[0])
	}

	// remove the separator after the path
	p.S, _ = strings.CutPrefix(p.S, " ")

	return mbox, nil
}

// adl parses the remainder of an A-d-l (the leading '@' has been consumed)
// including the terminating ':'.
func (p *Parser) adl() error {
	for {
		if _, err := p.domain(); err != nil {
			return err
		}
		if p.acceptByte(':') {
			return nil
		}
		if err := p.expectByte(','); err != nil {
			return err
		}
		if err := p.expectByte('@'); err != nil {
			return err
		}
	}
}

// Mailbox parses a Mailbox (RFC 5321 section 4.1.2, extended by RFC 6531
// section 3.3 for SMTPUTF8):
//
//	Mailbox    = Local-part "@" ( Domain / address-literal )
//	Local-part = Dot-string / Quoted-string
//
// The returned mailbox is in canonical form: a quoted local-part whose
// content is a valid Dot-string is returned unquoted (both forms are
// equivalent), any other quoted local-part is returned as a Quoted-string
// with only '"' and '\' escaped. A local-part starting with '"' is therefore
// always quoted, and an unquoted local-part never contains '@', '"', '<', '>'
// or whitespace.
func (p *Parser) Mailbox() (string, error) {
	localPart, err := p.localPart()
	if err != nil {
		return "", fmt.Errorf("in local-part: %v", err)
	}

	if err := p.expectByte('@'); err != nil {
		return "", err
	}

	var domain string
	if ch, _ := p.peekByte(); ch == '[' {
		domain, err = p.addressLiteral()
		if err != nil {
			return "", fmt.Errorf("in address-literal: %v", err)
		}
	} else {
		domain, err = p.domain()
		if err != nil {
			return "", fmt.Errorf("in domain: %v", err)
		}
	}

	return localPart + "@" + domain, nil
}

// localPart parses a Local-part and returns it in canonical form.
func (p *Parser) localPart() (string, error) {
	if !p.acceptByte('"') {
		s := p.dotString()
		if !isDotString(s) {
			return "", errors.New("malformed dot-string")
		}
		return s, nil
	}

	// Quoted-string   = DQUOTE *QcontentSMTP DQUOTE
	// QcontentSMTP    = qtextSMTP / quoted-pairSMTP
	// quoted-pairSMTP = %d92 %d32-126
	// qtextSMTP       = %d32-33 / %d35-91 / %d93-126 / UTF8-non-ascii
	var sb strings.Builder
	for {
		ch, ok := p.readByte()
		if !ok {
			return "", errors.New("unterminated quoted-string")
		}
		switch {
		case ch == '"':
			content := sb.String()
			if !utf8.ValidString(content) {
				return "", errors.New("invalid UTF-8 in quoted-string")
			}
			if isDotString(content) {
				return content, nil
			}
			return quoteLocalPart(content), nil
		case ch == '\\':
			ch, ok = p.readByte()
			if !ok || ch < 32 || ch > 126 {
				return "", errors.New("malformed quoted-pair")
			}
		case ch < 32 || ch == 127:
			return "", fmt.Errorf("invalid character %q in quoted-string", ch)
		default:
		}
		sb.WriteByte(ch)
	}
}

// dotString consumes the longest run of atext and '.' characters.
func (p *Parser) dotString() string {
	i := 0
	for i < len(p.S) && (p.S[i] == '.' || isAtext(p.S[i])) {
		i++
	}
	s := p.S[:i]
	p.S = p.S[i:]
	return s
}

// isDotString reports whether s is a Dot-string:
//
//	Dot-string = Atom *("."  Atom)
//	Atom       = 1*atext
func isDotString(s string) bool {
	if s == "" || !utf8.ValidString(s) {
		return false
	}
	for atom := range strings.SplitSeq(s, ".") {
		if atom == "" {
			return false
		}
		for i := 0; i < len(atom); i++ {
			if !isAtext(atom[i]) {
				return false
			}
		}
	}
	return true
}

// isAtext reports whether ch is atext (RFC 5322 section 3.2.3), extended by
// UTF8-non-ascii octets (RFC 6531 section 3.3).
func isAtext(ch byte) bool {
	switch {
	case 'a' <= ch && ch <= 'z', 'A' <= ch && ch <= 'Z', '0' <= ch && ch <= '9', ch >= 0x80:
		return true
	}
	return strings.IndexByte("!#$%&'*+-/=?^_`{|}~", ch) >= 0
}

// quoteLocalPart returns content as a Quoted-string.
func quoteLocalPart(content string) string {
	var sb strings.Builder
	sb.Grow(len(content) + 2)
	sb.WriteByte('"')
	for i := 0; i < len(content); i++ {
		if content[i] == '"' || content[i] == '\\' {
			sb.WriteByte('\\')
		}
		sb.WriteByte(content[i])
	}
	sb.WriteByte('"')
	return sb.String()
}

// domain parses a Domain (RFC 5321 section 4.1.2, extended by RFC 6531
// section 3.3 for U-labels):
//
//	Domain     = sub-domain *("." sub-domain)
//	sub-domain = Let-dig [Ldh-str] / U-label
func (p *Parser) domain() (string, error) {
	i := 0
	for i < len(p.S) && (p.S[i] == '.' || p.S[i] == '-' || isLetDig(p.S[i])) {
		i++
	}
	d := p.S[:i]
	if d == "" {
		return "", errors.New("domain is empty")
	}
	if !utf8.ValidString(d) {
		return "", errors.New("invalid UTF-8 in domain")
	}
	for label := range strings.SplitSeq(d, ".") {
		if label == "" || label[0] == '-' || label[len(label)-1] == '-' {
			return "", fmt.Errorf("malformed domain %q", d)
		}
	}
	p.S = p.S[i:]
	return d, nil
}

// isLetDig reports whether ch is Let-dig, extended by UTF8-non-ascii octets.
func isLetDig(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' || ch >= 0x80
}

// addressLiteral parses an address-literal (RFC 5321 section 4.1.3):
//
//	address-literal         = "[" ( IPv4-address-literal /
//	                          IPv6-address-literal /
//	                          General-address-literal ) "]"
//	IPv6-address-literal    = "IPv6:" IPv6-addr
//	General-address-literal = Standardized-tag ":" 1*dcontent
//	Standardized-tag        = Ldh-str
//	dcontent                = %d33-90 / %d94-126
func (p *Parser) addressLiteral() (string, error) {
	end := strings.IndexByte(p.S, ']')
	if p.S == "" || p.S[0] != '[' || end < 0 {
		return "", errors.New("malformed address-literal")
	}
	lit := p.S[1:end]
	for i := 0; i < len(lit); i++ {
		if lit[i] < 33 || lit[i] > 126 || lit[i] == '[' || lit[i] == '\\' {
			return "", fmt.Errorf("invalid character %q in address-literal", lit[i])
		}
	}

	tag, value, hasTag := strings.Cut(lit, ":")
	switch {
	case !hasTag:
		if addr, err := netip.ParseAddr(lit); err != nil || !addr.Is4() {
			return "", fmt.Errorf("malformed IPv4 address-literal %q", lit)
		}
	case strings.EqualFold(tag, "IPv6"):
		addr, err := netip.ParseAddr(value)
		if err != nil || !addr.Is6() || addr.Zone() != "" {
			return "", fmt.Errorf("malformed IPv6 address-literal %q", lit)
		}
	default:
		if !isLdhStr(tag) || value == "" {
			return "", fmt.Errorf("malformed general address-literal %q", lit)
		}
	}

	s := p.S[:end+1]
	p.S = p.S[end+1:]
	return s, nil
}

// isLdhStr reports whether s is Let-dig [Ldh-str] (ASCII only).
func isLdhStr(s string) bool {
	if s == "" || s[0] == '-' || s[len(s)-1] == '-' {
		return false
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 0x80 || ch != '-' && !isLetDig(ch) {
			return false
		}
	}
	return true
}

// IsASCII reports whether s only contains ASCII characters. Addresses
// containing non-ASCII characters require SMTPUTF8 (RFC 6531).
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// IsPrintableASCII checks if string contains only printable ascii.
func IsPrintableASCII(val string) bool {
	for _, ch := range val {
		if ch < ' ' || '~' < ch {
			return false
		}
	}
	return true
}

// CheckNotifySet checks if a DSNNotify array isn't malformed.
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
