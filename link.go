package irdata

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type linkResponse struct {
	Link    string    `json:"link"`
	Expires time.Time `json:"expires"`
}

func handleLink[T any](d *irdata, resp *http.Response, err error, output T) error {
	defer resp.Body.Close()

	if err != nil {
		return &ConfigurationError{Msg: "unable to make link request", Trigger: err}
	}

	if resp.StatusCode == http.StatusServiceUnavailable {
		return &ServiceUnavailableError{Msg: "service unavailable", Trigger: errors.New(resp.Status)}
	} else if resp.StatusCode == http.StatusForbidden {
		return &AuthenticationError{Msg: "forbidden"}
	} else if resp.StatusCode == http.StatusUnauthorized {
		return &AuthenticationError{Msg: "unauthorised"}
	} else if resp.StatusCode != http.StatusOK {
		return &ConfigurationError{Msg: "unexpected status code", Trigger: errors.New(resp.Status)}
	}

	linkBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ConfigurationError{Msg: "unable to read link response body", Trigger: err}
	}

	var link linkResponse
	err = json.Unmarshal(linkBody, &link)
	if err != nil {
		return &ConfigurationError{Msg: "unable to unmarshal link response body", Trigger: err}
	}

	r, err := d.client.Get(link.Link)
	if err != nil {
		return &ConfigurationError{Msg: "unable to make request", Trigger: err}
	}

	return handleResponse(r, err, output)
}

func handleResponse[T any](resp *http.Response, err error, output T) error {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &ConfigurationError{Msg: "unexpected status code", Trigger: errors.New(resp.Status)}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ConfigurationError{Msg: "unable to read response body", Trigger: err}
	}

	err = json.Unmarshal(body, output)
	if err != nil {
		return &ConfigurationError{Msg: "unable to unmarshal response body", Trigger: err}
	}

	return nil
}
