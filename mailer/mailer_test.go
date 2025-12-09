package mailer

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
	"github.com/uponusolutions/go-smtp/client"
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

func TestClient_ChunkingErrors(t *testing.T) {
	c := New(WithServerAddresses(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())

		// Calling again must be ok.
		assert.NoError(t, c.Disconnect())
	}()

	// server doesn't support chunking
	_, err := c.client.Bdat(0)
	require.ErrorContains(t, err, "doesn't support chunking")

	assert.NoError(t, c.Disconnect())

	c = New(WithServerAddresses(addr), WithBasic(client.WithChunkingMaxSize(-1)))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))

	// client chunking is disabled
	_, err = c.client.Bdat(0)
	require.ErrorContains(t, err, "chunking is disabled")

	assert.NoError(t, c.Disconnect())
}

func TestClient_SendMail(t *testing.T) {
	c := New(WithServerAddresses(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())

		// Calling again must be ok.
		assert.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "mal@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.Send(context.Background(), from, recipients, in)
	require.NoError(t, err)

	// Lookup email.
	m, found := tester.GetBackend(s).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMail_MultipleAddresses(t *testing.T) {
	c := New(WithServerAddresses(addr, "0.0.0.0")) // second is invalid
	require.NotNil(t, c)

	require.Equal(t, "", c.ServerAddress())
	require.NoError(t, c.Connect(context.Background()))
	require.Equal(t, addr, c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())
	require.NoError(t, c.Terminate())
	require.Equal(t, addr, c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())

	c = New(WithServerAddresses("0.0.0.0", addr)) // second is invalid
	require.NotNil(t, c)

	require.Equal(t, "", c.ServerAddress())
	require.NoError(t, c.Connect(context.Background()))
	require.Equal(t, addr, c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())
	require.NoError(t, c.Terminate())
	require.Equal(t, addr, c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())
}

func TestClient_SendMailUTF8Force(t *testing.T) {
	c := New(WithServerAddresses(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())

		// Calling again must be ok.
		assert.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "mal@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.SendAdvanced(context.Background(), from, &client.MailOptions{UTF8: client.UTF8Force}, recipients, in)
	require.ErrorContains(t, err, "server does not support SMTPUTF8")
}

func TestClient_VerifyUTF8Force(t *testing.T) {
	c := New(WithServerAddresses(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())

		// Calling again must be ok.
		assert.NoError(t, c.Disconnect())
	}()

	err := c.Verify("Bob@external.com", &client.VrfyOptions{UTF8: client.UTF8Force})
	require.ErrorContains(t, err, "server does not support SMTPUTF8")
}

func TestClient_InvalidLocalName(t *testing.T) {
	c := New(WithServerAddresses(addr), WithBasic(
		client.WithLocalName("hostinjection>\n\rDATA\r\nInjected message body\r\n.\r\nQUIT\r\n")),
	)
	require.NotNil(t, c)
	require.ErrorContains(t, c.Connect(context.Background()), "smtp: the local name must not contain CR or LF")
}

func TestClient_Send(t *testing.T) {
	c := New(WithServerAddresses(addr))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())

		// Calling again must be ok.
		assert.NoError(t, c.Disconnect())
	}()

	data := []byte("All your base are belong to us.")
	from := "alice1@internal.com"
	recipients := []string{"Bob1@external.com", "mal1@external.com"}

	_, _, err := c.Send(context.Background(), from, recipients, bytes.NewBuffer(data))
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

	c := New(WithServerAddresses(server), WithTLSConfig(&tls.Config{
		Certificates: []tls.Certificate{cert},
	}), WithSecurity(SecurityTLS))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())
		assert.NoError(t, c.Disconnect())
	}()

	_, _, err = c.Send(context.Background(), from, recipients, bytes.NewBuffer([]byte(eml)))
	require.NoError(t, err)
}
