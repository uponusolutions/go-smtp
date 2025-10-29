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

type testcase struct {
	eml  []byte
	name string
}

func Benchmark(b *testing.B) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
	slog.SetDefault(l)

	smallEml, err := embedFSTestadata.ReadFile("testdata/small.eml")
	require.NoError(b, err)

	largeEml, err := embedFSTestadata.ReadFile("testdata/large.eml")
	require.NoError(b, err)

	for _, t := range []testcase{
		{
			eml:  smallEml,
			name: "Small",
		},
		{
			eml:  largeEml,
			name: "Large",
		},
	} {
		_, s1, addr1, err := testServer(nil, server.WithEnableCHUNKING(true))
		require.NoError(b, err)

		b.Run(t.name+"WithChunking", func(b *testing.B) {
			if os.Getenv("SETBYTES") == "" {
				b.SetBytes(int64(len(t.eml)))
			}
			for b.Loop() {
				_ = sendMail(addr1, t.eml)
			}
		})

		b.Run(t.name+"WithChunkingSameConnection", func(b *testing.B) {
			if os.Getenv("SETBYTES") == "" {
				b.SetBytes(int64(len(t.eml)))
			}
			c := client.New(
				client.WithServerAddresses(addr1),
				client.WithSecurity(client.SecurityPlain),
				client.WithMailOptions(client.MailOptions{Size: int64(len(t.eml))}),
			)
			require.NotNil(b, c)
			require.NoError(b, c.Connect(context.Background()))

			for b.Loop() {
				_ = sendMailCon(c, t.eml)
			}

			err = c.Quit()
			require.NoError(b, err)
		})

		require.NoError(b, s1.Close())

		_, s2, addr2, err := testServer(nil, server.WithEnableCHUNKING(false))
		require.NoError(b, err)

		b.Run(t.name+"WithoutChunking", func(b *testing.B) {
			if os.Getenv("SETBYTES") == "" {
				b.SetBytes(int64(len(t.eml)))
			}
			for b.Loop() {
				_ = sendMail(addr2, t.eml)
			}
		})

		b.Run(t.name+"WithoutChunkingSameConnection", func(b *testing.B) {
			if os.Getenv("SETBYTES") == "" {
				b.SetBytes(int64(len(t.eml)))
			}
			c := client.New(
				client.WithServerAddresses(addr2),
				client.WithSecurity(client.SecurityPlain),
				client.WithMailOptions(client.MailOptions{Size: int64(len(t.eml))}),
			)
			require.NotNil(b, c)

			require.NoError(b, c.Connect(context.Background()))

			for b.Loop() {
				_ = sendMailCon(c, t.eml)
			}

			err = c.Quit()
			require.NoError(b, err)
		})
		require.NoError(b, s2.Close())
	}

	// require.EqualValues(b, be1.messages, be2.messages)
}
