// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
	"github.com/uponusolutions/go-smtp/tester"

	"github.com/uponusolutions/go-sasl"
)

// Don't send a trailing space on AUTH command when there's no initial response:
// https://github.com/golang/go/issues/17794
func TestClientAuthTrimSpace(t *testing.T) {
	server := "220 hello world\r\n" +
		"200 some more"
	wrote := &bytes.Buffer{}

	fake := tester.NewFakeConn(server, wrote)

	c := New()
	c.setConn(fake)
	require.Error(t, c.Auth(toServerNoRespAuth{}))
	require.NoError(t, c.Close())
	if got, want := wrote.String(), "AUTH FOOAUTH\r\n*\r\n"; got != want {
		t.Errorf("wrote %q; want %q", got, want)
	}
}

// toServerNoRespAuth is an implementation of Auth that only implements
// the Start method, and returns "FOOAUTH", nil, nil. Notably, it returns
// nil for "toServer" so we can test that we don't send spaces at the end of
// the line. See TestClientAuthTrimSpace.
type toServerNoRespAuth struct{}

func (toServerNoRespAuth) Start() (proto string, toServer []byte, err error) {
	return "FOOAUTH", nil, nil
}

func (toServerNoRespAuth) Next(_ []byte) (toServer []byte, err error) {
	panic("unexpected call")
}

