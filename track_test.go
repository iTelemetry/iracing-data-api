package irdata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsTracks(t *testing.T) {
	api := DefaultClient.Track()
	assert.NotNil(t, api)

	values, err := api.Get()
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.TrackID)
		assert.NotEmpty(t, c.TrackName)
		assert.NotEmpty(t, c.PackageID)
	}
}

func TestReturnsTrackAssets(t *testing.T) {
	api := DefaultClient.Track()
	assert.NotNil(t, api)

	values, err := api.Assets()
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.TrackID)
		assert.NotEmpty(t, c.TrackMap)
		assert.NotEmpty(t, c.TrackMapLayers)
	}
}
