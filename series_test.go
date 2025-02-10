package irdata

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReturnsSeriesAssets(t *testing.T) {
	api := DefaultClient.Series()
	assert.NotNil(t, api)

	values, err := api.Assets(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.SeriesID)
		assert.NotEmpty(t, c.SeriesCopy)
	}
}

func TestReturnsSeriesAssetsWithImageBaseUrl(t *testing.T) {
	api := DefaultClient.Series()
	assert.NotNil(t, api)

	url := "https://test.itelemetry.app"
	values, err := api.Assets(context.TODO(), WithImageBaseUrl(url))
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.SeriesID)
		assert.NotEmpty(t, c.SeriesCopy)

		if c.Logo != "" {
			assert.True(t, strings.HasPrefix(c.Logo, fmt.Sprintf("%s/", url)))
		}

		if c.LargeImage != "" {
			assert.True(t, strings.HasPrefix(c.LargeImage, fmt.Sprintf("%s/", url)))
		}

		if c.SmallImage != "" {
			assert.True(t, strings.HasPrefix(c.SmallImage, fmt.Sprintf("%s/", url)))
		}
	}
}

func TestReturnsSeries(t *testing.T) {
	api := DefaultClient.Series()
	assert.NotNil(t, api)

	values, err := api.Get(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.SeriesID)
		assert.NotEmpty(t, c.SeriesName)
	}
}

func TestReturnsSeriesSeasons(t *testing.T) {
	api := DefaultClient.Series()
	assert.NotNil(t, api)

	values, err := api.Seasons(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.SeriesID)
		assert.NotEmpty(t, c.SeasonName)
	}
}
