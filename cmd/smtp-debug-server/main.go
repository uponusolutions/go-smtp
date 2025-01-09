package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/server"
)

var addr = "127.0.0.1:1025"

func init() {
	flag.StringVar(&addr, "l", addr, "Listen address")
}

type backend struct{}

func (bkd *backend) NewSession(c *server.Conn) (server.Session, error) {
	return &session{}, nil
}

type session struct{}

func (s *session) AuthPlain(username, password string) error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	return nil
}

func (s *session) Data(r func() io.Reader) error {
	r()
	return nil
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer(&backend{})

	s.Addr = addr
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	s.Debug = os.Stdout

	log.Println("Starting SMTP server at", addr)
	log.Fatal(s.ListenAndServe())
}
