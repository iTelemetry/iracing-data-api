package irdata

import "net/http"

type Options interface {
	Apply(data *irdata) error
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

type OptionsMembersUrl struct {
	MembersUrl string
}

func (o OptionsMembersUrl) Apply(data *irdata) error {
	data.membersUrl = o.MembersUrl
	return nil
}
