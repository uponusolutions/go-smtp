package mailer

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"log/slog"
	"strings"
	"sync"
	"testing"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/server"
	"github.com/uponusolutions/go-smtp/tester/testserver"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/client"
)

// backend rejects every recipient whose local part starts with "notfound",
// naming the address in the reply so tests can tell the rejections apart.
var backend = testserver.Backend{
	Mails: sync.Map{},
	Rcpt: func(_ context.Context, to string, _ *smtp.RcptOptions) error {
		if strings.HasPrefix(to, "notfound") {
			return smtp.NewStatusS(550, smtp.NoEnhancedCode, "not found "+to)
		}
		return nil
	},
}

var s = []*server.Server{
	testserver.Standard(
		server.WithBackend(&backend),
		server.WithMaxRecipients(1000),
	),
	testserver.Standard(
		server.WithBackend(&backend),
		server.WithWriterSize(1),
		server.WithReaderSize(1),
		server.WithMaxRecipients(1000),
	),
}

var addr = []string{
	"",
	"",
}

func TestMain(m *testing.M) {
	for i, sc := range s {
		listen, err := sc.Listen()
		if err != nil {
			slog.Error("error listen server", slog.Any("error", err))
		}
		addr[i] = listen.Addr().String()
		go func() {
			if err := sc.Serve(context.Background(), listen); err != nil {
				log.Printf("smtp server response %s", err)
			}
		}()
		// nolint:revive
		defer func() {
			if err := sc.Close(); err != nil {
				slog.Error("error closing server", "err", err)
			}
		}()
	}

	m.Run()
}

func TestClient_DisconnectTwicePipeline(t *testing.T) {
	c := New(WithServerAddresses(addr[0]), WithBasic(client.WithPipelining(true)))
	require.NoError(t, c.Connect(t.Context()))
	require.NoError(t, c.Terminate())
	require.NoError(t, c.Disconnect())
}

func TestClient_ChunkingErrors(t *testing.T) {
	c := New(WithServerAddresses(addr[0]))
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

	c = New(WithServerAddresses(addr[0]), WithBasic(client.WithChunkingMaxSize(-1)))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))

	// client chunking is disabled
	_, err = c.client.Bdat(0)
	require.ErrorContains(t, err, "chunking is disabled")

	assert.NoError(t, c.Disconnect())
}

func TestClient_SendMailAutoconnect(t *testing.T) {
	c := New(WithServerAddresses(addr[0]))
	require.NotNil(t, c)

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
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMail(t *testing.T) {
	for i, sa := range addr {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			basic := WithBasic()
			if i == 1 {
				basic = WithBasic(client.WithReaderSize(1), client.WithWriterSize(1), client.WithPipelining(true))
			}

			c := New(WithServerAddresses(sa), basic)
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
			m, found := testserver.GetBackend(s[0]).Load(from, recipients)
			assert.True(t, found)

			t.Logf("Found %t, mail %+v\n", found, m)
		})
	}
}

func TestClient_SendMailTooManyRcptsAbortPipelining(t *testing.T) {
	c := New(WithServerAddresses(addr[0]),
		WithAbortOnRcptReject(true),
		WithBasic(client.WithPipelining(true)))
	require.NotNil(t, c)

	require.NoError(t, c.Connect(context.Background()))
	defer func() {
		assert.NoError(t, c.Terminate())

		// Calling again must be ok.
		assert.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := makeEmails(1001)

	in := bytes.NewBuffer(data)

	status, failures, err := c.Send(context.Background(), from, recipients, in)
	require.Equal(t,
		smtp.NewStatusS(452, smtp.EnhancedCode{4, 5, 3}, "Maximum limit of 1000 recipients reached"),
		err,
	)
	require.Nil(t, status)
	require.Equal(t, 0, len(failures))

	_, failures, err = c.Send(context.Background(), from, recipients[0:1000], in)
	require.NoError(t, err)
	require.Equal(t, 0, len(failures))
}

func TestClient_SendMailDirect(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "mal@external.com"}

	_, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
	)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailDirectPipelining(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "mal@external.com"}

	_, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
		WithBasic(client.WithPipelining(true)),
	)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func makeEmails(x int) []string {
	emails := make([]string, x)
	for i := range emails {
		emails[i] = fmt.Sprintf("%d@external.com", i)
	}
	return emails
}

func TestClient_SendMailDirectManyRcptsPipelining(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := makeEmails(1000)

	_, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
		WithBasic(client.WithPipelining(true)),
	)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	require.True(t, found)
	require.Equal(t, recipients, m.Recipients)
}

