package irdata

import (
	"net/http"
	"time"
)

type IRData interface {
	Email() string
	IsLoggedIn() bool
	GetLoginExpiration() time.Time

	Car() DataCar
	CarClass() DataCarClass
	Constant() DataConstant
	Hosted() DataHosted
	Track() DataTrack
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
}

func (d *irdata) Email() string {
	return d.email
}

func (d *irdata) IsLoggedIn() bool {
	return d.expiration.After(time.Now())
}

func (d *irdata) needsReauthorization() bool {
	return d.expiration.Before(time.Now().Add(-d.reauthorizeThreshold))
}

func (d *irdata) GetLoginExpiration() time.Time {
	return d.expiration
}

func (d *irdata) get(url string) (resp *http.Response, err error) {
	if d.autoReauthorize && d.needsReauthorization() {
		err = d.Reauthenticate()
		if err != nil {
			return nil, err
		}
	} else if !d.IsLoggedIn() {
		return nil, &AuthenticationError{Msg: "not logged in"}
	}

	return d.client.Get(url)
}

func (d *irdata) Reauthenticate() error {
	if d.IsLoggedIn() {
		return nil
	}

	return d.Authenticate()
}
