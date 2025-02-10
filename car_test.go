package irdata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsCars(t *testing.T) {
	car := DefaultClient.Car()
	assert.NotNil(t, car)

	cars, err := car.Get(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, cars)

	for _, c := range cars {
		assert.NotEmpty(t, c.CarID)
		assert.NotEmpty(t, c.CarName)
		assert.NotEmpty(t, c.Created)
	}
}
