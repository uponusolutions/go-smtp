package limit

import (
	"errors"
	"time"
)

// ErrRatelimit is returned if limit reached and strict mode is enabled.
var ErrRatelimit = errors.New("rate limit occured")

type RatelimitConfig struct {
	Rate     int
	Duration time.Duration
	Strict   bool
}

type Ratelimit struct {
	start  time.Time
	count  int
	config *RatelimitConfig
}

func New(config *RatelimitConfig) *Ratelimit {
	return &Ratelimit{
		config: config,
		start:  time.Now(),
		count:  0,
	}
}

func (c *Ratelimit) Take() error {
	c.count++
	if c.count <= c.config.Rate {
		return nil
	}
	now := time.Now()

	dur := now.Sub(c.start)

	if dur < c.config.Duration {
		if c.config.Strict {
			return ErrRatelimit
		}
		time.Sleep(c.config.Duration - dur)
		now = time.Now()
	}

	c.start = now
	c.count = 1

	return nil
}
