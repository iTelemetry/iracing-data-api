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
}

type irdata struct {
	email        string
	passwordHash string

	client     *http.Client
	membersUrl string

	cookies    []*http.Cookie
	expiration time.Time
}

func (d *irdata) Email() string {
	return d.email
}

func (d *irdata) IsLoggedIn() bool {
	return d.expiration.After(time.Now())
}

func (d *irdata) GetLoginExpiration() time.Time {
	return d.expiration
}
