package server_test

import (
	"bufio"
	"io"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/server"
)

// When the maximum message size is exceeded during DATA or BDAT the rest of the
// payload is still on the wire. The server must not try to interpret it as
// commands (desync / smuggling), so the limit breach is fatal: the server
// replies 552 and closes the connection.

// limitTestServer starts a server with the given max message size and opens a
// mail transaction (MAIL FROM + RCPT TO).
func limitTestServer(t *testing.T, be *backend, maxMessageBytes int64) (*backend, net.Conn, *bufio.Scanner) {
	t.Helper()

	if be == nil {
		be = &backend{}
	}
	be.closeErrors = make(chan error, 1)

	be, s, c, scanner := testServerAuthenticated(t, be,
		server.WithMaxMessageBytes(maxMessageBytes),
		server.WithEnableCHUNKING(true),
	)
	t.Cleanup(func() {
		_ = c.Close()
		_ = s.Close()
	})

	_, _ = io.WriteString(c, "MAIL FROM:<root@nsa.gov>\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))
	_, _ = io.WriteString(c, "RCPT TO:<root@gchq.gov.uk>\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))

	return be, c, scanner
}

// requireTooLargeAndClosed expects a 552 reply followed by the server closing
// the connection without answering anything left in the payload.
func requireTooLargeAndClosed(t *testing.T, be *backend, c net.Conn, scanner *bufio.Scanner) {
	t.Helper()

	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "552 5.3.4 "), "expected 552 reply, got %q", reply)

	if scan(t, c, scanner) {
		t.Fatalf("expected connection to be closed, got reply %q", scanner.Text())
	}
	require.NoError(t, scanner.Err(), "expected EOF, not a read timeout")

	select {
	case err := <-be.closeErrors:
		require.ErrorContains(t, err, "maximum message size exceeded")
	case <-time.After(5 * time.Second):
		t.Fatal("session was not closed")
	}
}

// requireSessionUsable checks that the connection is still in sync.
func requireSessionUsable(t *testing.T, c net.Conn, scanner *bufio.Scanner) {
	t.Helper()

	_, _ = io.WriteString(c, "NOOP\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "250 2.0.0 "), "expected NOOP reply, got %q", reply)
}

func TestServerLimit_Data_TooLargeClosesConnection(t *testing.T) {
	be, c, scanner := limitTestServer(t, nil, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))

	// Everything after the first 50 bytes must never be executed as commands.
	_, _ = io.WriteString(c, strings.Repeat("x", 60)+"\r\n"+
		"RSET\r\n"+
		"MAIL FROM:<smuggled@example.com>\r\n"+
		"NOOP\r\n"+
		".\r\n")

	requireTooLargeAndClosed(t, be, c, scanner)
	require.Empty(t, be.messages)
	require.Empty(t, be.anonmsgs)
}

func TestServerLimit_Data_BackendSwallowsError(t *testing.T) {
	be, c, scanner := limitTestServer(t, &backend{swallowDataErr: true}, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))

	_, _ = io.WriteString(c, strings.Repeat("x", 60)+"\r\nNOOP\r\n.\r\n")

	// Even if the backend ignores the read error, the message must not be accepted.
	requireTooLargeAndClosed(t, be, c, scanner)
}

func TestServerLimit_Data_BackendRejectsWithoutReading(t *testing.T) {
	// The backend rejects the message without reading it, the server discards
	// the rest itself and hits the limit while doing so.
	rejected := smtp.NewStatusS(554, smtp.EnhancedCode{5, 6, 0}, "rejected")
	be, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: rejected}, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))

	_, _ = io.WriteString(c, strings.Repeat("x", 60)+"\r\nNOOP\r\n.\r\n")

	requireTooLargeAndClosed(t, be, c, scanner)
}

func TestServerLimit_Data_BackendRejectsWithinLimit(t *testing.T) {
	// Rejecting a message within the limit keeps the connection usable.
	rejected := smtp.NewStatusS(554, smtp.EnhancedCode{5, 6, 0}, "rejected")
	_, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: rejected}, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))

	_, _ = io.WriteString(c, "Hello\r\n.\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "554 5.6.0 "), "got %q", reply)

	requireSessionUsable(t, c, scanner)
}

func TestServerLimit_Data_RejectedEndOfDataEndsTransaction(t *testing.T) {
	// A message rejected after the end of data ends the transaction, like an accepted one.
	rejected := smtp.NewStatusS(554, smtp.EnhancedCode{5, 6, 0}, "rejected")
	_, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: rejected}, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))
	_, _ = io.WriteString(c, "Hello\r\n.\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "554 5.6.0 "), "got %q", reply)

	// Pipelined RCPT and DATA must not reuse the previous envelope.
	_, _ = io.WriteString(c, "RCPT TO:<root@gchq.gov.uk>\r\nDATA\r\n")
	reply = readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "5"), "RCPT without MAIL, got %q", reply)
	reply = readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "5"), "DATA without MAIL, got %q", reply)

	// A new transaction can be started.
	_, _ = io.WriteString(c, "MAIL FROM:<root@nsa.gov>\r\n")
	reply = readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
}