func TestBasic(t *testing.T) {
	server := strings.Join(strings.Split(basicServer, "\n"), "\r\n")
	client := strings.Join(strings.Split(basicClient, "\n"), "\r\n")

	cmdbuf := &bytes.Buffer{}
	fake := tester.NewFakeConn(server, cmdbuf)

	c := &Client{text: textsmtp.NewTextproto(fake, 4096, 4096, 0), conn: fake, localName: "localhost"}

	if err := c.helo(); err != nil {
		t.Fatalf("HELO failed: %s", err)
	}
	if err := c.ehlo(); err == nil {
		t.Fatalf("Expected first EHLO to fail")
	}
	if err := c.ehlo(); err != nil {
		t.Fatalf("Second EHLO failed: %s", err)
	}

	if ok, args := c.Extension("aUtH"); !ok || args != "LOGIN PLAIN" {
		t.Fatalf("Expected AUTH supported")
	}
	if ok, _ := c.Extension("DSN"); ok {
		t.Fatalf("Shouldn't support DSN")
	}
	if !c.SupportsAuth("PLAIN") {
		t.Errorf("Expected AUTH PLAIN supported")
	}
	if size, ok := c.MaxMessageSize(); !ok {
		t.Errorf("Expected SIZE supported")
	} else if size != 35651584 {
		t.Errorf("Expected SIZE=35651584, got %v", size)
	}

	if err := c.Mail("user@gmail.com", nil); err == nil {
		t.Fatalf("MAIL should require authentication")
	}

	if err := c.Verify("user1@gmail.com", nil); err == nil {
		t.Fatalf("First VRFY: expected no verification")
	}
	if err := c.Verify("user2@gmail.com>\r\nDATA\r\nAnother injected message body\r\n.\r\nQUIT\r\n", nil); err == nil {
		t.Fatalf("VRFY should have failed due to a message injection attempt")
	}
	if err := c.Verify("user2@gmail.com", nil); err != nil {
		t.Fatalf("Second VRFY: expected verification, got %s", err)
	}

	c.serverName = "smtp.google.com"
	if err := c.Auth(sasl.NewPlainClient("", "user", "pass")); err != nil {
		t.Fatalf("AUTH failed: %s", err)
	}

	if err := c.Rcpt("golang-nuts@googlegroups.com>\r\nDATA\r\nInjected message body\r\n.\r\nQUIT\r\n", nil); err == nil {
		t.Fatalf("RCPT should have failed due to a message injection attempt")
	}
	if err := c.Mail("user@gmail.com>\r\nDATA\r\nAnother injected message body\r\n.\r\nQUIT\r\n", nil); err == nil {
		t.Fatalf("MAIL should have failed due to a message injection attempt")
	}
	if err := c.Mail("user@gmail.com", nil); err != nil {
		t.Fatalf("MAIL failed: %s", err)
	}
	if err := c.Rcpt("golang-nuts@googlegroups.com", nil); err != nil {
		t.Fatalf("RCPT failed: %s", err)
	}
	msg := `From: user@gmail.com
To: golang-nuts@googlegroups.com
Subject: Hooray for Go

Line 1
.Leading dot line .
Goodbye.`
	w, err := c.Data()
	if err != nil {
		t.Fatalf("DATA failed: %s", err)
	}
	if _, err := w.Write([]byte(msg)); err != nil {
		t.Fatalf("Data write failed: %s", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("Bad data response: %s", err)
	}

	if err := c.Quit(); err != nil {
		t.Fatalf("QUIT failed: %s", err)
	}

	actualcmds := cmdbuf.String()
	if client != actualcmds {
		t.Fatalf("Got:\n%s\nExpected:\n%s", actualcmds, client)
	}
}

func TestBasic_smtp(t *testing.T) {
	faultyServer := `220 mx.google.com at your service
250-mx.google.com at your service
250 ENHANCEDSTATUSCODES
500 5.0.0 Failing with enhanced code
500 Failing without enhanced code
500-5.0.0 Failing with multiline and enhanced code
500 5.0.0 ... still failing
`
	// RFC 2034 says that enhanced codes *SHOULD* be included in errors,
	// this means it can be violated hence we need to handle last
	// case properly.

	faultyServer = strings.Join(strings.Split(faultyServer, "\n"), "\r\n")

	wrote := &bytes.Buffer{}
	fake := tester.NewFakeConn(faultyServer, wrote)

	c := New()
	c.setConn(fake)

	require.NoError(t, c.greet())
	require.NoError(t, c.hello())

	err := c.Mail("whatever", nil)
	if err == nil {
		t.Fatal("MAIL succeeded")
	}
	smtpErr, ok := err.(*smtp.Status)
	if !ok {
		t.Fatal("Returned error is not smtp")
	}
	if smtpErr.Code != 500 {
		t.Fatalf("Wrong status code, got %d, want %d", smtpErr.Code, 500)
	}
	if smtpErr.EnhancedCode != (smtp.EnhancedCode{5, 0, 0}) {
		t.Fatalf("Wrong enhanced code, got %v, want %v", smtpErr.EnhancedCode, smtp.EnhancedCode{5, 0, 0})
	}
	if smtpErr.Message != "Failing with enhanced code" {
		t.Fatalf("Wrong message, got %s, want %s", smtpErr.Message, "Failing with enhanced code")
	}

	err = c.Mail("whatever", nil)
	if err == nil {
		t.Fatal("MAIL succeeded")
	}
	smtpErr, ok = err.(*smtp.Status)
	if !ok {
		t.Fatal("Returned error is not smtp")
	}
	if smtpErr.Code != 500 {
		t.Fatalf("Wrong status code, got %d, want %d", smtpErr.Code, 500)
	}
	if smtpErr.Message != "Failing without enhanced code" {
		t.Fatalf("Wrong message, got %s, want %s", smtpErr.Message, "Failing without enhanced code")
	}

	err = c.Mail("whatever", nil)
	if err == nil {
		t.Fatal("MAIL succeeded")
	}
	smtpErr, ok = err.(*smtp.Status)
	if !ok {
		t.Fatal("Returned error is not smtp")
	}
	if smtpErr.Code != 500 {
		t.Fatalf("Wrong status code, got %d, want %d", smtpErr.Code, 500)
	}
	if want := "Failing with multiline and enhanced code\n... still failing"; smtpErr.Message != want {
		t.Fatalf("Wrong message, got %s, want %s", smtpErr.Message, want)
	}
}

func TestClient_TooLongLine(t *testing.T) {
	faultyServer := []string{
		"220 mx.google.com at your service\r\n",
		"250 2.0.0 Kk\r\n",
		"500 5.0.0 nU6XC5JJUfiuIkC7NhrxZz36Rl/rXpkfx9QdeZJ+rno6W5J9k9HvniyWXBBi1gOZ/CUXEI6K7Uony70eiVGGGkdFhP1rEvMGny1dqIRo3NM2NifrvvLIKGeX6HrYmkc7NMn9BwHyAnt5oLe5eNVDI+grwIikVPNVFZi0Dg4Xatdg5Cs8rH1x9BWhqyDoxosJst4wRoX4AymYygUcftM3y16nVg/qcb1GJwxSNbah7VjOiSrk6MlTdGR/2AwIIcSw7pZVJjGbCorniOTvKBcyut1YdbrX/4a/dBhvLfZtdSccqyMZAdZno+tGrnu+N2ghFvz6cx6bBab9Z4JJQMlkK/g1y7xjEPr6nKwruAf71NzOclPK5wzs2hY3Ku9xEjU0Cd+g/OjAzVsmeJk2U0q+vmACZsFAiOlRynXKFPLqMAg8skM5lioRTm05K/u3aBaUq0RKloeBHZ/zNp/kfHNp6TmJKAzvsXD3Xdo+PRAgCZRTRAl3ydGdrOOjxTULCVlgOL6xSAJdj9zGkzQoEW4tRmp1OiIab4GSxCtkIo7XnAowJ7EPUfDGTV3hhl5Qn7jvZjPCPlruRTtzVTho7D3HBEouWv1qDsqdED23myw0Ma9ZlobSf9eHqsSv1MxjKG2D5DdFBACu6pXGz3ceGreOHYWnI74TkoHtQ5oNuF6VUkGjGN+f4fOaiypQ54GJ8skTNoSCHLK4XF8ZutSxWzMR+LKoJBWMb6bdAiFNt+vXZOUiTgmTqs6Sw79JXqDX9YFxryJMKjHMiFkm+RZbaK5sIOXqyq+RNmOJ+G0unrQHQMCES476c7uvOlYrNoJtq+uox1qFdisIE/8vfSoKBlTtw+r2m87djIQh4ip/hVmalvtiF5fnVTxigbtwLWv8rAOCXKoktU0c2ie0a5hGtvZT0SXxwX8K2CeYXb81AFD2IaLt/p8Q4WuZ82eOCeXP72qP9yWYj6mIZdgyimm8wjrDowt2yPJU28ZD6k3Ei6C31OKgMpCf8+MW504/VCwld7czAIwjJiZe3DxtUdfM7Q565OzLiWQgI8fxjsvlCKMiOY7q42IGGsVxXJAFMtDKdchgqQA1PJR1vrw+SbI3Mh4AGnn8vKn+WTsieB3qkloo7MZlpMz/bwPXg7XadOVkUaVeHrZ5OsqDWhsWOLtPZLi5XdNazPzn9uxWbpelXEBKAjZzfoawSUgGT5vCYACNfz/yIw1DB067N+HN1KvVddI6TNBA32lpqkQ6VwdWztq6pREE51sNl9p7MUzr+ef0331N5DqQsy+epmRDwebosCx15l/rpvBc91OnxmMMXDNtmxSzVxaZjyGDmJ7RDdTy/Su76AlaMP1zxivxg2MU/9zyTzM16coIAMOd/6Uo9ezKgbZEPeMROKTzAld9BhK9BBPWofoQ0mBkVc7btnahQe3u8HoD6SKCkr9xcTcC9ZKpLkc4svrmxT9e0858pjhis9BbWD/owa6552n2+KwUMRyB8ys7rPL86hh9lBTS+05cVL+BmJfNHOA6ZizdGc3lpwIVbFmzMR5BM0HRf3OCntkWojgsdsP8BGZWHiCGGqA7YGa5AOleR887r8Zhyp47DT3Cn3Rg/icYurIx7Yh0p696gxfANo4jEkE2BOroIscDnhauwck5CCJMcabpTrGwzK8NJ+xZnCUplXnZiIaj85Uh9+yI670B4bybWlZoVmALUxxuQ8bSMAp7CAzMcMWbYJHwBqLF8V2qMj3/g81S3KOptn8b7Idh7IMzAkV8VxE3qAguzwS0zEu8l894sOFUPiJq2/llFeiHNOcEQUGJ+8ATJSAFOMDXAeQS2FoIDOYdesO6yacL0zUkvDydWbA84VXHW8DvdHPli/8hmc++dn5CXSDeBJfC/yypvrpLgkSilZMuHEYHEYHEYEHYEHEYEHEYEHEYEYEYEYEYEYEYEYEYEYEYEYEYEYEYEYEYEYEYYEYEYEYEYEYEYEYYEYEYEYEYEYEYEYEY\r\n",
		"250 2.0.0 Kk\r\n",
	}

	// The pipe is used to avoid bufio.Reader reading the too long line ahead
	// of time (in NewClient) and failing eariler than we expect.
	pr, pw := io.Pipe()

	go func() {
		for _, l := range faultyServer {
			_, err := pw.Write([]byte(l))
			require.NoError(t, err)
		}
		require.NoError(t, pw.Close())
	}()

	wrote := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(pr, wrote)

	c := New()
	c.setConn(fake)

	require.NoError(t, c.greet())
	require.NoError(t, c.hello())

	err := c.Mail("whatever", nil)
	if err != textsmtp.ErrTooLongLine {
		t.Fatal("MAIL succeeded or returned a different error:", err)
	}

	// ErrTooLongLine is "sticky" since the connection is in broken state and
	// the only reasonable way to recover is to close it.
	err = c.Mail("whatever", nil)
	if err != textsmtp.ErrTooLongLine {
		t.Fatal("Second MAIL succeeded or returned a different error:", err)
	}
}

var basicServer = `250 mx.google.com at your service
502 Unrecognized command.
250-mx.google.com at your service
250-SIZE 35651584
250-AUTH LOGIN PLAIN
250 8BITMIME
530 Authentication required
252 Send some mail, I'll try my best
250 User is valid
235 Accepted
250 Sender OK
250 Receiver OK
354 Go ahead
250 Data OK
221 OK
`

var basicClient = `HELO localhost
EHLO localhost
EHLO localhost
MAIL FROM:<user@gmail.com> BODY=8BITMIME
VRFY user1@gmail.com
VRFY user2@gmail.com
AUTH PLAIN AHVzZXIAcGFzcw==
MAIL FROM:<user@gmail.com> BODY=8BITMIME
RCPT TO:<golang-nuts@googlegroups.com>
DATA
From: user@gmail.com
To: golang-nuts@googlegroups.com
Subject: Hooray for Go

Line 1
..Leading dot line .
Goodbye.
.
QUIT
`

func TestNewClient(t *testing.T) {
	server := strings.Join(strings.Split(newClientServer, "\n"), "\r\n")
	client := strings.Join(strings.Split(newClientClient, "\n"), "\r\n")

	cmdbuf := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), cmdbuf)

	c := New()
	c.setConn(fake)
	defer func() { _ = c.Close() }()

	require.NoError(t, c.greet())
	require.NoError(t, c.hello())

	if ok, args := c.Extension("aUtH"); !ok || args != "LOGIN PLAIN" {
		t.Fatalf("Expected AUTH supported")
	}
	if ok, _ := c.Extension("DSN"); ok {
		t.Fatalf("Shouldn't support DSN")
	}
	if err := c.Quit(); err != nil {
		t.Fatalf("QUIT failed: %s", err)
	}

	actualcmds := cmdbuf.String()
	if client != actualcmds {
		t.Fatalf("Got:\n%s\nExpected:\n%s", actualcmds, client)
	}
}

