package irdata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRateLimitUpdates(t *testing.T) {
	api := DefaultClient.Season()
	assert.NotNil(t, api)

	_, err := api.RaceGuide(context.TODO())
	assert.Nil(t, err)

	rl := DefaultClient.RateLimit()
	assert.NotNil(t, rl)

	remaining := rl.Remaining()
	total := rl.TotalLimit()
	reset := rl.ResetAt()

	assert.True(t, reset.After(time.Now()))

	_, err = api.RaceGuide(context.TODO())
	assert.Nil(t, err)

	assert.NotEqual(t, remaining, total, "remaining rate limit should not be equal to total")
	assert.NotEqual(t, rl.Remaining(), remaining, "remaining rate limit should not be equal to remaining on previous request")
	assert.True(t, reset.After(time.Now()))
}

//func TestRateLimit(t *testing.T) {
//	const workers = 5
//	const requests = 50
//
//	wg := new(sync.WaitGroup)
//
//	ch := make(chan int, requests)
//	for i := 0; i < workers; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//
//			for {
//				select {
//				case req, ok := <-ch:
//					if !ok {
//						return
//					}
//
//					car := DefaultClient.Car()
//					assert.NotNil(t, car)
//
//					cars, err := car.Get(context.TODO())
//					assert.NoError(t, err)
//					assert.NotEmpty(t, cars)
//
//					slog.Info("completed", "iteration", req, "remaining", DefaultClient.RateLimit().Remaining(), "reset", DefaultClient.RateLimit().ResetAt())
//				}
//			}
//		}()
//	}
//
//	for i := 0; i < requests; i++ {
//		ch <- i
//	}
//
//	close(ch)
//	wg.Wait()
//}
