package irdata

import (
	"context"
	"fmt"
	"net/url"
)

func (d *irdata) Results() DataResults {
	return &irdataResults{parent: d}
}

type irdataResults struct {
	parent *irdata
}

type DataResults interface {
	Get(ctx context.Context, subsessionId int, opts ...ResultsGetOption) (ResultSession, error)
}

type ResultsGetOption interface {
	ApplyResultsGet(*url.Values)
}

type ResultsFromOption struct {
	IncludeLicenses bool
}

func (o *ResultsFromOption) ApplyResultsGet(v *url.Values) {
	v.Set("include_licenses", fmt.Sprintf("%t", o.IncludeLicenses))
}

func (c *irdataResults) Get(ctx context.Context, subsessionId int, opts ...ResultsGetOption) (ResultSession, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/results/get", d.membersUrl))
	if err != nil {
		return ResultSession{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	q.Set("subsession_id", fmt.Sprintf("%d", subsessionId))
	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output ResultSession
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return ResultSession{}, err
	}

	return output, nil
}