func TestClient_SendMailDirectTooManyRcptsPipelining(t *testing.T) {
	for i, sa := range addr {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			data := []byte("Hello World!")
			from := "alice@internal.com"
			recipients := makeEmails(1001)

			basic := WithBasic(client.WithPipelining(true))
			if i == 1 {
				basic = WithBasic(client.WithReaderSize(1), client.WithWriterSize(1), client.WithPipelining(true))
			}

			res, err := Send(
				context.Background(),
				from,
				recipients,
				func() io.Reader { return bytes.NewReader(data) },
				WithServerAddresses(sa),
				basic,
			)
			require.NoError(t, err)
			require.Equal(t, 1, len(res.Responses))
			require.Equal(t, 1000, len(res.Responses[0].Rcpts))
			require.Equal(t, 1, len(res.Failures))
			require.Equal(t,
				smtp.NewStatusS(452, smtp.EnhancedCode{4, 5, 3}, "Maximum limit of 1000 recipients reached"),
				res.Failures[0].Error,
			)
			require.Equal(t,
				[]string{recipients[1000]},
				res.Failures[0].Rcpts,
			)
		})
	}
}

func TestClient_SendMailDirectFailPipelining(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "notfound@external.com", "mal@external.com"}

	_, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
		WithBasic(client.WithPipelining(true)),
	)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, []string{"Bob@external.com", "mal@external.com"})
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailDirectAbortOnRcptRejectPipelining(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "notfound@external.com", "mal@external.com"}

	_, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
		WithAbortOnRcptReject(true),
		WithBasic(client.WithPipelining(true)),
	)
	require.ErrorContains(t, err, "not supported")
}

func TestClient_SendMailDirectAllRejectedPipelining(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"notfound@external.com"}

	res, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
		WithBasic(client.WithPipelining(true)),
	)
	require.NoError(t, err)

	require.Equal(t, 0, len(res.Responses))
	require.Equal(t, 1, len(res.Failures))
	require.Equal(t, recipients, res.Failures[0].Rcpts)
}

var pipeliningCases = []struct {
	name       string
	pipelining bool
}{
	{name: "without pipelining"},
	{name: "with pipelining", pipelining: true},
}

// sendReport sends a mail to rcpts and returns the report, failing the test on
// a transport level error.
func sendReport(t *testing.T, rcpts []string, pipelining bool) Report {
	t.Helper()

	opts := []Option{WithServerAddresses(addr[0])}
	if pipelining {
		opts = append(opts, WithBasic(client.WithPipelining(true)))
	}

	res, err := Send(
		context.Background(),
		"alice@internal.com",
		rcpts,
		func() io.Reader { return bytes.NewReader([]byte("Hello World!")) },
		opts...,
	)
	require.NoError(t, err)

	return res
}

// TestClient_SendMailAllRejectedKeepsPerRcptFailures checks that the individual
// rejection reasons survive when every recipient is refused.
//
// A server following RFC 2920 3.2 (3) answers DATA negatively once no valid
// recipient is left, so the transaction ends in an error. That error describes
// the transaction ("no valid recipients"), not any single address, and must not
// replace the per recipient 550s that were already collected.
func TestClient_SendMailAllRejectedKeepsPerRcptFailures(t *testing.T) {
	rcpts := []string{"notfound1@external.com", "notfound2@external.com", "notfound3@external.com"}

	for _, tc := range pipeliningCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sendReport(t, rcpts, tc.pipelining)

			assert.Empty(t, res.Responses, "nothing was delivered")
			require.Len(t, res.Failures, len(rcpts), "one failure per recipient")

			byRcpt := map[string]error{}
			for _, f := range res.Failures {
				require.Len(t, f.Rcpts, 1,
					"a failure must name a single recipient, got %v", f.Rcpts)
				byRcpt[f.Rcpts[0]] = f.Error
			}

			for _, rcpt := range rcpts {
				err, ok := byRcpt[rcpt]
				require.True(t, ok, "no failure reported for %s", rcpt)

				var status *smtp.Status
				require.ErrorAs(t, err, &status)
				assert.Equal(t, 550, status.Code,
					"%s must carry its own rejection, not the DATA rejection", rcpt)
				assert.Contains(t, err.Error(), rcpt,
					"%s must carry its own reply text", rcpt)
			}
		})
	}
}

// TestClient_SendMailPartiallyRejectedKeepsPerRcptFailures is the counterpart:
// the accepted recipients still get a response and are not blamed for the
// rejection of another address.
func TestClient_SendMailPartiallyRejectedKeepsPerRcptFailures(t *testing.T) {
	accepted := []string{"Bob@external.com", "mal@external.com"}
	rejected := "notfound1@external.com"
	rcpts := []string{accepted[0], rejected, accepted[1]}

	for _, tc := range pipeliningCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sendReport(t, rcpts, tc.pipelining)

			require.Len(t, res.Failures, 1)
			assert.Equal(t, []string{rejected}, res.Failures[0].Rcpts)

			require.Len(t, res.Responses, 1)
			assert.Equal(t, accepted, res.Responses[0].Rcpts,
				"delivered recipients must exclude the rejected one")
		})
	}
}

