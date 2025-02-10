package irdata

import (
	"context"
	"fmt"
)

func (d *irdata) Car() DataCar {
	return &irdataCar{parent: d}
}

type irdataCar struct {
	parent *irdata
}

type DataCar interface {
	Get(ctx context.Context) ([]Car, error)
}

func (c *irdataCar) Get(ctx context.Context) ([]Car, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/car/get", d.membersUrl))
	var output []Car
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
