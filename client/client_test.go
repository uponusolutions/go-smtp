package client

import (
	"bytes"
	"crypto/tls"
	"log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/davrux/go-smtptester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var s = smtptester.Standard()

func TestMain(m *testing.M) {
	go func() {
		if err := s.ListenAndServe(); err != nil {
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
	c := NewClient(WithServerAddress("127.0.0.1:2525"))
	require.NotNil(t, c)

	require.NoError(t, c.Connect())
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
	m, found := smtptester.GetBackend(s).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_InvalidLocalName(t *testing.T) {
	c := NewClient(WithServerAddress("127.0.0.1:2525"), WithLocalName("hostinjection>\n\rDATA\r\nInjected message body\r\n.\r\nQUIT\r\n"))
	require.NotNil(t, c)
	require.ErrorContains(t, c.Connect(), "smtp: the local name must not contain CR or LF")
}

func TestClient_Send(t *testing.T) {
	c := NewClient(WithServerAddress("127.0.0.1:2525"))
	require.NotNil(t, c)

	require.NoError(t, c.Connect())
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
	m, found := smtptester.GetBackend(s).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

var server = "" // ends with .mail.protection.outlook.com:25
var priv = ``
var certs = ``
var eml = ``
var from = ""
var recipients = []string{}

func TestClient_SendMicrosoft(t *testing.T) {
	t.Skip()
	cert, err := tls.X509KeyPair([]byte(certs), []byte(priv))
	require.NoError(t, err)

	c := NewClient(WithServerAddress(server), WithTLSConfig(&tls.Config{
		Certificates: []tls.Certificate{cert},
	}), WithSecurity(Security_TLS))
	require.NotNil(t, c)

	require.NoError(t, c.Connect())
	defer func() {
		assert.NoError(t, c.Close())
		assert.NoError(t, c.Quit())
	}()

	err = c.Send(from, recipients, []byte(eml))
	require.NoError(t, err)
}
