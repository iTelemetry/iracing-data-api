package irdata

import "fmt"

func (d *irdata) Constant() DataConstant {
	return &irdataConstant{parent: d}
}

type irdataConstant struct {
	parent *irdata
}

type DataConstant interface {
	GetCategories() ([]Category, error)
	GetDivisions() ([]Division, error)
	GetEventTypes() ([]EventType, error)
}

func (c *irdataConstant) GetCategories() ([]Category, error) {
	d := c.parent

	resp, err := d.get(fmt.Sprintf("%s/data/constants/categories", d.membersUrl))
	var output []Category
	err = handleResponse(resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *irdataConstant) GetDivisions() ([]Division, error) {
	d := c.parent

	resp, err := d.get(fmt.Sprintf("%s/data/constants/divisions", d.membersUrl))
	var output []Division
	err = handleResponse(resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *irdataConstant) GetEventTypes() ([]EventType, error) {
	d := c.parent

	resp, err := d.get(fmt.Sprintf("%s/data/constants/event_types", d.membersUrl))
	var output []EventType
	err = handleResponse(resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
