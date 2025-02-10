package irdata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsConstantCategories(t *testing.T) {
	api := DefaultClient.Constant()
	assert.NotNil(t, api)

	values, err := api.GetCategories(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.Label)
		assert.NotEmpty(t, c.Value)
	}
}

func TestReturnsConstantDivisions(t *testing.T) {
	api := DefaultClient.Constant()
	assert.NotNil(t, api)

	values, err := api.GetDivisions(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.Label)
		if c.Label != "Division 1" {
			assert.NotEmpty(t, c.Value)
		}
	}
}

func TestReturnsConstantEventTypes(t *testing.T) {
	api := DefaultClient.Constant()
	assert.NotNil(t, api)

	values, err := api.GetEventTypes(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.Label)
		assert.NotEmpty(t, c.Value)
	}
}
