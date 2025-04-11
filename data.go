package irdata

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type IRData interface {
	Email() string
	IsLoggedIn() bool
	GetLoginExpiration() time.Time

	RateLimit() *RateLimit
	Car() DataCar
	CarClass() DataCarClass
	Constant() DataConstant
	Hosted() DataHosted
	League() DataLeague
	Track() DataTrack
	Season() DataSeason
	Series() DataSeries
	Results() DataResults
}

type irdata struct {
	email        string
	passwordHash string

	client               *http.Client
	membersUrl           string
	autoReauthorize      bool
	reauthorizeThreshold time.Duration

	cookies    []*http.Cookie
	expiration time.Time

	rateLimit *RateLimit
}

func (d *irdata) Email() string {
	return d.email
}

func (d *irdata) IsLoggedIn() bool {
	return d.expiration.After(time.Now())
}

func (d *irdata) needsReauthorization() bool {
	return d.expiration.Before(time.Now().Add(d.reauthorizeThreshold))
}

func (d *irdata) GetLoginExpiration() time.Time {
	return d.expiration
}

func (d *irdata) RateLimit() *RateLimit {
	return d.rateLimit
}

func (d *irdata) get(ctx context.Context, url string) (resp *http.Response, err error) {
	for i := 0; i < d.rateLimit.Attempts(); i++ {
		//if i > 0 {
		//	slog.Debug("retrying", "attempt", i, "request", url)
		//}

		if err = d.RateLimit().Wait(ctx); err != nil {
			return nil, fmt.Errorf("waiting for rate limit: %w", err)
		}

		tmr := &RateLimitExceededError{}
		if err = d.Reauthenticate(); err != nil && errors.As(err, &tmr) {
			continue
		} else if err != nil {
			return
		}

		var req *http.Request
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}

		req.WithContext(ctx)
		resp, err = d.client.Do(req)
		if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
			continue
		} else {
			return
		}
	}

	return
}

func (d *irdata) Reauthenticate() error {
	if d.autoReauthorize && d.needsReauthorization() {
		err := d.Authenticate()
		if err != nil {
			return err
		}
	} else if !d.IsLoggedIn() {
		return &AuthenticationError{Msg: "not logged in"}
	}

	return nil
}
