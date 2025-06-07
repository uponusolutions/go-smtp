package tester

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"testing"

	s "net/smtp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var srv = Standard()

var addr string

func TestMain(m *testing.M) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	listen, err := srv.Listen()
	if err != nil {
		slog.Error("error listen server", slog.Any("error", err))
	}

	addr = listen.Addr().String()

	go func() {
		if err := srv.Serve(ctx, listen); err != nil {
			slog.Error("smtp server response %s", slog.Any("error", err))
		}
	}()

	exitVal := m.Run()

	if err := srv.Close(); err != nil {
		slog.Error("error closing server", slog.Any("error", err))
	}

	os.Exit(exitVal)
}

func TestBackend_AddLoad(t *testing.T) {
	b := Backend{}

	m := &Mail{
		From:       "alice@i.com",
		Recipients: []string{"bob@e.com"},
		Data:       []byte("test"),
	}

	b.Add(m)
	m1, found := b.Load(m.From, m.Recipients)
	assert.True(t, found)
	assert.Equal(t, m, m1)
}

func TestSendMail(t *testing.T) {
	from := "alice@i.com"
	recipients := []string{"bob@e.com"}
	data := []byte("Test mail\r\n")

	// Send without TLS.
	require.Nil(t, s.SendMail(addr, nil, from, recipients, data))

	m, found := GetBackend(srv).Load(from, recipients)
	assert.True(t, found)
	assert.Equal(t, from, m.From)
	assert.Equal(t, recipients, m.Recipients)
	assert.Equal(t, data, m.Data)
}
