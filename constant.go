package irdata

import (
	"context"
	"fmt"
)

func (d *irdata) Constant() DataConstant {
	return &irdataConstant{parent: d}
}

type irdataConstant struct {
	parent *irdata
}

type DataConstant interface {
	GetCategories(ctx context.Context) ([]Category, error)
	GetDivisions(ctx context.Context) ([]Division, error)
	GetEventTypes(ctx context.Context) ([]EventType, error)
}

func (c *irdataConstant) GetCategories(ctx context.Context) ([]Category, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/constants/categories", d.membersUrl))
	var output []Category
	err = handleResponse(resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *irdataConstant) GetDivisions(ctx context.Context) ([]Division, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/constants/divisions", d.membersUrl))
	var output []Division
	err = handleResponse(resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *irdataConstant) GetEventTypes(ctx context.Context) ([]EventType, error) {
	d := c.parent

	resp, err := d.get(ctx, fmt.Sprintf("%s/data/constants/event_types", d.membersUrl))
	var output []EventType
	err = handleResponse(resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
