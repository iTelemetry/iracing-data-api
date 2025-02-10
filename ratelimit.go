package irdata

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RateLimit struct {
	totalLimit int
	remaining  int
	resetAt    time.Time

	attempts int

	wait        bool
	waitTimeout time.Duration

	locking bool
	lock    *sync.Mutex
}

func (rl *RateLimit) TotalLimit() int {
	return rl.totalLimit
}

func (rl *RateLimit) Remaining() int {
	return rl.remaining
}

func (rl *RateLimit) ResetAt() time.Time {
	return rl.resetAt
}

func (rl *RateLimit) Attempts() int {
	if rl.attempts == 0 {
		rl.attempts = 1
	}

	return rl.attempts
}

func (rl *RateLimit) Wait(ctx context.Context) error {
	if !rl.wait || rl.remaining > 0 || rl.resetAt.IsZero() {
		return nil
	}

	requiredTimeout := rl.resetAt.Sub(time.Now())
	if requiredTimeout < rl.waitTimeout {
		delay := rl.resetAt.Sub(time.Now())
		// slog.Debug("waiting before reset", "delay", delay)

		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}

			return ctx.Err()
		case <-timer.C:
		}
	} else {
		return &RateLimitExceededError{
			Msg: fmt.Sprintf("rate limit exceeded, max wait timeout exceeded (%s), %s", rl.waitTimeout, requiredTimeout),
		}
	}

	return nil
}

func (rl *RateLimit) update(resp *http.Response) error {
	if rl.locking {
		rl.lock.Lock()
		defer rl.lock.Unlock()
	}

	v := resp.Header.Get("X-RateLimit-Limit")
	if v != "" {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid value for total rate limit (%s): %w", v, err)
		}

		rl.totalLimit = limit
	}

	v = resp.Header.Get("X-RateLimit-Remaining")
	if v != "" {
		remaining, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid value for remaining rate limit (%s): %v", v, err)
		}

		rl.remaining = remaining
	}

	v = resp.Header.Get("X-RateLimit-Reset")
	if v != "" {
		resetAtEpoch, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid rate limit reset (%s): %w", v, err)
		}

		resetAt := time.Unix(resetAtEpoch, 0).UTC()
		rl.resetAt = resetAt
	}

	return nil
}
