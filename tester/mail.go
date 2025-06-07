package tester

import "strings"

// Mail is one mail received by SMTP server.
type Mail struct {
	From       string
	Recipients []string
	Data       []byte
}

// LookupKey call LookupKey for current mail.
func (m *Mail) LookupKey() string {
	return m.From + "+" + strings.Join(m.Recipients, "+")
}

// LookupKey returns a key of the format:
//
//	m.From+m.Recipient_1+m.Recipient_2...
func LookupKey(f string, r []string) string {
	return f + "+" + strings.Join(r, "+")
}
