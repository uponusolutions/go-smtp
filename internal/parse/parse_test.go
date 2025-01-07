package parse

import (
	"testing"
)

func TestParser(t *testing.T) {
	validReversePaths := []struct {
		raw, path, after string
	}{
		{"<>", "", ""},
		{"<root@nsa.gov>", "root@nsa.gov", ""},
		{"root@nsa.gov", "root@nsa.gov", ""},
		{"<root@nsa.gov> AUTH=asdf@example.org", "root@nsa.gov", " AUTH=asdf@example.org"},
		{"root@nsa.gov AUTH=asdf@example.org", "root@nsa.gov", " AUTH=asdf@example.org"},
	}
	for _, tc := range validReversePaths {
		p := Parser{tc.raw}
		path, err := p.ReversePath()
		if err != nil {
			t.Errorf("parser.parseReversePath(%q) = %v", tc.raw, err)
		} else if path != tc.path {
			t.Errorf("parser.parseReversePath(%q) = %q, want %q", tc.raw, path, tc.path)
		} else if p.S != tc.after {
			t.Errorf("parser.parseReversePath(%q): got after = %q, want %q", tc.raw, p.S, tc.after)
		}
	}

	invalidReversePaths := []string{
		"",
		" ",
		"asdf",
		"<Foo Bar <root@nsa.gov>>",
		" BODY=8BITMIME SIZE=12345",
		"a:b:c@example.org",
		"<root@nsa.gov",
	}
	for _, tc := range invalidReversePaths {
		p := Parser{tc}
		if path, err := p.ReversePath(); err == nil {
			t.Errorf("parser.parseReversePath(%q) = %q, want error", tc, path)
		}
	}
}