func TestServerLimit_Data_RejectedCommandKeepsTransaction(t *testing.T) {
	// A DATA command rejected before 354 carried no message, the transaction stays open.
	rejected := smtp.NewStatusS(554, smtp.EnhancedCode{5, 6, 0}, "rejected")
	_, c, scanner := limitTestServer(t, &backend{dataErr: rejected}, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "554 5.6.0 "), "got %q", reply)

	_, _ = io.WriteString(c, "RCPT TO:<other@gchq.gov.uk>\r\n")
	reply = readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
}

func TestServerLimit_Bdat_BackendRejectsWithoutReading(t *testing.T) {
	rejected := smtp.NewStatusS(554, smtp.EnhancedCode{5, 6, 0}, "rejected")
	be, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: rejected}, 50)

	payload := strings.Repeat("x", 60) + "\r\nNOOP\r\n"
	_, _ = io.WriteString(c, "BDAT "+strconv.Itoa(len(payload))+" LAST\r\n"+payload)

	requireTooLargeAndClosed(t, be, c, scanner)
}

func TestServerLimit_Bdat_BackendRejectsWithinLimit(t *testing.T) {
	rejected := smtp.NewStatusS(554, smtp.EnhancedCode{5, 6, 0}, "rejected")
	_, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: rejected}, 50)

	_, _ = io.WriteString(c, "BDAT 6 LAST\r\nNOOP\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "554 5.6.0 "), "got %q", reply)

	requireSessionUsable(t, c, scanner)
}

func TestServerLimit_Data_ExactLimit(t *testing.T) {
	be, c, scanner := limitTestServer(t, nil, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))

	body := strings.Repeat("x", 48) + "\r\n"
	_, _ = io.WriteString(c, body+".\r\n")

	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
	require.Len(t, be.messages, 1)
	require.Equal(t, body, string(be.messages[0].Data))

	requireSessionUsable(t, c, scanner)
}

func TestServerLimit_Bdat_SingleChunkTooLarge(t *testing.T) {
	be, c, scanner := limitTestServer(t, nil, 50)

	// The chunk content looks like commands, it must not be executed.
	payload := strings.Repeat("x", 40) + "\r\nRSET\r\nNOOP\r\nMAIL FROM:<smuggled@example.com>\r\n"
	_, _ = io.WriteString(c, "BDAT "+strconv.Itoa(len(payload))+" LAST\r\n"+payload)

	requireTooLargeAndClosed(t, be, c, scanner)
	require.Empty(t, be.messages)
	require.Empty(t, be.anonmsgs)
}

func TestServerLimit_Bdat_SecondChunkTooLarge(t *testing.T) {
	be, c, scanner := limitTestServer(t, nil, 50)

	_, _ = io.WriteString(c, "BDAT 30\r\n"+strings.Repeat("a", 30))
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))

	payload := strings.Repeat("b", 15) + "\r\nNOOP\r\nRSET\r\n"
	_, _ = io.WriteString(c, "BDAT "+strconv.Itoa(len(payload))+" LAST\r\n"+payload)

	requireTooLargeAndClosed(t, be, c, scanner)
	require.Empty(t, be.messages)
	require.Empty(t, be.anonmsgs)
}

func TestServerLimit_Bdat_PipelinedChunksTooLarge(t *testing.T) {
	be, c, scanner := limitTestServer(t, nil, 50)

	// All chunks are pipelined, the second one breaches the limit.
	_, _ = io.WriteString(c,
		"BDAT 30\r\n"+strings.Repeat("a", 30)+
			"BDAT 30\r\nNOOP\r\nNOOP\r\nNOOP\r\nNOOP\r\nNOOP\r\n"+
			"BDAT 6 LAST\r\nNOOP\r\n")

	// reply for the first chunk
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))
	requireTooLargeAndClosed(t, be, c, scanner)
	require.Empty(t, be.messages)
}

func TestServerLimit_Bdat_BackendSwallowsError(t *testing.T) {
	be, c, scanner := limitTestServer(t, &backend{swallowDataErr: true}, 50)

	payload := strings.Repeat("x", 60) + "\r\nNOOP\r\n"
	_, _ = io.WriteString(c, "BDAT "+strconv.Itoa(len(payload))+" LAST\r\n"+payload)

	requireTooLargeAndClosed(t, be, c, scanner)
}

