package irdata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsCarClasses(t *testing.T) {
	api := DefaultClient.CarClass()
	assert.NotNil(t, api)

	values, err := api.Get()
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		if c.CarClassID == 0 && c.Name == "Hosted All Cars Class" {
			// car class 0 is "Hosted All Cars", test fails because 0 is "empty"
			continue
		}

		assert.NotEmpty(t, c.CarClassID)
		assert.NotEmpty(t, c.Name)
		assert.NotEmpty(t, c.CarsInClass)
	}
}
