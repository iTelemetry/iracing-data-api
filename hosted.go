package irdata

import (
	"context"
	"fmt"
	"net/url"
)

func (d *irdata) Hosted() DataHosted {
	return &irdataHosted{parent: d}
}

type irdataHosted struct {
	parent *irdata
}

type DataHosted interface {
	GetSessions(ctx context.Context) (HostedSessions, error)
	GetCombinedSessions(ctx context.Context, opts ...HostedCombinedSessionsOption) (HostedSessions, error)
}

func (c *irdataHosted) GetSessions(ctx context.Context) (HostedSessions, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/hosted/sessions", d.membersUrl))
	var output HostedSessions
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return HostedSessions{}, err
	}

	return output, nil
}

type HostedCombinedSessionsOption interface {
	ApplyHostedCombinedSessions(*url.Values)
}

type PackageIDOption struct {
	PackageID int
}

func (o *PackageIDOption) ApplyHostedCombinedSessions(v *url.Values) {
	v.Set("package_id", fmt.Sprintf("%d", o.PackageID))
}

func (c *irdataHosted) GetCombinedSessions(ctx context.Context, opts ...HostedCombinedSessionsOption) (HostedSessions, error) {
	d := c.parent

	u, err := url.Parse(fmt.Sprintf("%s/data/hosted/combined_sessions", d.membersUrl))
	if err != nil {
		return HostedSessions{}, &ConfigurationError{Msg: "unable to parse URL", Trigger: err}
	}

	q := u.Query()
	for _, opt := range opts {
		opt.ApplyHostedCombinedSessions(&q)
	}

	u.RawQuery = q.Encode()

	resp, err := d.get(ctx, u.String())
	var output HostedSessions
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return HostedSessions{}, err
	}

	return output, nil
}
