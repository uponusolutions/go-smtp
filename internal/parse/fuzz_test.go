package parse

import (
	"strings"
	"testing"
)

// FuzzCmd parses arbitrary command lines.
func FuzzCmd(f *testing.F) {
	f.Add("MAIL FROM:<root@nsa.gov>")
	f.Add("starttls")
	f.Add("NOOP")
	f.Add("x")

	f.Fuzz(func(t *testing.T, line string) {
		cmd, _, err := Cmd(line)
		if err != nil {
			return
		}
		if cmd != strings.ToUpper(cmd) {
			t.Fatalf("Cmd(%q) returned non uppercase command %q", line, cmd)
		}
	})
}

// FuzzArgs parses arbitrary ESMTP argument strings.
func FuzzArgs(f *testing.F) {
	f.Add(" BODY=8BITMIME SIZE=1024 SMTPUTF8")
	f.Add("A=B=C")
	f.Add("=")

	f.Fuzz(func(t *testing.T, s string) {
		args, err := Args(s)
		if err != nil {
			return
		}
		if len(args) > len(strings.Fields(s)) {
			t.Fatalf("Args(%q) returned more args than fields", s)
		}
		for _, arg := range args {
			if arg.Key != strings.ToUpper(arg.Key) {
				t.Fatalf("Args(%q) returned non uppercase key %q", s, arg.Key)
			}
		}
	})
}

// FuzzParser parses arbitrary path and mailbox strings.
func FuzzParser(f *testing.F) {
	f.Add("<root@nsa.gov> AUTH=asdf@example.org")
	f.Add("<@a,@b:c@d>")
	f.Add("<\"quoted string\"@example.org>")
	f.Add("<>")

	f.Fuzz(func(t *testing.T, s string) {
		for _, parse := range []func(p *Parser) (string, error){
			func(p *Parser) (string, error) { return p.ReversePath() },
			func(p *Parser) (string, error) { return p.Path() },
			func(p *Parser) (string, error) { return p.Mailbox() },
		} {
			p := Parser{S: s}
			if _, err := parse(&p); err == nil {
				if !strings.HasSuffix(s, p.S) {
					t.Fatalf("parser did not consume a prefix of %q, rest %q", s, p.S)
				}
			}
		}
	})
}
