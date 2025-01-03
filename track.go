package irdata

import (
	"fmt"
)

func (d *irdata) Track() DataTrack {
	return &irdataTrack{parent: d}
}

type irdataTrack struct {
	parent *irdata
}

type DataTrack interface {
	Get() ([]Track, error)
	Assets() (map[string]TrackAssets, error)
}

func (c *irdataTrack) Get() ([]Track, error) {
	d := c.parent

	resp, err := d.get(fmt.Sprintf("%s/data/track/get", d.membersUrl))
	var output []Track
	err = handleLink(d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *irdataTrack) Assets() (map[string]TrackAssets, error) {
	d := c.parent

	resp, err := d.get(fmt.Sprintf("%s/data/track/assets", d.membersUrl))
	var output map[string]TrackAssets
	err = handleLink(d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
