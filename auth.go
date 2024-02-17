package irdata

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(email string, password string, options ...Options) (IRData, error) {
	if email == "" {
		return nil, &ConfigurationError{Msg: "email must not be empty"}
	} else if password == "" {
		return nil, &ConfigurationError{Msg: "password must not be empty"}
	}

	auth := fmt.Sprintf("%s:%s", password, email)
	encodedHash := base64.StdEncoding.EncodeToString([]byte(auth))

	data := &irdata{
		client:     http.DefaultClient,
		membersUrl: "https://members-ng.iracing.com",

		email:        email,
		passwordHash: encodedHash,
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

	if resp.StatusCode == 401 {
		return &AuthenticationError{Msg: "invalid credentials"}
	} else if resp.StatusCode == 503 {
		return &ServiceUnavailableError{Msg: "service unavailable"}
	} else if resp.StatusCode != 200 {
		return &ServiceUnavailableError{Msg: "unexpected error", Trigger: errors.New(resp.Status)}
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

	if !membersCookie {
		return &AuthenticationError{Msg: "unable to find 'authtoken_members' cookie"}
	}

	return nil
}
