package benchmark

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/client"
	"github.com/uponusolutions/go-smtp/server"
)

func Benchmark(b *testing.B) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
	slog.SetDefault(l)

	eml, err := embedFSTestadata.ReadFile("testdata/small.eml")
	require.NoError(b, err)

	_, s1, addr1, err := testServer(nil, server.WithEnableCHUNKING(true))
	require.NoError(b, err)

	b.Run("SmallWithChunking", func(b *testing.B) {
		b.SetBytes(int64(len(eml)))
		for b.Loop() {
			_ = sendMail(addr1, eml)
		}
	})

	b.Run("SmallWithChunkingSameConnection", func(b *testing.B) {
		b.SetBytes(int64(len(eml)))
		c := client.New(client.WithServerAddresses(addr1), client.WithSecurity(client.SecurityPlain))
		require.NotNil(b, c)
		require.NoError(b, c.Connect(context.Background()))

		for b.Loop() {
			_ = sendMailCon(c, eml)
		}

		err = c.Quit()
		require.NoError(b, err)
	})

	require.NoError(b, s1.Close())

	_, s2, addr2, err := testServer(nil, server.WithEnableCHUNKING(false))
	require.NoError(b, err)

	b.Run("SmallWithoutChunking", func(b *testing.B) {
		b.SetBytes(int64(len(eml)))

		for b.Loop() {
			_ = sendMail(addr2, eml)
		}
	})

	b.Run("SmallWithoutChunkingSameConnection", func(b *testing.B) {
		b.SetBytes(int64(len(eml)))

		c := client.New(client.WithServerAddresses(addr2), client.WithSecurity(client.SecurityPlain))
		require.NotNil(b, c)

		require.NoError(b, c.Connect(context.Background()))

		for b.Loop() {
			_ = sendMailCon(c, eml)
		}

		err = c.Quit()
		require.NoError(b, err)
	})
	require.NoError(b, s2.Close())

	// require.EqualValues(b, be1.messages, be2.messages)
}
