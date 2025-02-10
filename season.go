package irdata

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func (d *irdata) Season() DataSeason {
	return &irdataSeason{parent: d}
}

type irdataSeason struct {
	parent *irdata
}

type DataSeason interface {
	List(ctx context.Context, year int, quarter int) (Seasons, error)
	RaceGuide(ctx context.Context, opts ...RaceGuideOption) (RaceGuide, error)
	SpectatorSubSessionIDs(ctx context.Context, opts ...SpectatorSubSessionIDsOption) (SpectatorSubSessionIDs, error)
}

func (c *irdataSeason) List(ctx context.Context, year int, quarter int) (Seasons, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/season/list", d.membersUrl))
	if err != nil {
		return Seasons{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	q.Set("season_year", fmt.Sprintf("%d", year))
	q.Set("season_quarter", fmt.Sprintf("%d", quarter))
	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output Seasons
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return Seasons{}, err
	}

	return output, nil
}

type RaceGuideOption interface {
	ApplyRaceGuide(*url.Values)
}

type FromOption struct {
	From time.Time
}

func (o *FromOption) ApplyRaceGuide(v *url.Values) {
	v.Set("from", fmt.Sprintf("%s", o.From.Format(time.RFC3339)))
}

type IncludeEndAfterFromOption struct {
	Include bool
}

func (o *IncludeEndAfterFromOption) ApplyRaceGuide(v *url.Values) {
	v.Set("include_end_after_from", fmt.Sprintf("%t", o.Include))
}

func (c *irdataSeason) RaceGuide(ctx context.Context, opts ...RaceGuideOption) (RaceGuide, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/season/race_guide", d.membersUrl))
	if err != nil {
		return RaceGuide{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	for _, opt := range opts {
		opt.ApplyRaceGuide(&q)
	}

	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output RaceGuide
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return RaceGuide{}, err
	}

	return output, nil
}

type SpectatorSubSessionIDsOption interface {
	ApplySpectatorSubSessionIDs(*url.Values)
}

type EventTypes struct {
	EventTypes []int
}

func (o *EventTypes) ApplySpectatorSubSessionIDs(v *url.Values) {
	str := make([]string, len(o.EventTypes))
	for i, et := range o.EventTypes {
		str[i] = fmt.Sprintf("%d", et)
	}

	v.Set("event_types", fmt.Sprintf("%s", strings.Join(str, ",")))
}

func (c *irdataSeason) SpectatorSubSessionIDs(ctx context.Context, opts ...SpectatorSubSessionIDsOption) (SpectatorSubSessionIDs, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/season/spectator_subsessionids", d.membersUrl))
	if err != nil {
		return SpectatorSubSessionIDs{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	for _, opt := range opts {
		opt.ApplySpectatorSubSessionIDs(&q)
	}

	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output SpectatorSubSessionIDs
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return SpectatorSubSessionIDs{}, err
	}

	return output, nil
}
