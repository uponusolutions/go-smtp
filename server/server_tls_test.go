package server_test

import (
	"bufio"
	"crypto/tls"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/server"
	"github.com/uponusolutions/go-smtp/tester"
)

func TestServerEnforceSecureConnection(t *testing.T) {
	_, _, c, scanner, caps := testServerEhlo(t, nil, server.WithEnforceSecureConnection(true))

	if _, ok := caps["AUTH PLAIN"]; !ok {
		t.Fatal("AUTH PLAIN capability is missing when auth is enabled")
	}

	io.WriteString(c, "AUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk\r\n")
	scanner.Scan()
	if !strings.HasPrefix(scanner.Text(), "530 ") {
		t.Fatal("Should enforce STARTTLS:", scanner.Text())
	}
}

func TestServerEnforceSecureConnectionImplicitTls(t *testing.T) {
	cert, err := tester.GenX509KeyPair("localhost")
	require.NoError(t, err)

	be := new(backend)
	be.clientTls = true

	_, _, c, scanner, caps := testServerEhlo(
		t,
		be,
		server.WithEnforceSecureConnection(true),
		server.WithImplicitTLS(true),
		server.WithTLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
		}),
	)

	if _, ok := caps["AUTH PLAIN"]; !ok {
		t.Fatal("AUTH PLAIN capability is missing when auth is enabled")
	}

	io.WriteString(c, "AUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk\r\n")
	scanner.Scan()
	if !strings.HasPrefix(scanner.Text(), "235 ") {
		t.Fatal("Should succeed:", scanner.Text())
	}
}

func TestServerEnforceSecureConnectionStartTls(t *testing.T) {
	cert, err := tester.GenX509KeyPair("localhost")
	require.NoError(t, err)

	_, _, c, scanner, caps := testServerEhlo(
		t,
		nil,
		server.WithEnforceSecureConnection(true),
		server.WithTLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
		}),
	)

	if _, ok := caps["AUTH PLAIN"]; !ok {
		t.Fatal("AUTH PLAIN capability is missing when auth is enabled")
	}

	io.WriteString(c, "AUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk\r\n")
	scanner.Scan()
	if !strings.HasPrefix(scanner.Text(), "530 ") {
		t.Fatal("STARTTLS first expected:", scanner.Text())
	}
}

func TestServerEnforceSecureConnectionStartTlsStarttls(t *testing.T) {
	cert, err := tester.GenX509KeyPair("localhost")
	require.NoError(t, err)

	_, _, c, _, _ := testServerEhlo(
		t,
		nil,
		server.WithEnforceSecureConnection(true),
		server.WithTLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
		}),
	)

	io.WriteString(c, "STARTTLS\r\n")

	buf := make([]byte, 30)
	c.Read(buf)

	if string(buf) != "220 2.0.0 Ready to start TLS\r\n" {
		t.Fatal("Ready to start expected:", string(buf))
	}

	// Upgrade to TLS
	c = tls.Client(c, &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}})

	scanner := bufio.NewScanner(c)

	io.WriteString(c, "HELO localhost\r\n")
	scanner.Scan()
	if !strings.HasPrefix(scanner.Text(), "250 ") {
		t.Fatal("hello expected expected:", scanner.Text())
	}

	io.WriteString(c, "AUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk\r\n")
	scanner.Scan()
	if !strings.HasPrefix(scanner.Text(), "235 ") {
		t.Fatal("Should succeed:", scanner.Text())
	}
}
