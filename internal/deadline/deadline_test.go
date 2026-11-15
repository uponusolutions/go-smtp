package deadline

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAt(t *testing.T) {
	require.True(t, at(0).IsZero())
	require.True(t, at(-time.Second).IsZero())

	before := time.Now()
	d := at(time.Minute)
	require.False(t, d.Before(before.Add(time.Minute)))
	require.False(t, d.After(time.Now().Add(time.Minute)))
}

func TestSet(t *testing.T) {
	c1, c2 := net.Pipe()
	defer func() { _ = c1.Close(); _ = c2.Close() }()

	buf := make([]byte, 1)
	var netErr net.Error

	Set(c1, time.Millisecond)
	_, err := c1.Read(buf)
	require.ErrorAs(t, err, &netErr)
	require.True(t, netErr.Timeout())

	SetRead(c1, 0) // no deadline anymore
	go func() { _, _ = c2.Write([]byte("x")) }()
	_, err = c1.Read(buf)
	require.NoError(t, err)

	SetWrite(c1, time.Millisecond)
	_, err = c1.Write(buf)
	require.ErrorAs(t, err, &netErr)
	require.True(t, netErr.Timeout())
}