func TestServerLimit_Bdat_ExactLimit(t *testing.T) {
	be, c, scanner := limitTestServer(t, nil, 50)

	_, _ = io.WriteString(c, "BDAT 25\r\n"+strings.Repeat("a", 25))
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))
	_, _ = io.WriteString(c, "BDAT 25 LAST\r\n"+strings.Repeat("b", 25))
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)

	require.Len(t, be.messages, 1)
	require.Equal(t, strings.Repeat("a", 25)+strings.Repeat("b", 25), string(be.messages[0].Data))

	requireSessionUsable(t, c, scanner)
}

// A backend may reject with 552 on its own (e.g. its own size policy) without
// the server limit being hit. The stream is in sync, so the connection stays open.
func TestServerLimit_Data_BackendOwnTooLarge(t *testing.T) {
	_, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: smtp.ErrDataTooLarge}, 50)

	_, _ = io.WriteString(c, "DATA\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "354 "))

	_, _ = io.WriteString(c, "Hello\r\n.\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "552 5.3.4 "), "got %q", reply)

	requireSessionUsable(t, c, scanner)
}

func TestServerLimit_Bdat_BackendOwnTooLarge(t *testing.T) {
	_, c, scanner := limitTestServer(t, &backend{rejectAfterOpen: smtp.ErrDataTooLarge}, 50)

	_, _ = io.WriteString(c, "BDAT 6 LAST\r\nNOOP\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "552 5.3.4 "), "got %q", reply)

	requireSessionUsable(t, c, scanner)
}

// dialGreeted opens a connection and returns the first reply line.
func dialGreeted(t *testing.T, addr string) (net.Conn, string) {
	t.Helper()

	c, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	t.Cleanup(func() { _ = c.Close() })

	scanner := bufio.NewScanner(c)
	return c, readReply(t, c, scanner)
}

func TestServerLimit_Connections(t *testing.T) {
	noPerIP := server.WithMaxConnectionsPerIP(0)
	perIP := "421 4.7.0 localhost Error: too many connections from 127.0.0.1"
	total := "421 4.3.2 Too many concurrent connections, try again later"

	tests := []struct {
		name string
		opts []server.Option
		// Connections that must be accepted.
		accepted int
		// Reply to the next connection, empty if it must be accepted.
		rejected string
	}{
		{name: "per ip default", accepted: 10, rejected: perIP},
		{name: "per ip custom", opts: []server.Option{server.WithMaxConnectionsPerIP(2)}, accepted: 2, rejected: perIP},
		{name: "per ip disabled", opts: []server.Option{noPerIP}, accepted: 20},
		{name: "total default", opts: []server.Option{noPerIP}, accepted: 1000, rejected: total},
		{name: "total custom", opts: []server.Option{noPerIP, server.WithMaxConnections(2)}, accepted: 2, rejected: total},
		{
			name:     "total before per ip",
			opts:     []server.Option{server.WithMaxConnections(2), server.WithMaxConnectionsPerIP(2)},
			accepted: 2,
			rejected: total,
		},
		{name: "total disabled", opts: []server.Option{noPerIP, server.WithMaxConnections(0)}, accepted: 1001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, s, c, scanner := testServerGreeted(t, nil, tt.opts...)
			t.Cleanup(func() {
				_ = c.Close()
				_ = s.Close()
			})
			addr := c.RemoteAddr().String()

			// The first connection is already open.
			for range tt.accepted - 1 {
				_, reply := dialGreeted(t, addr)
				require.True(t, strings.HasPrefix(reply, "220 "), "got %q", reply)
			}

			if tt.rejected == "" {
				return
			}

			cc, reply := dialGreeted(t, addr)
			require.Equal(t, tt.rejected, reply)
			if scan(t, cc, bufio.NewScanner(cc)) {
				t.Fatal("expected connection to be closed")
			}

			// Closing a connection frees a slot.
			_, _ = io.WriteString(c, "QUIT\r\n")
			require.True(t, strings.HasPrefix(readReply(t, c, scanner), "221 "))
			_ = c.Close()

			require.Eventually(t, func() bool {
				cc, err := net.Dial("tcp", addr)
				if err != nil {
					return false
				}
				defer func() { _ = cc.Close() }()
				sc := bufio.NewScanner(cc)
				return scan(t, cc, sc) && strings.HasPrefix(sc.Text(), "220 ")
			}, 5*time.Second, 10*time.Millisecond)
		})
	}
}

// limitsCap returns the LIMITS capability of the EHLO response, or "" if missing.
func limitsCap(caps map[string]bool) string {
	for c := range caps {
		if strings.HasPrefix(c, "LIMITS") {
			return c
		}
	}
	return ""
}

func TestServerLimit_LimitsAdvertised(t *testing.T) {
	tests := []struct {
		name string
		opts []server.Option
		want string
	}{
		{
			name: "custom", opts: []server.Option{server.WithMaxTransactions(5), server.WithMaxRecipients(20)},
			want: "LIMITS MAILMAX=5 RCPTMAX=20",
		},
		{name: "mailmax disabled", opts: []server.Option{server.WithMaxTransactions(0)}, want: "LIMITS RCPTMAX=1000"},
		{name: "rcptmax disabled", opts: []server.Option{server.WithMaxRecipients(0)}, want: "LIMITS MAILMAX=1000"},
		{name: "both disabled", opts: []server.Option{server.WithMaxTransactions(0), server.WithMaxRecipients(0)}},
		// RFC 9422 allows at most six digits.
		{
			name: "maximum", opts: []server.Option{server.WithMaxTransactions(999999), server.WithMaxRecipients(1000000)},
			want: "LIMITS MAILMAX=999999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, s, c, _, caps := testServerEhlo(t, nil, tt.opts...)
			t.Cleanup(func() {
				_ = c.Close()
				_ = s.Close()
			})
			require.Equal(t, tt.want, limitsCap(caps))
		})
	}
}

