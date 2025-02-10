package irdata

import (
	"context"
	"fmt"
)

func (d *irdata) CarClass() DataCarClass {
	return &irdataCarClass{parent: d}
}

type irdataCarClass struct {
	parent *irdata
}

type DataCarClass interface {
	Get(ctx context.Context) ([]CarClass, error)
}

func (c *irdataCarClass) Get(ctx context.Context) ([]CarClass, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/carclass/get", d.membersUrl))
	var output []CarClass
	err = handleLink(ctx, d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
