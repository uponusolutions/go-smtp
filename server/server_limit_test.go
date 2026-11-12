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
