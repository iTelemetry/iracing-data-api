package irdata

import "fmt"

func (d *irdata) CarClass() DataCarClass {
	return &irdataCarClass{parent: d}
}

type irdataCarClass struct {
	parent *irdata
}

type DataCarClass interface {
	Get() ([]CarClass, error)
}

func (c *irdataCarClass) Get() ([]CarClass, error) {
	d := c.parent

	resp, err := d.client.Get(fmt.Sprintf("%s/data/carclass/get", d.membersUrl))
	var output []CarClass
	err = handleLink(d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