var newClientServer = `220 hello world
250-mx.google.com at your service
250-SIZE 35651584
250-AUTH LOGIN PLAIN
250 8BITMIME
221 OK
`

var newClientClient = `EHLO localhost
QUIT
`

func TestNewClient2(t *testing.T) {
	server := strings.Join(strings.Split(newClient2Server, "\n"), "\r\n")
	client := strings.Join(strings.Split(newClient2Client, "\n"), "\r\n")

	cmdbuf := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), cmdbuf)

	c := New()
	c.setConn(fake)
	defer func() { _ = c.Close() }()

	err := c.greet()
	require.NoError(t, err)

	err = c.hello()
	require.NoError(t, err)

	if ok, _ := c.Extension("DSN"); ok {
		t.Fatalf("Shouldn't support DSN")
	}
	if err := c.Quit(); err != nil {
		t.Fatalf("QUIT failed: %s", err)
	}

	actualcmds := cmdbuf.String()
	if client != actualcmds {
		t.Fatalf("Got:\n%s\nExpected:\n%s", actualcmds, client)
	}
}

var newClient2Server = `220 hello world
502 EH?
250-mx.google.com at your service
250-SIZE 35651584
250-AUTH LOGIN PLAIN
250 8BITMIME
221 OK
`

var newClient2Client = `EHLO localhost
HELO localhost
QUIT
`

