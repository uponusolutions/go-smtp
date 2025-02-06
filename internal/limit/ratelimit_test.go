package limit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRateLimitStrict(t *testing.T) {
	limit := New(&RatelimitConfig{
		Duration: time.Duration(10 * time.Millisecond),
		Rate:     1,
		Strict:   true,
	})

	err := limit.Take()
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	err = limit.Take()
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	err = limit.Take()
	require.NoError(t, err)

	err = limit.Take()
	require.Equal(t, ErrRatelimit, err)
}

func TestRateLimitLax(t *testing.T) {
	limit := New(&RatelimitConfig{
		Duration: time.Duration(10 * time.Millisecond),
		Rate:     1,
		Strict:   false,
	})

	err := limit.Take()
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	err = limit.Take()
	require.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	now := time.Now()
	err = limit.Take()
	require.NoError(t, err)
	require.True(t, time.Since(now) < time.Duration(5*time.Millisecond))

	now = time.Now()
	err = limit.Take()
	require.NoError(t, err)
	require.True(t, time.Since(now) > time.Duration(5*time.Millisecond))

	time.Sleep(10 * time.Millisecond)

	now = time.Now()
	err = limit.Take()
	require.NoError(t, err)
	require.True(t, time.Since(now) < time.Duration(5*time.Millisecond))
}