// RFC 9422 MAILMAX: every MAIL command counts, also rejected ones and across RSET.
func TestServerLimit_MaxTransactions(t *testing.T) {
	be := &backend{closeErrors: make(chan error, 1)}
	_, s, c, scanner := testServerAuthenticated(t, be, server.WithMaxTransactions(3))
	t.Cleanup(func() {
		_ = c.Close()
		_ = s.Close()
	})

	_, _ = io.WriteString(c, "MAIL FROM:invalid\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "501 "), "got %q", reply)

	for range 2 {
		_, _ = io.WriteString(c, "MAIL FROM:<root@nsa.gov>\r\n")
		reply = readReply(t, c, scanner)
		require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
		_, _ = io.WriteString(c, "RSET\r\n")
		reply = readReply(t, c, scanner)
		require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
	}

	_, _ = io.WriteString(c, "MAIL FROM:<root@nsa.gov>\r\n")
	require.Equal(t, "421 4.7.0 Too many transactions in this session", readReply(t, c, scanner))
	if scan(t, c, scanner) {
		t.Fatalf("expected connection to be closed, got %q", scanner.Text())
	}

	select {
	case err := <-be.closeErrors:
		require.ErrorContains(t, err, "maximum transactions exceeded")
	case <-time.After(5 * time.Second):
		t.Fatal("session was not closed")
	}
}

// RFC 9422 RCPTMAX doesn't require counting rejected RCPT commands, a client
// keeping to the advertised limit works either way. Only accepted recipients count.
func TestServerLimit_MaxRecipients(t *testing.T) {
	_, c, scanner := limitTestServerRcpt(t, server.WithMaxRecipients(3))

	// limitTestServerRcpt already sent one RCPT, a rejected one doesn't count.
	_, _ = io.WriteString(c, "RCPT TO:invalid\r\n")
	reply := readReply(t, c, scanner)
	require.True(t, strings.HasPrefix(reply, "501 "), "got %q", reply)

	for _, rcpt := range []string{"a", "b"} {
		_, _ = io.WriteString(c, "RCPT TO:<"+rcpt+"@gchq.gov.uk>\r\n")
		reply = readReply(t, c, scanner)
		require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
	}

	_, _ = io.WriteString(c, "RCPT TO:<c@gchq.gov.uk>\r\n")
	require.Equal(t, "452 4.5.3 Maximum limit of 3 recipients reached", readReply(t, c, scanner))

	// A new transaction starts counting again.
	_, _ = io.WriteString(c, "RSET\r\nMAIL FROM:<root@nsa.gov>\r\nRCPT TO:<b@gchq.gov.uk>\r\n")
	for range 3 {
		reply = readReply(t, c, scanner)
		require.True(t, strings.HasPrefix(reply, "250 "), "got %q", reply)
	}
}

// limitTestServerRcpt starts a server with the given options and opens a mail
// transaction with one recipient.
func limitTestServerRcpt(t *testing.T, opts ...server.Option) (*backend, net.Conn, *bufio.Scanner) {
	t.Helper()

	be, s, c, scanner := testServerAuthenticated(t, nil, opts...)
	t.Cleanup(func() {
		_ = c.Close()
		_ = s.Close()
	})

	_, _ = io.WriteString(c, "MAIL FROM:<root@nsa.gov>\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))
	_, _ = io.WriteString(c, "RCPT TO:<root@gchq.gov.uk>\r\n")
	require.True(t, strings.HasPrefix(readReply(t, c, scanner), "250 "))

	return be, c, scanner
}