func TestHello(t *testing.T) {
	if len(helloServer) != len(helloClient) {
		t.Fatalf("Hello server and client size mismatch")
	}

	for i := range helloServer {
		HelloCase(t, i)
	}
}

func HelloCase(t *testing.T, i int) {
	server := strings.Join(strings.Split(baseHelloServer+helloServer[i], "\n"), "\r\n")
	client := strings.Join(strings.Split(baseHelloClient+helloClient[i], "\n"), "\r\n")

	cmdbuf := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), cmdbuf)

	c := New()
	c.setConn(fake)
	defer func() { _ = c.Close() }()

	c.serverName = "fake.host"
	c.localName = "customhost"

	require.NoError(t, c.greet())
	if i > 0 {
		require.NoError(t, c.hello())
	}

	var err error
	switch i {
	case 0:
		c.localName = "customhost"
		err = c.hello()
	case 1:
		err = c.startTLS()
		if err.Error() == "SMTP error 502: Not implemented" {
			err = nil
		}
	case 2:
		err = c.Verify("test@example.com", nil)
	case 3:
		c.serverName = "smtp.google.com"
		err = c.Auth(sasl.NewPlainClient("", "user", "pass"))
	case 4:
		err = c.Mail("test@example.com", nil)
	case 5:
		ok, _ := c.Extension("feature")
		if ok {
			t.Errorf("Expected FEATURE not to be supported")
		}
	case 6:
		err = c.Reset()
	case 7:
		err = c.Quit()
	case 8:
		err = c.Verify("test@example.com", nil)
		if err != nil {
			c.localName = "customhost"
			err = c.hello()
			if err != nil {
				t.Errorf("Want error, got none")
			}
		}
	case 9:
		err = c.Noop()
	default:
		t.Fatalf("Unhandled command")
	}

	if err != nil {
		t.Errorf("Command %d failed: %v", i, err)
	}

	actualcmds := cmdbuf.String()
	if client != actualcmds {
		t.Errorf("Got:\n%s\nExpected:\n%s", actualcmds, client)
	}
}

