package irdata

import (
	"net/http"
	"time"
)

type Options interface {
	Apply(data *irdata) error
}

func WithHttpClient(client *http.Client) Options {
	return OptionsHttpClient{HttpClient: client}
}

type OptionsHttpClient struct {
	HttpClient *http.Client
}

func (o OptionsHttpClient) Apply(data *irdata) error {
	if o.HttpClient == nil {
		return &ConfigurationError{Msg: "HttpClient must not be nil"}
	}

	data.client = o.HttpClient
	return nil
}

func WithMembersUrl(membersUrl string) Options {
	return OptionsMembersUrl{MembersUrl: membersUrl}
}

type OptionsMembersUrl struct {
	MembersUrl string
}

func (o OptionsMembersUrl) Apply(data *irdata) error {
	data.membersUrl = o.MembersUrl
	return nil
}

func WithAutoReauthorize(autoReauthorize bool) Options {
	return OptionsAutoReauthorize{AutoReauthorize: autoReauthorize}
}

type OptionsAutoReauthorize struct {
	AutoReauthorize bool
}

func (o OptionsAutoReauthorize) Apply(data *irdata) error {
	data.autoReauthorize = o.AutoReauthorize
	return nil
}

func WithAutoReauthorizeThreshold(reauthorizeThreshold time.Duration) Options {
	return OptionsAutoReauthorizeThreshold{ReauthorizeThreshold: reauthorizeThreshold}
}

type OptionsAutoReauthorizeThreshold struct {
	ReauthorizeThreshold time.Duration
}

func (o OptionsAutoReauthorizeThreshold) Apply(data *irdata) error {
	data.reauthorizeThreshold = o.ReauthorizeThreshold
	return nil
}
