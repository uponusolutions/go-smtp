package mailer

import (
	"bytes"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/client"
	"github.com/uponusolutions/go-smtp/tester"
)

// newStrictServer returns a tester.FakeServer that, like Postfix, refuses a MAIL
// inside an open transaction (503) and only ends a transaction on RSET or the
// end of mail data.
//
// Recipients starting with "notfound" are rejected. A sender starting with
// "nodata" makes DATA fail before 354, one starting with "reject" gets the
// message rejected after the end of data.
func newStrictServer(t *testing.T, pipelining bool) (*tester.FakeServer, string) {
	t.Helper()

	// open transaction of a connection: *tester.FakeSession -> *strictTx
	var txs sync.Map

	f := &tester.FakeServer{Handlers: map[string]tester.FakeHandler{
		"MAIL": func(fs *tester.FakeSession, line string) bool {
			if _, open := txs.LoadOrStore(fs, &strictTx{from: line}); open {
				fs.Reply("503 5.5.1 Error: nested MAIL command")
			} else {
				fs.Reply("250 2.1.0 Ok")
			}
			return true
		},
		"RCPT": func(fs *tester.FakeSession, line string) bool {
			if strings.Contains(line, "<notfound") {
				fs.Reply("550 5.1.1 not found")
				return true
			}
			v, _ := txs.Load(fs)
			if tx, open := v.(*strictTx); open {
				tx.rcpts++
			}
			fs.Reply("250 2.1.5 Ok")
			return true
		},
		"DATA": func(fs *tester.FakeSession, _ string) bool {
			v, _ := txs.Load(fs)
			tx, open := v.(*strictTx)
			if !open || tx.rcpts == 0 {
				fs.Reply("554 5.5.1 Error: no valid recipients")
				return true
			}
			if strings.Contains(tx.from, "<nodata") {
				fs.Reply("451 4.3.0 Error: queue file write error")
				return true
			}
			fs.Reply("354 End data with <CR><LF>.<CR><LF>")
			err := fs.ReadMessage()
			// the end of mail data ends the transaction
			txs.Delete(fs)
			if err != nil {
				return false
			}
			if strings.Contains(tx.from, "<reject") {
				fs.Reply("554 5.7.1 rejected")
			} else {
				fs.Reply("250 2.0.0 Ok: queued")
			}
			return true
		},
		"RSET": func(fs *tester.FakeSession, _ string) bool {
			txs.Delete(fs)
			fs.Reply("250 2.0.0 Ok")
			return true
		},
	}}
	if pipelining {
		f.Extensions = []string{"PIPELINING"}
	}

	return f, f.Listen(t)
}

// strictTx is an open transaction: the MAIL line and the number of accepted recipients.
type strictTx struct {
	from  string
	rcpts int
}

// connections counts the connections to f, each starts with EHLO.
func connections(f *tester.FakeServer) int {
	n := 0
	for _, c := range f.Commands() {
		if strings.HasPrefix(c, "EHLO ") {
			n++
		}
	}
	return n
}

// TestMailer_FailedTransactionDoesNotBreakConnection sends a failing mail and
// then a good one over the same connection. The failed transaction must be
// ended, otherwise the following MAIL is refused as nested.
func TestMailer_FailedTransactionDoesNotBreakConnection(t *testing.T) {
	cases := []struct {
		name    string
		from    string
		rcpts   []string
		wantErr string
	}{
		{name: "all recipients rejected", from: "alice@internal.com", rcpts: []string{"notfound@external.com"}},
		{
			name: "data rejected before 354", wantErr: "451",
			from: "nodata@internal.com", rcpts: []string{"bob@external.com"},
		},
		{
			name: "message rejected after end of data", wantErr: "554",
			from: "reject@internal.com", rcpts: []string{"bob@external.com"},
		},
	}

	for _, tc := range cases {
		for _, pc := range pipeliningCases {
			t.Run(tc.name+" "+pc.name, func(t *testing.T) {
				srv, addr := newStrictServer(t, pc.pipelining)

				m := New(
					WithServerAddresses(addr),
					WithSecurity(SecurityPlain),
					WithBasic(client.WithPipelining(pc.pipelining)),
				)
				t.Cleanup(func() { _ = m.Terminate() })

				_, failures, err := m.Send(t.Context(), tc.from, tc.rcpts, bytes.NewBufferString("Hello World!\r\n"))
				if tc.wantErr != "" {
					require.ErrorContains(t, err, tc.wantErr)
				} else if !pc.pipelining {
					// without pipelining DATA isn't sent at all
					require.NoError(t, err)
				}
				if tc.name == "all recipients rejected" {
					require.Len(t, failures, 1)
				}
				assert.True(t, m.Connected(), "a rejected transaction must not close the connection")

				status, failures, err := m.Send(t.Context(), "alice@internal.com",
					[]string{"bob@external.com"}, bytes.NewBufferString("Hello World!\r\n"))
				require.NoError(t, err)
				require.Empty(t, failures)
				require.NotNil(t, status)

				// a transaction left open would have failed the second Send with 503
				assert.Equal(t, 1, connections(srv), "connection must be reused")
			})
		}
	}
}

// statusReader fails like the data reader of a server whose size limit was hit.
type statusReader struct{ sent bool }

func (r *statusReader) Read(p []byte) (int, error) {
	if !r.sent {
		r.sent = true
		return copy(p, "Hello"), nil
	}
	return 0, smtp.ErrDataTooLarge
}

// TestMailer_SourceStatusErrorClosesConnection checks that a failing message
// source closes the connection, even with an *smtp.Status error. Part of the
// message is already on the wire, a RSET would end up inside the message.
func TestMailer_SourceStatusErrorClosesConnection(t *testing.T) {
	srv, addr := newStrictServer(t, false)

	m := New(WithServerAddresses(addr), WithSecurity(SecurityPlain))
	t.Cleanup(func() { _ = m.Terminate() })

	_, _, err := m.Send(t.Context(), "alice@internal.com", []string{"bob@external.com"}, &statusReader{})
	require.ErrorIs(t, err, smtp.ErrDataTooLarge)
	assert.False(t, m.Connected())

	_, _, err = m.Send(t.Context(), "alice@internal.com",
		[]string{"bob@external.com"}, bytes.NewBufferString("Hello World!\r\n"))
	require.NoError(t, err)
	assert.Equal(t, 2, connections(srv))
}