var baseHelloServer = `220 hello world
502 EH?
250-mx.google.com at your service
250 FEATURE
`

var helloServer = []string{
	"",
	"502 Not implemented\n",
	"250 User is valid\n",
	"235 Accepted\n",
	"250 Sender ok\n",
	"",
	"250 Reset ok\n",
	"221 Goodbye\n",
	"250 Sender ok\n",
	"250 ok\n",
}

var baseHelloClient = `EHLO customhost
HELO customhost
`

var helloClient = []string{
	"",
	"STARTTLS\nQUIT\n",
	"VRFY test@example.com\n",
	"AUTH PLAIN AHVzZXIAcGFzcw==\n",
	"MAIL FROM:<test@example.com>\n",
	"",
	"RSET\n",
	"QUIT\n",
	"VRFY test@example.com\n",
	"NOOP\n",
}

var shuttingDownServerHello = `220 hello world
421 Service not available, closing transmission channel
`

func TestHello_421Response(t *testing.T) {
	server := strings.Join(strings.Split(shuttingDownServerHello, "\n"), "\r\n")
	client := "EHLO customhost\r\n"

	cmdbuf := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), cmdbuf)

	c := New()
	c.setConn(fake)
	defer func() { _ = c.Close() }()

	require.NoError(t, c.greet())

	c.serverName = "fake.host"
	c.localName = "customhost"

	err := c.hello()
	if err == nil {
		t.Errorf("Expected Hello to fail")
	}

	var smtp *smtp.Status
	if !errors.As(err, &smtp) || smtp.Code != 421 || smtp.Message != "Service not available, closing transmission channel" {
		t.Errorf("Expected error 421, got %v", err)
	}

	actualcmds := cmdbuf.String()
	if client != actualcmds {
		t.Errorf("Got:\n%s\nExpected:\n%s", actualcmds, client)
	}
}

/* unused

var sendMailServer = `220 hello world
502 EH?
250 mx.google.com at your service
250 Sender ok
250 Receiver ok
354 Go ahead
250 Data ok
221 Goodbye
`

var sendMailClient = `EHLO localhost
HELO localhost
MAIL FROM:<test@example.com>
RCPT TO:<other@example.com>
DATA
From: test@example.com
To: other@example.com
Subject: SendMail test

SendMail is working for me.
.
QUIT
`
*/

