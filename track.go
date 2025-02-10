package irdata

import (
	"context"
	"fmt"
)

func (d *irdata) Track() DataTrack {
	return &irdataTrack{parent: d}
}

type irdataTrack struct {
	parent *irdata
}

type DataTrack interface {
	Get(ctx context.Context) ([]Track, error)
	Assets(ctx context.Context) (map[string]TrackAssets, error)
}

func (c *irdataTrack) Get(ctx context.Context) ([]Track, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/track/get", d.membersUrl))
	var output []Track
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *irdataTrack) Assets(ctx context.Context) (map[string]TrackAssets, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/track/assets", d.membersUrl))
	var output map[string]TrackAssets
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
