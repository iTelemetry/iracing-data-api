package irdata

import (
	"context"
	"fmt"
	"net/url"
)

// DataLeague interface defines the methods for interacting with leagues
type DataLeague interface {
	Get(ctx context.Context, leagueId int, opts ...LeagueGetOption) (League, error)
	Seasons(ctx context.Context, leagueId int, opts ...LeagueSeasonsOption) (LeagueSeasons, error)
	SeasonSessions(ctx context.Context, leagueId int, seasonId int, opts ...LeagueSeasonSessionsOption) (LeagueSeasonSessions, error)
}

func (d *irdata) League() DataLeague {
	return &irdataLeague{parent: d}
}

type irdataLeague struct {
	parent *irdata
}

// Get method and its options

// LeagueGetOption interface for options to the Get method
type LeagueGetOption interface {
	ApplyLeagueGet(*url.Values)
}

// IncludeLicensesOption implements LeagueGetOption
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

// Seasons method and its options

// LeagueSeasonsOption interface for options to the Seasons method
type LeagueSeasonsOption interface {
	ApplyLeagueSeasons(*url.Values)
}

// RetiredOption implements LeagueSeasonsOption
type RetiredOption struct {
	Retired bool
}

func (o *RetiredOption) ApplyLeagueSeasons(v *url.Values) {
	v.Set("retired", fmt.Sprintf("%t", o.Retired))
}

func (c *irdataLeague) Seasons(ctx context.Context, leagueId int, opts ...LeagueSeasonsOption) (LeagueSeasons, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/league/seasons", d.membersUrl))
	if err != nil {
		return LeagueSeasons{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	for _, opt := range opts {
		opt.ApplyLeagueSeasons(&q)
	}

	q.Set("league_id", fmt.Sprintf("%d", leagueId))
	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output LeagueSeasons
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return LeagueSeasons{}, err
	}

	return output, nil
}

// SeasonSessions method and its options

// LeagueSeasonSessionsOption interface for options to the SeasonSessions method
type LeagueSeasonSessionsOption interface {
	ApplyLeagueSeasonSessions(*url.Values)
}

// ResultsOnlyOption implements LeagueSeasonSessionsOption
type ResultsOnlyOption struct {
	ResultsOnly bool
}

func (o *ResultsOnlyOption) ApplyLeagueSeasonSessions(v *url.Values) {
	v.Set("results_only", fmt.Sprintf("%t", o.ResultsOnly))
}

func (c *irdataLeague) SeasonSessions(ctx context.Context, leagueId int, seasonId int, opts ...LeagueSeasonSessionsOption) (LeagueSeasonSessions, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/league/season_sessions", d.membersUrl))
	if err != nil {
		return LeagueSeasonSessions{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	for _, opt := range opts {
		opt.ApplyLeagueSeasonSessions(&q)
	}

	q.Set("league_id", fmt.Sprintf("%d", leagueId))
	q.Set("season_id", fmt.Sprintf("%d", seasonId))
	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output LeagueSeasonSessions
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return LeagueSeasonSessions{}, err
	}

	return output, nil
}
