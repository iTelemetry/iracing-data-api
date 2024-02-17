package irdata

import "fmt"

func (d *irdata) Car() DataCar {
	return &irdataCar{parent: d}
}

type irdataCar struct {
	parent *irdata
}

type DataCar interface {
	Get() ([]Car, error)
}

func (c *irdataCar) Get() ([]Car, error) {
	d := c.parent

	resp, err := d.client.Get(fmt.Sprintf("%s/data/car/get", d.membersUrl))
	var output []Car
	err = handleLink(d, resp, err, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
