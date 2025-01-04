package irdata

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRateLimitUpdates(t *testing.T) {
	api := DefaultClient.Season()
	assert.NotNil(t, api)

	_, err := api.RaceGuide()
	assert.Nil(t, err)

	rl := DefaultClient.RateLimit()
	assert.NotNil(t, rl)

	remaining := rl.Remaining()
	total := rl.TotalLimit()
	reset := rl.ResetAt()

	assert.True(t, reset.After(time.Now()))

	_, err = api.RaceGuide()
	assert.Nil(t, err)

	assert.NotEqual(t, remaining, total, "remaining rate limit should not be equal to total")
	assert.NotEqual(t, rl.Remaining(), remaining, "remaining rate limit should not be equal to remaining on previous request")
	assert.True(t, reset.After(time.Now()))
}
