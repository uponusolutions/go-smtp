package mailer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-sasl"
	"github.com/uponusolutions/go-smtp/tester"
)

// fakeStartTLSServer advertises STARTTLS and AUTH and answers STARTTLS with h.
func fakeStartTLSServer(h tester.FakeHandler) *tester.FakeServer {
	return &tester.FakeServer{
		Extensions: []string{"STARTTLS", "AUTH PLAIN"},
		Handlers:   map[string]tester.FakeHandler{"STARTTLS": h},
	}
}

func TestSecurity_Default(t *testing.T) {
	require.Equal(t, SecurityStartTLS, DefaultConfig().extra.security)
}

// H1: with mandatory STARTTLS (the default), a rejected STARTTLS must not leave a plaintext
// connection behind that a later Send silently reuses.
func TestSecurity_StartTLSRejectedNoPlaintextReuse(t *testing.T) {
	f := fakeStartTLSServer(func(s *tester.FakeSession, _ string) bool {
		s.Reply("454 TLS not available")
		return true
	})
	c := New(WithServerAddresses(f.Listen(t)))

	for range 2 {
		msg := strings.NewReader("Subject: x\r\n\r\nsecret\r\n")
		_, _, err := c.Send(t.Context(), "from@example.com", []string{"to@example.com"}, msg)
		require.Error(t, err)
		assert.False(t, c.Connected(), "plaintext connection must be closed")
	}
	assert.False(t, f.Received("MAIL"), "envelope was sent in cleartext")
	assert.False(t, f.Received("DATA"), "body was sent in cleartext")
}

// H2: with the default (prefer STARTTLS) a failed TLS handshake falls back to
// plain and still authenticates. This is intended and documented; mandatory
// STARTTLS must never send the credentials.
func TestSecurity_BrokenHandshakeAuth(t *testing.T) {
	for security, wantAuth := range map[Security]bool{
		SecurityPreferStartTLS: true,
		SecurityStartTLS:       false,
	} {
		f := fakeStartTLSServer(func(s *tester.FakeSession, _ string) bool {
			s.Reply("220 ready")
			return false // break the handshake
		})
		c := New(
			WithServerAddresses(f.Listen(t)),
			WithSecurity(security),
			WithSASLClient(sasl.NewPlainClient("", "user", "password")),
		)

		_ = c.Connect(t.Context())
		_ = c.Terminate()

		assert.Equal(t, wantAuth, f.Received("AUTH"), "security %d", security)
	}
}
