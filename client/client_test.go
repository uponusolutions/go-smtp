package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/tester"
)

var s = tester.Standard()

var addr string

func TestMain(m *testing.M) {
	listen, err := s.Listen()
	if err != nil {
		slog.Error("error listen server", slog.Any("error", err))
	}

	addr = listen.Addr().String()

	go func() {
		if err := s.Serve(context.Background(), listen); err != nil {
			log.Printf("smtp server response %s", err)
		}
	}()

	defer func() {
		if err := s.Close(); err != nil {
			slog.Error("error closing server", "err", err)
		}
	}()

	// Wait a second to let the server come up.
	time.Sleep(time.Second)

	ret := m.Run()

	os.Exit(ret)
}

func TestClient_SendMail(t *testing.T) {
	c := New(WithServerAddress(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Close())

		// Calling again must be ok.
		assert.NoError(t, c.Quit())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "mal@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.SendMail(from, recipients, in)
	require.NoError(t, err)

	// Lookup email.
	m, found := tester.GetBackend(s).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailUTF8Force(t *testing.T) {
	c := New(WithServerAddress(addr), WithMailOptions(MailOptions{UTF8: UTF8Force}))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Close())

		// Calling again must be ok.
		assert.NoError(t, c.Quit())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "mal@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.SendMail(from, recipients, in)
	require.ErrorContains(t, err, "server does not support SMTPUTF8")

	// simulate from a client perspective that the server does support smtputf8
	c.ext["SMTPUTF8"] = ""

	_, _, err = c.SendMail(from, recipients, in)
	require.ErrorContains(t, err, "504: SMTPUTF8 is not implemented")
}

func TestClient_VerifyUTF8Force(t *testing.T) {
	c := New(WithServerAddress(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Close())

		// Calling again must be ok.
		assert.NoError(t, c.Quit())
	}()

	err := c.Verify("Bob@external.com", &VrfyOptions{UTF8: UTF8Force})
	require.ErrorContains(t, err, "server does not support SMTPUTF8")

	// simulate from a client perspective that the server does support smtputf8
	c.ext["SMTPUTF8"] = ""

	err = c.Verify("Bob@external.com", &VrfyOptions{UTF8: UTF8Force})
	require.ErrorContains(t, err, "504: SMTPUTF8 is not implemented")
}

func TestClient_InvalidLocalName(t *testing.T) {
	c := New(WithServerAddress(addr), WithLocalName("hostinjection>\n\rDATA\r\nInjected message body\r\n.\r\nQUIT\r\n"))
	require.NotNil(t, c)
	require.ErrorContains(t, c.Connect(context.Background()), "smtp: the local name must not contain CR or LF")
}

func TestClient_ServerAddress(t *testing.T) {
	c := New(WithServerAddress("test"))
	require.NotNil(t, c)
	require.Equal(t, "test", c.ServerAddress())
}

func TestClient_Send(t *testing.T) {
	c := New(WithServerAddress(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Close())

		// Calling again must be ok.
		assert.NoError(t, c.Quit())
	}()

	data := []byte("All your base are belong to us.")
	from := "alice1@internal.com"
	recipients := []string{"Bob1@external.com", "mal1@external.com"}

	err := c.Send(from, recipients, data)
	require.NoError(t, err)

	// Lookup email.
	m, found := tester.GetBackend(s).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

var (
	server     = "" // ends with .mail.protection.outlook.com:25
	priv       = ``
	certs      = ``
	eml        = ``
	from       = ""
	recipients = []string{}
)

func TestClient_SendMicrosoft(t *testing.T) {
	t.Skip()
	cert, err := tls.X509KeyPair([]byte(certs), []byte(priv))
	require.NoError(t, err)

	c := New(WithServerAddress(server), WithTLSConfig(&tls.Config{
		Certificates: []tls.Certificate{cert},
	}), WithSecurity(SecurityTLS))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Close())
		assert.NoError(t, c.Quit())
	}()

	err = c.Send(from, recipients, []byte(eml))
	require.NoError(t, err)
}
