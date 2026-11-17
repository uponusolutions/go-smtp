package server_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/server"
)

func TestUTF8Clean(t *testing.T) {
	r := func(x rune) string { return string(x) }

	tests := []struct {
		name string
		addr string
		ok   bool
	}{
		{"ascii", "alice@example.org", true},
		{"empty", "", true},
		{"umlauts", "jösé@bücher.example", true},
		{"cjk", "用户@例子.广告", true},
		{"zero width joiner", "a" + r(0x200D) + "b@example.org", true},
		{"zero width non-joiner", "a" + r(0x200C) + "b@example.org", true},
		{"left-to-right mark", "a" + r(0x200E) + "b@example.org", true},
		{"right-to-left mark", "a" + r(0x200F) + "b@example.org", true},
		{"arabic letter mark", "a" + r(0x061C) + "b@example.org", true},
		{"no-break space", "a" + r(0x00A0) + "b@example.org", true},

		{"left-to-right embedding", "a" + r(0x202A) + "b@example.org", false},
		{"right-to-left override", "a" + r(0x202E) + "b@example.org", false},
		{"left-to-right isolate", "a" + r(0x2066) + "b@example.org", false},
		{"pop directional isolate", "a" + r(0x2069) + "b@example.org", false},
		{"next line", "a" + r(0x0085) + "b@example.org", false},
		{"line separator", "a" + r(0x2028) + "b@example.org", false},
		{"paragraph separator", "a" + r(0x2029) + "b@example.org", false},
		{"zero width space", "a" + r(0x200B) + "b@example.org", false},
		{"word joiner", "a" + r(0x2060) + "b@example.org", false},
		{"byte order mark", r(0xFEFF) + "a@example.org", false},
		{"in domain", "a@exa" + r(0x200B) + "mple.org", false},
		{"in quoted local-part", `"a` + r(0x202E) + `b"@example.org`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := server.UTF8Clean(tt.addr)
			if tt.ok {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}

	// the error names the character
	require.ErrorContains(t, server.UTF8Clean("a"+r(0x202E)+"b@example.org"), "U+202E")
}

// Without an address validator everything the parser accepts is passed on.
func TestServer_DefaultAddressValidator(t *testing.T) {
	be, s, c, scanner := testServerAuthenticated(t, nil, server.WithEnableSMTPUTF8(true))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	cmd := func(line, want string) {
		t.Helper()
		_, _ = io.WriteString(c, line+"\r\n")
		scanOk(t, c, scanner)
		require.True(t, strings.HasPrefix(scanner.Text(), want), "%q: got %q, want prefix %q", line, scanner.Text(), want)
	}

	rlo := string(rune(0x202E))
	cmd("MAIL FROM:<a"+rlo+"b@example.org> SMTPUTF8", "250 ")
	cmd("RCPT TO:<c"+rlo+"d@example.org>", "250 ")
	cmd("VRFY <e"+rlo+"f@example.org> SMTPUTF8", "252 ")
	cmd("DATA", "354 ")
	cmd("Hey\r\n.", "250 ")
	require.Len(t, be.messages, 1)
	require.Equal(t, "a"+rlo+"b@example.org", be.messages[0].From)
}

func TestServer_AddressValidator(t *testing.T) {
	var seen []string
	validator := func(addr string) error {
		seen = append(seen, addr)
		switch {
		case strings.HasPrefix(addr, "bad"):
			return errors.New("bad address")
		case strings.HasPrefix(addr, "custom"):
			return smtp.NewStatusS(550, smtp.EnhancedCode{5, 7, 1}, "Not for you")
		case strings.HasPrefix(addr, "positive"):
			return smtp.NewStatusS(250, smtp.EnhancedCode{2, 0, 0}, "Fine")
		}
		return nil
	}

	be, s, c, scanner := testServerAuthenticated(t, nil, server.WithAddressValidator(validator))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	cmd := func(line, want string) {
		t.Helper()
		_, _ = io.WriteString(c, line+"\r\n")
		scanOk(t, c, scanner)
		require.True(t, strings.HasPrefix(scanner.Text(), want), "%q: got %q, want prefix %q", line, scanner.Text(), want)
	}

	// MAIL: plain errors, custom statuses and positive statuses all reject.
	cmd("MAIL FROM:<bad@example.org>", "553 5.1.7 ")
	cmd("MAIL FROM:<custom@example.org>", "550 5.7.1 ")
	cmd("MAIL FROM:<positive@example.org>", "553 5.1.7 ")

	// The null reverse-path is no address.
	seen = nil
	cmd("MAIL FROM:<>", "250 ")
	require.Empty(t, seen)
	cmd("RSET", "250 ")

	// RCPT, a rejected recipient is not counted.
	cmd("MAIL FROM:<alice@example.org>", "250 ")
	cmd("RCPT TO:<bad@example.org>", "553 5.1.3 ")
	cmd("RCPT TO:<custom@example.org>", "550 5.7.1 ")
	cmd("DATA", "502 5.5.1 ")
	cmd("RCPT TO:<bob@example.org>", "250 ")

	// VRFY
	cmd("VRFY <bad@example.org>", "553 5.1.3 ")
	cmd("VRFY <custom@example.org>", "550 5.7.1 ")
	cmd("VRFY <bob@example.org>", "252 ")

	// The validator gets the canonical form of the address.
	seen = nil
	cmd("RSET", "250 ")
	cmd(`MAIL FROM:<"carol"@example.org>`, "250 ")
	cmd(`RCPT TO:<"a b"@example.org>`, "250 ")
	require.Equal(t, []string{"carol@example.org", `"a b"@example.org`}, seen)

	cmd("DATA", "354 ")
	cmd("Hey\r\n.", "250 ")
	require.Len(t, be.messages, 1)
	require.Equal(t, []string{`"a b"@example.org`}, be.messages[0].To)
}

func TestServer_AddressValidatorUTF8Clean(t *testing.T) {
	_, s, c, scanner := testServerAuthenticated(t, nil,
		server.WithEnableSMTPUTF8(true),
		server.WithAddressValidator(server.UTF8Clean),
	)
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	cmd := func(line, want string) {
		t.Helper()
		_, _ = io.WriteString(c, line+"\r\n")
		scanOk(t, c, scanner)
		require.True(t, strings.HasPrefix(scanner.Text(), want), "%q: got %q, want prefix %q", line, scanner.Text(), want)
	}

	rlo := string(rune(0x202E))
	cmd("MAIL FROM:<a"+rlo+"b@example.org> SMTPUTF8", "553 5.1.7 ")
	cmd("MAIL FROM:<jösé@example.org> SMTPUTF8", "250 ")
	cmd("RCPT TO:<c"+rlo+"d@example.org>", "553 5.1.3 ")
	cmd("RCPT TO:<bob@bücher.example>", "250 ")
}
