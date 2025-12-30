package resolvemx

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeResolver struct{}

func (*fakeResolver) LookupMX(_ context.Context, name string) ([]*net.MX, error) {
	switch name {
	case "y.local":
		return nil, &net.DNSError{IsNotFound: false, Name: "realerror"}
	case "z.local":
		return nil, &net.DNSError{IsNotFound: true, Name: "notfound"}
	case "a.local":
		return []*net.MX{
			{
				Host: "smtpa.local",
				Pref: 1,
			},
		}, nil
	case "b.local", "d.local":
		return []*net.MX{
			{
				Host: "smtpa.local",
				Pref: 1,
			},
			{
				Host: "smtpb.local",
				Pref: 2,
			},
		}, nil
	case "c.local":
		return []*net.MX{
			{
				Host: "smtpa.local",
				Pref: 1,
			},
			{
				Host: "smtpb.local",
				Pref: 1,
			},
			{
				Host: "smtpc.local",
				Pref: 2,
			},
			{
				Host: "smtpd.local",
				Pref: 2,
			},
			{
				Host: "smtpe.local",
				Pref: 3,
			},
		}, nil
	}

	return nil, nil
}

func TestMxResolveError(t *testing.T) {
	resolveMx := New(&fakeResolver{})
	_, err := resolveMx.Lookup(context.Background(), "y.local")
	require.ErrorContains(t, err, "realerror")

	_, err = resolveMx.Recipients(context.Background(), []string{"test@a.local", "test@y.local"})
	require.ErrorContains(t, err, "realerror")

	res, err := resolveMx.Lookup(context.Background(), "z.local")
	require.NoError(t, err)
	require.True(t, res == nil)
}

func TestMxResolve(t *testing.T) {
	resolveMx := New(&fakeResolver{})
	res, err := resolveMx.Lookup(context.Background(), "a.local")
	require.NoError(t, err)
	require.Equal(t, [][]string{{"smtpa.local:25"}}, res)

	res, err = resolveMx.Lookup(context.Background(), "b.local")
	require.NoError(t, err)
	require.Equal(t, [][]string{{"smtpa.local:25"}, {"smtpb.local:25"}}, res)

	res, err = resolveMx.Lookup(context.Background(), "c.local")
	require.NoError(t, err)
	require.Equal(t, [][]string{
		{"smtpa.local:25", "smtpb.local:25"}, {"smtpc.local:25", "smtpd.local:25"}, {"smtpe.local:25"},
	}, res)
}

func TestMxResolveFail(t *testing.T) {
	resolveMx := New(nil)
	res, err := resolveMx.Lookup(context.Background(), "notexisting.mx-record.de")
	require.NoError(t, err)
	require.Nil(t, res)
}

func TestRecipients(t *testing.T) {
	resolveMx := New(&fakeResolver{})
	res, err := resolveMx.Recipients(context.Background(), []string{
		"test@a.local",
		"test@b.local",
		"test@c.local",
		"test2@c.local",
		"test@d.local",
		"test@z.local",
		"test2@z.local",
		"invalid",
	})
	require.NoError(t, err)

	require.Equal(t, 2, len(res.Failures))
	require.ErrorContains(t, res.Failures[0].Error, "no mx record found for z.local")
	require.Equal(t, []string{"test@z.local", "test2@z.local"}, res.Failures[0].Rcpts)
	require.ErrorContains(t, res.Failures[1].Error, "part of invalid")
	require.Equal(t, []string{"invalid"}, res.Failures[1].Rcpts)

	require.ErrorContains(t, res.Error(), "part of invalid")
	require.ErrorContains(t, res.Error(), "no mx record found for z.local")

	require.Equal(t, []Server{
		{
			Addresses: [][]string{{"smtpa.local:25"}},
			Rcpts:     []string{"test@a.local"},
		},
		{
			Addresses: [][]string{{"smtpa.local:25"}, {"smtpb.local:25"}},
			Rcpts:     []string{"test@b.local", "test@d.local"},
		},
		{
			Addresses: [][]string{
				{"smtpa.local:25", "smtpb.local:25"}, {"smtpc.local:25", "smtpd.local:25"}, {"smtpe.local:25"},
			},
			Rcpts: []string{"test@c.local", "test2@c.local"},
		},
	}, res.Servers)
}

func TestSliceEqual(t *testing.T) {
	require.True(t,
		sliceEqual(
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local"}},
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local"}},
		),
	)

	require.False(t,
		sliceEqual(
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local"}},
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpde.local"}},
		),
	)

	require.False(t,
		sliceEqual(
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local"}},
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local", "test.local"}},
		),
	)

	require.False(t,
		sliceEqual(
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local"}, {"test"}},
			[][]string{{"smtpa.local", "smtpb.local"}, {"smtpc.local", "smtpd.local"}},
		),
	)
}