func TestAuthFailed(t *testing.T) {
	server := strings.Join(strings.Split(authFailedServer, "\n"), "\r\n")
	client := strings.Join(strings.Split(authFailedClient, "\n"), "\r\n")

	cmdbuf := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), cmdbuf)

	c := New()
	c.setConn(fake)
	defer func() { _ = c.Close() }()

	require.NoError(t, c.greet())
	require.NoError(t, c.hello())

	c.serverName = "smtp.google.com"
	err := c.Auth(sasl.NewPlainClient("", "user", "pass"))

	if err == nil {
		t.Error("Auth: expected error; got none")
	} else if err.Error() != "SMTP error 535: Invalid credentials\nplease see www.example.com" {
		t.Errorf("Auth: got error: %v, want: %s", err, "Invalid credentials\nplease see www.example.com")
	}

	actualcmds := cmdbuf.String()
	if client != actualcmds {
		t.Errorf("Got:\n%s\nExpected:\n%s", actualcmds, client)
	}
}

var authFailedServer = `220 hello world
250-mx.google.com at your service
250 AUTH LOGIN PLAIN
535-Invalid credentials
535 please see www.example.com
221 Goodbye
`

var authFailedClient = `EHLO localhost
AUTH PLAIN AHVzZXIAcGFzcw==
*
`

func TestTLSConnState(t *testing.T) {
	ln := newLocalListener(t)
	defer func() { require.NoError(t, ln.Close()) }()
	clientDone := make(chan bool)
	serverDone := make(chan bool)
	go func() {
		defer close(serverDone)
		c, err := ln.Accept()
		if err != nil {
			t.Errorf("Server accept: %v", err)
			return
		}
		defer func() { _ = c.Close() }()
		if err := serverHandle(c, t); err != nil {
			t.Errorf("server error: %v", err)
		}
	}()
	go func() {
		defer close(clientDone)
		cfg := &tls.Config{ServerName: "example.com", RootCAs: testRootCAs}

		c := New(
			WithServerAddress(ln.Addr().String()),
			WithTLSConfig(cfg),
			WithSecurity(SecurityStartTLS),
		)
		err := c.Connect(context.Background())
		if err != nil {
			t.Errorf("Client dial: %v", err)
			return
		}

		defer func() { require.NoError(t, c.Quit()) }()
		cs, ok := c.TLSConnectionState()
		if !ok {
			t.Errorf("TLSConnectionState returned ok == false; want true")
			return
		}
		if cs.Version == 0 || !cs.HandshakeComplete {
			t.Errorf("ConnectionState = %#v; expect non-zero Version and HandshakeComplete", cs)
		}
	}()
	<-clientDone
	<-serverDone
}

func newLocalListener(t *testing.T) net.Listener {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		ln, err = net.Listen("tcp6", "[::1]:0")
	}
	if err != nil {
		t.Fatal(err)
	}
	return ln
}

type smtpSender struct {
	w io.Writer
}

func (s smtpSender) send(t *testing.T, f string) {
	_, err := s.w.Write([]byte(f + "\r\n"))
	require.NoError(t, err)
}

// smtp server, finely tailored to deal with our own client only!
func serverHandle(c net.Conn, t *testing.T) error {
	send := func(d string) { smtpSender{c}.send(t, d) }
	send("220 127.0.0.1 ESMTP service ready")
	s := bufio.NewScanner(c)
	for s.Scan() {
		switch s.Text() {
		case "EHLO localhost":
			send("250-127.0.0.1 ESMTP offers a warm hug of welcome")
			send("250-STARTTLS")
			send("250 Ok")
		case "STARTTLS":
			send("220 Go ahead")
			keypair, err := tls.X509KeyPair(localhostCert, localhostKey)
			if err != nil {
				return err
			}
			config := &tls.Config{Certificates: []tls.Certificate{keypair}}
			c = tls.Server(c, config)
			err = serverHandleTLS(c, t)
			_ = c.Close()
			return err
		default:
			t.Fatalf("unrecognized command: %q", s.Text())
		}
	}
	return s.Err()
}

