package irdata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Authenticator interface {
	Authenticate(username string, password string, force bool) (*url.URL, []*http.Cookie, time.Time, error)
}

type DefaultAuthenticator struct {
	Authenticator

	irdata IRData
}

func NewDefaultAuthenticator(irdata IRData) (*DefaultAuthenticator, error) {
	if irdata == nil {
		return nil, errors.New("irdata is nil")
	}

	return &DefaultAuthenticator{
		irdata: irdata,
	}, nil
}

func (a *DefaultAuthenticator) Authenticate(username string, password string, _ bool) (*url.URL, []*http.Cookie, time.Time, error) {
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
		Email:    username,
		Password: password,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, nil, time.Time{}, &ConfigurationError{Msg: "unable to marshal request body", Trigger: err}
	}

	bodyReader := bytes.NewReader(body)

	resp, err := a.irdata.HttpClient().Post(fmt.Sprintf("%s/auth", a.irdata.MembersUrl()), "application/json", bodyReader)
	if err != nil {
		return nil, nil, time.Time{}, &ConfigurationError{Msg: "unable to make request", Trigger: err}
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, nil, time.Time{}, &ConfigurationError{Msg: "unable to read authentication response", Trigger: err}
	}

	err = a.irdata.RateLimit().update(resp)
	if err != nil {
		return nil, nil, time.Time{}, &ConfigurationError{Msg: "unable to update rate limit", Trigger: err}
	}

	if resp.StatusCode == 401 {
		return nil, nil, time.Time{}, &AuthenticationError{Msg: "invalid credentials"}
	} else if resp.StatusCode == 503 {
		return nil, nil, time.Time{}, &ServiceUnavailableError{Msg: "service unavailable"}
	} else if resp.StatusCode == 209 {
		return nil, nil, time.Time{}, &RateLimitExceededError{Msg: "too many requests"}
	} else if resp.StatusCode != 200 {
		return nil, nil, time.Time{}, &ServiceUnavailableError{Msg: "unexpected error", Trigger: errors.New(resp.Status)}
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, time.Time{}, &ConfigurationError{Msg: "unable to read authentication response body", Trigger: err}
	}

	responseBody := AuthResponse{}
	err = json.Unmarshal(r, &responseBody)
	if err != nil {
		return nil, nil, time.Time{}, &ConfigurationError{Msg: "unable to unmarshal authentication response body", Trigger: err}
	}

	if responseBody.AuthCode == 0 || responseBody.AuthCode == float64(0) {
		return nil, nil, time.Time{}, &AuthenticationError{Msg: "authentication failed", Trigger: errors.New(responseBody.Message)}
	}

	expiresAt := time.Now()
	membersCookie := false
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "authtoken_members" {
			membersCookie = true
			expiresAt = cookie.Expires
			break
		}
	}

	if !membersCookie {
		return nil, nil, time.Time{}, &AuthenticationError{Msg: "unable to find 'authtoken_members' cookie"}
	}

	return resp.Request.URL, cookies, expiresAt, nil
}
