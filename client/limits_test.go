package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// RFC 9422 section 3 and 3.7.
func TestLimit(t *testing.T) {
	tests := []struct {
		name    string
		ext     map[string]string
		mailMax int // 0 means not conveyed
		rcptMax int
	}{
		{name: "no LIMITS", ext: map[string]string{"SIZE": "1000"}},
		{name: "both", ext: map[string]string{"LIMITS": "MAILMAX=5 RCPTMAX=20"}, mailMax: 5, rcptMax: 20},
		{name: "one", ext: map[string]string{"LIMITS": "RCPTMAX=20"}, rcptMax: 20},
		{name: "case insensitive name", ext: map[string]string{"LIMITS": "mailMax=5"}, mailMax: 5},
		{name: "maximum", ext: map[string]string{"LIMITS": "MAILMAX=999999"}, mailMax: 999999},
		{name: "unknown limits ignored", ext: map[string]string{"LIMITS": "FOO BAR=x-y RCPTMAX=20"}, rcptMax: 20},
		{name: "RCPTDOMAINMAX ignored", ext: map[string]string{"LIMITS": "RCPTDOMAINMAX=3 MAILMAX=5"}, mailMax: 5},
		{name: "first wins", ext: map[string]string{"LIMITS": "MAILMAX=5 MAILMAX=7"}, mailMax: 5},

		// An invalid value only drops that limit.
		{name: "zero", ext: map[string]string{"LIMITS": "MAILMAX=0 RCPTMAX=20"}, rcptMax: 20},
		{name: "leading zero", ext: map[string]string{"LIMITS": "MAILMAX=05 RCPTMAX=20"}, rcptMax: 20},
		{name: "seven digits", ext: map[string]string{"LIMITS": "MAILMAX=1000000 RCPTMAX=20"}, rcptMax: 20},
		{name: "not a number", ext: map[string]string{"LIMITS": "MAILMAX=five RCPTMAX=20"}, rcptMax: 20},
		{name: "signed", ext: map[string]string{"LIMITS": "MAILMAX=+5 RCPTMAX=20"}, rcptMax: 20},
		{name: "no value", ext: map[string]string{"LIMITS": "MAILMAX RCPTMAX=20"}, rcptMax: 20},

		// Invalid syntax drops all limits.
		{name: "empty", ext: map[string]string{"LIMITS": ""}},
		{name: "empty value", ext: map[string]string{"LIMITS": "MAILMAX= RCPTMAX=20"}},
		{name: "semicolon in value", ext: map[string]string{"LIMITS": "FOO=a;b RCPTMAX=20"}},
		{name: "invalid name", ext: map[string]string{"LIMITS": "FOO.BAR=1 RCPTMAX=20"}},
		{name: "double space", ext: map[string]string{"LIMITS": "MAILMAX=5  RCPTMAX=20"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.setExt(tt.ext)

			n, ok := c.MaxTransactions()
			assert.Equal(t, tt.mailMax != 0, ok, "MAILMAX conveyed")
			assert.Equal(t, tt.mailMax, n)

			n, ok = c.MaxRecipients()
			assert.Equal(t, tt.rcptMax != 0, ok, "RCPTMAX conveyed")
			assert.Equal(t, tt.rcptMax, n)
		})
	}
}

// RFC 9422 section 3.6: the most recent EHLO replaces previous limits.
func TestLimit_LatestEhloWins(t *testing.T) {
	c := New()
	c.setExt(map[string]string{"LIMITS": "MAILMAX=5"})
	c.setExt(map[string]string{"SIZE": ""})

	_, ok := c.MaxTransactions()
	assert.False(t, ok)
}