func serverHandleTLS(c net.Conn, t *testing.T) error {
	send := func(d string) { smtpSender{c}.send(t, d) }
	s := bufio.NewScanner(c)
	for s.Scan() {
		switch s.Text() {
		case "EHLO localhost":
			send("250 Ok")
		case "MAIL FROM:<joe1@example.com>":
			send("250 Ok")
		case "RCPT TO:<joe2@example.com>":
			send("250 Ok")
		case "DATA":
			send("354 send the mail data, end with .")
			send("250 Ok")
		case "Subject: test":
		case "":
		case "howdy!":
		case ".":
		case "QUIT":
			send("221 127.0.0.1 Service closing transmission channel")
			return nil
		default:
			t.Fatalf("unrecognized command during TLS: %q", s.Text())
		}
	}
	return s.Err()
}

var testRootCAs *x509.CertPool

func init() {
	testRootCAs = x509.NewCertPool()
	testRootCAs.AppendCertsFromPEM(localhostCert)
}

// localhostCert is a PEM-encoded TLS cert generated from src/crypto/tls:
//
//	go run generate_cert.go --rsa-bits 1024 --host 127.0.0.1,::1,example.com \
//			--ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`
-----BEGIN CERTIFICATE-----
MIICFDCCAX2gAwIBAgIRAK0xjnaPuNDSreeXb+z+0u4wDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCBnzANBgkqhkiG9w0BAQEFAAOBjQAw
gYkCgYEA0nFbQQuOWsjbGtejcpWz153OlziZM4bVjJ9jYruNw5n2Ry6uYQAffhqa
JOInCmmcVe2siJglsyH9aRh6vKiobBbIUXXUU1ABd56ebAzlt0LobLlx7pZEMy30
LqIi9E6zmL3YvdGzpYlkFRnRrqwEtWYbGBf3znO250S56CCWH2UCAwEAAaNoMGYw
DgYDVR0PAQH/BAQDAgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1UdEwEB/wQF
MAMBAf8wLgYDVR0RBCcwJYILZXhhbXBsZS5jb22HBH8AAAGHEAAAAAAAAAAAAAAA
AAAAAAEwDQYJKoZIhvcNAQELBQADgYEAbZtDS2dVuBYvb+MnolWnCNqvw1w5Gtgi
NmvQQPOMgM3m+oQSCPRTNGSg25e1Qbo7bgQDv8ZTnq8FgOJ/rbkyERw2JckkHpD4
n4qcK27WkEDBtQFlPihIM8hLIuzWoi/9wygiElTy/tVL3y7fGCvY2/k1KBthtZGF
tN8URjVmyEo=
-----END CERTIFICATE-----`)

// localhostKey is the private key for localhostCert.
var localhostKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDScVtBC45ayNsa16NylbPXnc6XOJkzhtWMn2Niu43DmfZHLq5h
AB9+Gpok4icKaZxV7ayImCWzIf1pGHq8qKhsFshRddRTUAF3np5sDOW3QuhsuXHu
lkQzLfQuoiL0TrOYvdi90bOliWQVGdGurAS1ZhsYF/fOc7bnRLnoIJYfZQIDAQAB
AoGBAMst7OgpKyFV6c3JwyI/jWqxDySL3caU+RuTTBaodKAUx2ZEmNJIlx9eudLA
kucHvoxsM/eRxlxkhdFxdBcwU6J+zqooTnhu/FE3jhrT1lPrbhfGhyKnUrB0KKMM
VY3IQZyiehpxaeXAwoAou6TbWoTpl9t8ImAqAMY8hlULCUqlAkEA+9+Ry5FSYK/m
542LujIcCaIGoG1/Te6Sxr3hsPagKC2rH20rDLqXwEedSFOpSS0vpzlPAzy/6Rbb
PHTJUhNdwwJBANXkA+TkMdbJI5do9/mn//U0LfrCR9NkcoYohxfKz8JuhgRQxzF2
6jpo3q7CdTuuRixLWVfeJzcrAyNrVcBq87cCQFkTCtOMNC7fZnCTPUv+9q1tcJyB
vNjJu3yvoEZeIeuzouX9TJE21/33FaeDdsXbRhQEj23cqR38qFHsF1qAYNMCQQDP
QXLEiJoClkR2orAmqjPLVhR3t2oB3INcnEjLNSq8LHyQEfXyaFfu4U9l5+fRPL2i
jiC0k/9L5dHUsF0XZothAkEA23ddgRs+Id/HxtojqqUT27B8MT/IGNrYsp4DvS/c
qgkeluku4GjxRlDMBuXk94xOBEinUs+p/hwP1Alll80Tpg==
-----END RSA PRIVATE KEY-----`)

