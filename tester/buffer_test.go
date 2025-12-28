package tester

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuffer(t *testing.T) {
	buf := NewBuffer([]byte{'a', 'b', 'c'})

	cache := make([]byte, 2)

	i, err := buf.Read(cache)
	require.NoError(t, err)
	require.Equal(t, 2, i)
	require.Equal(t, []byte{'a', 'b'}, cache)

	i, err = buf.Read(cache)
	require.NoError(t, err)
	require.Equal(t, 1, i)
	require.Equal(t, []byte{'c'}, cache[:i])

	i, err = buf.Read(cache)
	require.Error(t, io.EOF, err)
	require.Equal(t, 0, i)
}
