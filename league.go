package irdata

import (
	"context"
	"fmt"
	"net/url"
)

func (d *irdata) League() DataLeague {
	return &irdataLeague{parent: d}
}

type irdataLeague struct {
	parent *irdata
}

type DataLeague interface {
	Get(ctx context.Context, leagueId int, opts ...LeagueGetOption) (League, error)
}

type LeagueGetOption interface {
	ApplyLeagueGet(*url.Values)
}

type IncludeLicensesOption struct {
	IncludeLicenses bool
}

func (o *IncludeLicensesOption) ApplyLeagueGet(v *url.Values) {
	v.Set("include_licenses", fmt.Sprintf("%t", o.IncludeLicenses))
}

func (c *irdataLeague) Get(ctx context.Context, leagueId int, opts ...LeagueGetOption) (League, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/league/get", d.membersUrl))
	if err != nil {
		return League{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	for _, opt := range opts {
		opt.ApplyLeagueGet(&q)
	}

	q.Set("league_id", fmt.Sprintf("%d", leagueId))
	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output League
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return League{}, err
	}

	return output, nil
}