var xtextClient = `MAIL FROM:<e=mc2@example.com> AUTH=e+3Dmc2@example.com
RCPT TO:<e=mc2@example.com> ORCPT=UTF-8;e\x{3D}mc2@example.com
`

func TestClientXtext(t *testing.T) {
	server := "220 hello world\r\n" +
		"250 ok\r\n" +
		"250 ok"
	client := strings.Join(strings.Split(xtextClient, "\n"), "\r\n")

	wrote := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), wrote)

	c := New()
	c.setConn(fake)

	c.ext = map[string]string{"AUTH": "PLAIN", "DSN": ""}
	email := "e=mc2@example.com"
	require.Error(t, c.Mail(email, &MailOptions{Auth: &email}))
	require.NoError(t, c.Rcpt(email, &smtp.RcptOptions{
		OriginalRecipientType: smtp.DSNAddressTypeUTF8,
		OriginalRecipient:     email,
	}))
	require.NoError(t, c.Close())
	if got := wrote.String(); got != client {
		t.Errorf("wrote %q; want %q", got, client)
	}
}

const (
	dsnEnvelopeID  = "e=mc2"
	dsnEmailRFC822 = "e=mc2@example.com"
	dsnEmailUTF8   = "e=mc2@ドメイン名例.jp"
)

var dsnServer = `220 hello world
250 ok
250 ok
250 ok
250 ok
`

var dsnClient = `MAIL FROM:<e=mc2@example.com> RET=HDRS ENVID=e+3Dmc2
RCPT TO:<e=mc2@example.com> NOTIFY=NEVER ORCPT=RFC822;e+3Dmc2@example.com
RCPT TO:<e=mc2@example.com> NOTIFY=FAILURE,DELAY ORCPT=UTF-8;e\x{3D}mc2@\x{30C9}\x{30E1}\x{30A4}\x{30F3}\x{540D}\x{4F8B}.jp
RCPT TO:<e=mc2@ドメイン名例.jp> ORCPT=UTF-8;e\x{3D}mc2@ドメイン名例.jp
`

func TestClientDSN(t *testing.T) {
	server := strings.Join(strings.Split(dsnServer, "\n"), "\r\n")
	client := strings.Join(strings.Split(dsnClient, "\n"), "\r\n")

	wrote := &bytes.Buffer{}
	fake := tester.NewFakeConnStream(strings.NewReader(server), wrote)

	c := New()
	c.setConn(fake)

	c.ext = map[string]string{"DSN": ""}
	require.Error(t, c.Mail(dsnEmailRFC822, &MailOptions{
		Return:     smtp.DSNReturnHeaders,
		EnvelopeID: dsnEnvelopeID,
	}))
	require.NoError(t, c.Rcpt(dsnEmailRFC822, &smtp.RcptOptions{
		OriginalRecipientType: smtp.DSNAddressTypeRFC822,
		OriginalRecipient:     dsnEmailRFC822,
		Notify:                []smtp.DSNNotify{smtp.DSNNotifyNever},
	}))
	require.NoError(t, c.Rcpt(dsnEmailRFC822, &smtp.RcptOptions{
		OriginalRecipientType: smtp.DSNAddressTypeUTF8,
		OriginalRecipient:     dsnEmailUTF8,
		Notify:                []smtp.DSNNotify{smtp.DSNNotifyFailure, smtp.DSNNotifyDelayed},
	}))
	c.ext["SMTPUTF8"] = ""
	require.NoError(t, c.Rcpt(dsnEmailUTF8, &smtp.RcptOptions{
		OriginalRecipientType: smtp.DSNAddressTypeUTF8,
		OriginalRecipient:     dsnEmailUTF8,
	}))
	require.NoError(t, c.Close())
	if actualcmds := wrote.String(); client != actualcmds {
		t.Errorf("wrote %q; want %q", actualcmds, client)
	}
}
