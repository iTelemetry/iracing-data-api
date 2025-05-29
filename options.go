package irdata

import (
	"fmt"
	"net/http"
	"sync"
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

func WithRateLimitWait(maxWait time.Duration) Options {
	return OptionsRateLimitWait{MaxWait: maxWait}
}

type OptionsRateLimitWait struct {
	MaxWait time.Duration
}

func (o OptionsRateLimitWait) Apply(data *irdata) error {
	data.rateLimit.wait = true
	data.rateLimit.waitTimeout = o.MaxWait
	return nil
}

func WithRateLimitLocking(enabled bool) Options {
	return OptionsRateLimitLocking{Enabled: enabled}
}

type OptionsRateLimitLocking struct {
	Enabled bool
}

func (o OptionsRateLimitLocking) Apply(data *irdata) error {
	data.rateLimit.locking = o.Enabled
	if !o.Enabled {
		data.rateLimit.lock = nil
	} else {
		data.rateLimit.lock = new(sync.Mutex)
	}

	return nil
}

func WithRateLimitRetry(attempts int) Options {
	return OptionsRateLimitRetry{Attempts: attempts}
}

type OptionsRateLimitRetry struct {
	Attempts int
}

func (o OptionsRateLimitRetry) Apply(data *irdata) error {
	if o.Attempts < 1 {
		return fmt.Errorf("attempts must be greater than zero")
	}

	data.rateLimit.attempts = o.Attempts
	return nil
}

func WithAuthenticator(f func(ird IRData, def Authenticator) Authenticator) Options {
	return OptionsAuthenticator{
		AuthenticatorInit: f,
	}
}

type OptionsAuthenticator struct {
	AuthenticatorInit func(ird IRData, def Authenticator) Authenticator
}

func (o OptionsAuthenticator) Apply(data *irdata) error {
	if o.AuthenticatorInit == nil {
		return &ConfigurationError{Msg: "authenticator init func must not be nil"}
	}

	data.authenticator = o.AuthenticatorInit(data, data.authenticator)
	if data.authenticator == nil {
		return &ConfigurationError{Msg: "authenticator init func result must not be nil"}
	}

	return nil
}
