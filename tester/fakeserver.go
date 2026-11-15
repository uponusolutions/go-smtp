package tester

import (
	"bufio"
	"crypto/tls"
	"net"
	"slices"
	"strings"
	"sync"
	"testing"
)

// FakeHandler handles a single command line. Returning false closes the connection.
type FakeHandler func(s *FakeSession, line string) bool

// FakeServer is a scripted SMTP server for client tests. It records every
// received command line (message content excluded) and answers with a default
// reply unless a handler for the verb is set: EHLO/HELO advertise Extensions,
// DATA consumes the message, QUIT closes the connection and everything else
// gets 250.
type FakeServer struct {
	Extensions []string               // advertised in the EHLO reply
	Handlers   map[string]FakeHandler // keyed by upper case verb, e.g. "STARTTLS"

	mu       sync.Mutex
	commands []string
}

// FakeSession is a single connection of a FakeServer.
type FakeSession struct {
	conn net.Conn
	r    *bufio.Reader
}

// Reply writes each line terminated by CRLF.
func (s *FakeSession) Reply(lines ...string) {
	for _, l := range lines {
		_, _ = s.conn.Write([]byte(l + "\r\n"))
	}
}

// StartTLS upgrades the connection to TLS, the handshake runs on the next read.
func (s *FakeSession) StartTLS(cfg *tls.Config) {
	s.conn = tls.Server(s.conn, cfg)
	s.r = bufio.NewReader(s.conn)
}

// Commands returns all received command lines in order, without CRLF.
func (f *FakeServer) Commands() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return slices.Clone(f.commands)
}

// Received reports whether a command with the given verb was received.
func (f *FakeServer) Received(verb string) bool {
	return slices.ContainsFunc(f.Commands(), func(c string) bool { return fakeVerb(c) == verb })
}

func fakeVerb(line string) string {
	verb, _, _ := strings.Cut(strings.ToUpper(line), " ")
	return verb
}

// Listen starts the server on a random local port and returns its address.
// The server is stopped when the test ends.
func (f *FakeServer) Listen(t testing.TB) string {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = l.Close() })

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}
			go f.serve(&FakeSession{conn: conn, r: bufio.NewReader(conn)})
		}
	}()
	return l.Addr().String()
}

func (f *FakeServer) serve(s *FakeSession) {
	defer func() { _ = s.conn.Close() }()

	s.Reply("220 fake ESMTP")
	for {
		line, err := s.r.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimRight(line, "\r\n")
		verb := fakeVerb(line)

		f.mu.Lock()
		f.commands = append(f.commands, line)
		f.mu.Unlock()

		if h, ok := f.Handlers[verb]; ok {
			if !h(s, line) {
				return
			}
			continue
		}

		switch verb {
		case "EHLO", "HELO":
			reply := []string{"fake"}
			reply = append(reply, f.Extensions...)
			for i, r := range reply {
				sep := "-"
				if i == len(reply)-1 {
					sep = " "
				}
				s.Reply("250" + sep + r)
			}
		case "DATA":
			s.Reply("354 go ahead")
			for {
				l, err := s.r.ReadString('\n')
				if err != nil || l == ".\r\n" {
					break
				}
			}
			s.Reply("250 queued")
		case "QUIT":
			s.Reply("221 bye")
			return
		default:
			s.Reply("250 ok")
		}
	}
}
