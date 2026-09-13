package server_test

import (
	"io"
	"strings"
	"testing"

	"github.com/uponusolutions/go-smtp/server"
)

// The server always propagates that it is supporting pipelining.
// If the client uses pipelining, MAIL FROM and RCPT TO and BDAT can happen in the same group.
// Which means when the server receives a BDAT command
//  - before MAIL FROM (or after a failed MAIL FROM - pipelining)
//  - before any RCPT TO (or after only failed RCPT TO - pipelining)
// The server needs to read and discard the chunk.

// TestServerBdatDiscard checks if a BDAT chunk is discarded if it is issued directly after connection.
func TestServerBdatDiscard(t *testing.T) {
	_, s, c, scanner, caps := testServerEhlo(t, nil, server.WithEnableCHUNKING(true))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	if !caps["CHUNKING"] {
		t.Fatal("server does not advertise CHUNKING")
	}

	// MAIL FROM as chunk content
	_, _ = io.WriteString(c, "BDAT 9 LAST\r\nMAIL FROM")
	scan(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "5") {
		t.Fatal("Invalid BDAT response:", reply)
	}

	_, _ = io.WriteString(c, "NOOP\r\n")
	scan(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "250 2.0.0") {
		// If BDAT chunk isn't discarded, MAIL FROM is executed as command
		// 501 5.5.2 Was expecting MAIL arg syntax of FROM:<address>
		t.Fatalf("Invalid NOOP response: %v", reply)
	}
}

// TestServerBdatDiscard checks if a BDAT chunk is discarded if it is issued without recipients.
func TestServerBdatNoRecipientsDiscard(t *testing.T) {
	_, s, c, scanner, caps := testServerEhlo(t, nil, server.WithEnableCHUNKING(true))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	if !caps["CHUNKING"] {
		t.Fatal("server does not advertise CHUNKING")
	}

	_, _ = io.WriteString(c, "MAIL FROM:<alice@wonderland.book>\r\n")
	scan(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "250 2.0.0") {
		t.Fatalf("Invalid MAIL FROM response: %v", reply)
	}

	// RCPT TO as chunk content
	_, _ = io.WriteString(c, "BDAT 7 LAST\r\nRCPT TO")
	scan(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "5") {
		t.Fatal("Invalid BDAT response:", reply)
	}

	_, _ = io.WriteString(c, "NOOP\r\n")
	scan(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "250 2.0.0") {
		// If BDAT chunk isn't discarded, MAIL FROM is executed as command
		// 501 5.5.2 Was expecting MAIL arg syntax of FROM:<address>
		t.Fatalf("Invalid NOOP response: %v", reply)
	}
}
