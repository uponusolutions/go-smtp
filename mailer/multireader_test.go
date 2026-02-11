package mailer_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/mailer"
)

func TestClient_Multireader(t *testing.T) {
	a := bytes.NewBuffer([]byte{'a', 'b'})
	b := bytes.NewBuffer([]byte{'c'})
	c := bytes.NewBuffer([]byte{'d'})

	res := mailer.MultiReader(a, b, c)

	require.Equal(t, 4, res.Len())

	bytes := make([]byte, 1)
	i, err := res.Read(bytes)
	require.Equal(t, i, 1)
	require.Equal(t, bytes, []byte{'a'})
	require.NoError(t, err)
	require.Equal(t, 3, res.Len())

	bytes = make([]byte, 2)
	i, err = res.Read(bytes)
	require.NoError(t, err)
	require.Equal(t, i, 1)
	require.Equal(t, 2, res.Len())
	i, err = res.Read(bytes[1:])
	require.NoError(t, err)
	require.Equal(t, i, 1)
	require.Equal(t, bytes, []byte{'b', 'c'})
	require.Equal(t, 1, res.Len())

	bytes = make([]byte, 1)
	i, err = res.Read(bytes)
	require.Equal(t, i, 1)
	require.Equal(t, bytes, []byte{'d'})
	require.NoError(t, err)
	require.Equal(t, 0, res.Len())
}
