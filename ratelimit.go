package irdata

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type RateLimit struct {
	totalLimit int
	remaining  int
	resetAt    time.Time
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

func (rl *RateLimit) update(resp *http.Response) error {
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
