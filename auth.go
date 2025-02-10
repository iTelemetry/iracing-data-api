package irdata

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	err := data.Authenticate()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *irdata) Authenticate() error {
	/**
	 * data := url.Values{}
	 * data.Set("username", d.email)
	 * data.Set("password", d.passwordHash)
	 * resp, err := d.client.Post(fmt.Sprintf("%s/Login", d.membersUrl), "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	 * if err != nil {
	 * 	return &ConfigurationError{Msg: "unable to make request", Trigger: err}
	 * }
	 */

	requestBody := AuthRequest{
		Email:    d.email,
		Password: d.passwordHash,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return &ConfigurationError{Msg: "unable to marshal request body", Trigger: err}
	}

	bodyReader := bytes.NewReader(body)

	resp, err := d.client.Post(fmt.Sprintf("%s/auth", d.membersUrl), "application/json", bodyReader)
	if err != nil {
		return &ConfigurationError{Msg: "unable to make request", Trigger: err}
	}

	defer resp.Body.Close()

	if err != nil {
		return &ConfigurationError{Msg: "unable to read authentication response", Trigger: err}
	}

	err = d.RateLimit().update(resp)
	if err != nil {
		return &ConfigurationError{Msg: "unable to update rate limit", Trigger: err}
	}

	if resp.StatusCode == 401 {
		return &AuthenticationError{Msg: "invalid credentials"}
	} else if resp.StatusCode == 503 {
		return &ServiceUnavailableError{Msg: "service unavailable"}
	} else if resp.StatusCode == 209 {
		return &RateLimitExceededError{Msg: "too many requests"}
	} else if resp.StatusCode != 200 {
		return &ServiceUnavailableError{Msg: "unexpected error", Trigger: errors.New(resp.Status)}
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ConfigurationError{Msg: "unable to read authentication response body", Trigger: err}
	}

	responseBody := AuthResponse{}
	err = json.Unmarshal(r, &responseBody)
	if err != nil {
		return &ConfigurationError{Msg: "unable to unmarshal authentication response body", Trigger: err}
	}

	if responseBody.AuthCode == 0 || responseBody.AuthCode == float64(0) {
		return &AuthenticationError{Msg: "authentication failed", Trigger: errors.New(responseBody.Message)}
	}

	membersCookie := false
	d.cookies = resp.Cookies()
	for _, cookie := range d.cookies {
		if cookie.Name == "authtoken_members" {
			d.expiration = cookie.Expires
			membersCookie = true
			break
		}
	}

	if d.client.Jar == nil {
		d.client.Jar, _ = cookiejar.New(&cookiejar.Options{})
	}

	d.client.Jar.SetCookies(resp.Request.URL, d.cookies)

	if !membersCookie {
		return &AuthenticationError{Msg: "unable to find 'authtoken_members' cookie"}
	}

	return nil
}
