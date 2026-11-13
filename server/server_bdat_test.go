package server_test

import (
	"io"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

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

// writeSlowly writes s in parts with pauses, taking longer in total than the read timeout.
func writeSlowly(c net.Conn, s string, parts int, pause time.Duration) {
	step := len(s) / parts
	for i := 0; i < parts-1; i++ {
		_, _ = io.WriteString(c, s[i*step:(i+1)*step])
		time.Sleep(pause)
	}
	_, _ = io.WriteString(c, s[(parts-1)*step:])
}

// BDAT uses the data timeout, a chunk taking longer than the read timeout must succeed.
func TestServerBdatReadTimeoutSlowChunk(t *testing.T) {
	be, s, c, scanner, _ := testServerEhlo(t, nil,
		server.WithEnableCHUNKING(true), server.WithReadTimeout(150*time.Millisecond))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	_, _ = io.WriteString(c, "MAIL FROM:<alice@wonderland.book>\r\n")
	scanOk(t, c, scanner)
	_, _ = io.WriteString(c, "RCPT TO:<bob@wonderland.book>\r\n")
	scanOk(t, c, scanner)

	body := "From: alice@wonderland.book\r\n\r\nslow but steady\r\n"
	_, _ = io.WriteString(c, "BDAT "+strconv.Itoa(len(body))+" LAST\r\n")
	writeSlowly(c, body, 6, 50*time.Millisecond)

	scanOk(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "250 ") {
		t.Fatal("Invalid BDAT response:", reply)
	}
	if len(be.anonmsgs) != 1 || be.anonmsgs[0].Data == nil {
		t.Fatal("Invalid number of sent messages:", be.anonmsgs)
	}
}

// Discarding a BDAT chunk uses the data timeout, it may take longer than the read timeout.
func TestServerBdatReadTimeoutSlowDiscard(t *testing.T) {
	_, s, c, scanner, _ := testServerEhlo(t, nil,
		server.WithEnableCHUNKING(true), server.WithReadTimeout(150*time.Millisecond))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	_, _ = io.WriteString(c, "MAIL FROM:<alice@wonderland.book>\r\n")
	scanOk(t, c, scanner)

	// no recipients, the chunk gets discarded
	body := strings.Repeat("x", 60)
	_, _ = io.WriteString(c, "BDAT "+strconv.Itoa(len(body))+" LAST\r\n")
	writeSlowly(c, body, 6, 50*time.Millisecond)

	scanOk(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "5") {
		t.Fatal("Invalid BDAT response:", reply)
	}

	_, _ = io.WriteString(c, "NOOP\r\n")
	scanOk(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "250 ") {
		t.Fatal("Invalid NOOP response:", reply)
	}
}

// A client going silent in the middle of a BDAT chunk is disconnected after the data timeout.
func TestServerBdatDataTimeoutStalled(t *testing.T) {
	_, s, c, scanner, _ := testServerEhlo(t, nil,
		server.WithEnableCHUNKING(true), server.WithDataTimeout(100*time.Millisecond))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	_, _ = io.WriteString(c, "MAIL FROM:<alice@wonderland.book>\r\n")
	scanOk(t, c, scanner)
	_, _ = io.WriteString(c, "RCPT TO:<bob@wonderland.book>\r\n")
	scanOk(t, c, scanner)

	_, _ = io.WriteString(c, "BDAT 100 LAST\r\nonly a part")

	scanOk(t, c, scanner)
	if reply := scanner.Text(); !strings.HasPrefix(reply, "421 4.4.2 ") {
		t.Fatal("Invalid response for stalled BDAT:", reply)
	}
}

// The data timeout covers all chunks of a BDAT transfer, including the command lines in between.
func TestServerBdatDataTimeoutWholeTransfer(t *testing.T) {
	_, s, c, scanner, _ := testServerEhlo(t, nil,
		server.WithEnableCHUNKING(true), server.WithDataTimeout(150*time.Millisecond))
	defer func() {
		_ = s.Close()
		_ = c.Close()
	}()

	_, _ = io.WriteString(c, "MAIL FROM:<alice@wonderland.book>\r\n")
	scanOk(t, c, scanner)
	_, _ = io.WriteString(c, "RCPT TO:<bob@wonderland.book>\r\n")
	scanOk(t, c, scanner)

	// every chunk is fast, but the transfer as a whole is too slow
	for i := range 10 {
		_, _ = io.WriteString(c, "BDAT 3\r\nabc")
		scanOk(t, c, scanner)
		reply := scanner.Text()
		if strings.HasPrefix(reply, "421 4.4.2 ") {
			if i == 0 {
				t.Fatal("First chunk should be accepted")
			}
			return
		}
		if !strings.HasPrefix(reply, "250 ") {
			t.Fatal("Invalid BDAT response:", reply)
		}
		time.Sleep(50 * time.Millisecond)
	}
	t.Fatal("Expected the data timeout to end the BDAT transfer")
}
