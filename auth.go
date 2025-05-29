package irdata

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AuthCode        interface{} `json:"authcode"`
	AutoLoginSeries string      `json:"autoLoginSeries"`
	AutoLoginToken  string      `json:"autoLoginToken"`
	CustId          int         `json:"custId"`
	Email           string      `json:"email"`
	SsoCookieDomain string      `json:"ssoCookieDomain"`
	SsoCookieName   string      `json:"ssoCookieName"`
	SsoCookiePath   string      `json:"ssoCookiePath"`
	SsoCookieValue  string      `json:"ssoCookieValue"`

	Message              string `json:"message"`
	Inactive             bool   `json:"inactive"`
	VerificationRequired bool   `json:"verificationRequired"`
}

func encodePassword(email string, password string) string {
	email = strings.ToLower(email)
	auth := fmt.Sprintf("%s%s", password, email)

	h := sha256.New()
	h.Write([]byte(auth))

	encodedHash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return encodedHash
}

func Login(email string, password string, options ...Options) (IRData, error) {
	if email == "" {
		return nil, &ConfigurationError{Msg: "email must not be empty"}
	} else if password == "" {
		return nil, &ConfigurationError{Msg: "password must not be empty"}
	}

	encodedHash := encodePassword(email, password)

	data := &irdata{
		client:               http.DefaultClient,
		membersUrl:           "https://members-ng.iracing.com",
		reauthorizeThreshold: 1 * time.Hour,

		email:        email,
		passwordHash: encodedHash,
		rateLimit: &RateLimit{
			remaining: 1,
		},
	}

	data.authenticator = &DefaultAuthenticator{
		irdata: data,
	}

	for _, option := range options {
		err := option.Apply(data)
		if err != nil {
			var cerr *ConfigurationError
			if errors.As(err, &cerr) {
				return nil, cerr
			}

			return nil, &ConfigurationError{Msg: fmt.Sprintf("unable to apply option %T", option), Trigger: err}
		}
	}

	err := data.Authenticate(false)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *irdata) Authenticate(force bool) error {
	cookieUrl, cookies, expiration, err := d.authenticator.Authenticate(d.email, d.passwordHash, force)
	if err != nil {
		return err
	}

	d.cookies = cookies
	d.expiration = expiration
	if d.client.Jar == nil {
		d.client.Jar, _ = cookiejar.New(&cookiejar.Options{})
	}

	d.client.Jar.SetCookies(cookieUrl, d.cookies)
	return nil
}