func TestClient_SendMailAutoconnectAbortOnRcptReject(t *testing.T) {
	c := New(WithServerAddresses(addr[0]), WithAbortOnRcptReject(true))
	require.NotNil(t, c)

	defer func() {
		require.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "notfound@external.com", "mal@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.Send(context.Background(), from, recipients, in)
	require.ErrorContains(t, err, "notfound@external.com")

	recipients = []string{"Bob@external.com", "mal@external.com"}

	_, _, err = c.Send(context.Background(), from, recipients, in)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailAutoconnectAbortOnRcptRejectPipelining(t *testing.T) {
	c := New(WithServerAddresses(addr[0]), WithAbortOnRcptReject(true), WithBasic(client.WithPipelining(true)))
	require.NotNil(t, c)

	defer func() {
		require.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "notfound@external.com", "mal@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.Send(context.Background(), from, recipients, in)
	require.ErrorContains(t, err, "notfound@external.com")

	recipients = []string{"Bob@external.com", "mal@external.com"}

	_, _, err = c.Send(context.Background(), from, recipients, in)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailAutoconnectAbortOnRcptRejectAll(t *testing.T) {
	c := New(WithServerAddresses(addr[0]), WithAbortOnRcptReject(true))
	require.NotNil(t, c)

	defer func() {
		require.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"notfound@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.Send(context.Background(), from, recipients, in)
	require.ErrorContains(t, err, "notfound@external.com")

	recipients = []string{"Bob@external.com"}

	_, _, err = c.Send(context.Background(), from, recipients, in)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailAutoconnectAbortOnRcptRejectAllPipelining(t *testing.T) {
	c := New(WithServerAddresses(addr[0]), WithAbortOnRcptReject(true), WithBasic(client.WithPipelining(true)))
	require.NotNil(t, c)

	defer func() {
		require.NoError(t, c.Disconnect())
	}()

	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"notfound@external.com"}

	in := bytes.NewBuffer(data)

	_, _, err := c.Send(context.Background(), from, recipients, in)
	require.ErrorContains(t, err, "notfound@external.com")

	recipients = []string{"Bob@external.com"}

	_, _, err = c.Send(context.Background(), from, recipients, in)
	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMailDirectFail(t *testing.T) {
	data := []byte("Hello World!")
	from := "alice@internal.com"
	recipients := []string{"Bob@external.com", "notfound@external.com"}

	rec, err := Send(
		context.Background(),
		from,
		recipients,
		func() io.Reader { return bytes.NewReader(data) },
		WithServerAddresses(addr[0]),
	)
	require.Equal(t, 1, len(rec.Failures))
	require.Equal(t, []string{"notfound@external.com"}, rec.Failures[0].Rcpts)
	require.ErrorContains(t, rec.Failures[0].Error, "550")

	require.NoError(t, err)

	// Lookup email.
	m, found := testserver.GetBackend(s[0]).Load(from, []string{"Bob@external.com"})
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

func TestClient_SendMail_MultipleAddresses(t *testing.T) {
	c := New(WithServerAddresses(addr[0], "0.0.0.0")) // second is invalid
	require.NotNil(t, c)

	require.Equal(t, "", c.ServerAddress())
	require.NoError(t, c.Connect(context.Background()))
	require.Equal(t, addr[0], c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())
	require.NoError(t, c.Terminate())
	require.Equal(t, addr[0], c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())

	c = New(WithServerAddresses("0.0.0.0", addr[0])) // second is invalid
	require.NotNil(t, c)

	require.Equal(t, "", c.ServerAddress())
	require.NoError(t, c.Connect(context.Background()))
	require.Equal(t, addr[0], c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())
	require.NoError(t, c.Terminate())
	require.Equal(t, addr[0], c.ServerAddress())
	require.Equal(t, "localhost", c.ServerName())
}

func TestClient_SendMailUTF8Force(t *testing.T) {
	c := New(WithServerAddresses(addr[0]))
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

	_, _, err := c.SendAdvanced(
		context.Background(),
		from,
		&client.MailOptions{UTF8: client.UTF8Force},
		recipients,
		nil,
		in,
	)
	require.ErrorContains(t, err, "server does not support SMTPUTF8")
}

func TestClient_VerifyUTF8Force(t *testing.T) {
	c := New(WithServerAddresses(addr[0]))
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
	c := New(WithServerAddresses(addr[0]), WithBasic(
		client.WithLocalName("hostinjection>\n\rDATA\r\nInjected message body\r\n.\r\nQUIT\r\n")),
	)
	require.NotNil(t, c)
	require.ErrorContains(t, c.Connect(context.Background()), "smtp: the local name must not contain CR or LF")
}

func TestClient_Client(t *testing.T) {
	c := New(WithServerAddresses(addr[0]))
	require.NotNil(t, c)
	require.NotNil(t, c.Client())
}

func TestClient_Send(t *testing.T) {
	c := New(WithServerAddresses(addr[0]))
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
	m, found := testserver.GetBackend(s[0]).Load(from, recipients)
	assert.True(t, found)

	t.Logf("Found %t, mail %+v\n", found, m)
}

var (
	address    = "" // ends with .mail.protection.outlook.com:25
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

	c := New(WithServerAddresses(address), WithTLSConfig(&tls.Config{
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
