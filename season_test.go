package irdata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsSeasonList(t *testing.T) {
	api := DefaultClient.Season()
	assert.NotNil(t, api)

	year := 2024
	quarter := 2

	values, err := api.List(year, quarter)
	assert.NoError(t, err)
	assert.NotEmpty(t, values)
	assert.NotEmpty(t, values.Seasons)
	assert.Equal(t, year, values.SeasonYear)
	assert.Equal(t, quarter, values.SeasonQuarter)

	for _, c := range values.Seasons {
		assert.NotEmpty(t, c.SeriesID)
		assert.NotEmpty(t, c.SeriesName)
		assert.Equal(t, year, c.SeasonYear)
		assert.Equal(t, quarter, c.SeasonQuarter)
	}
}

func TestReturnsSeasonRaceGuide(t *testing.T) {
	api := DefaultClient.Season()
	assert.NotNil(t, api)

	values, err := api.RaceGuide()
	assert.NoError(t, err)
	assert.NotEmpty(t, values)
	assert.NotEmpty(t, values.Sessions)

	for _, c := range values.Sessions {
		assert.NotEmpty(t, c.SeriesID)
		assert.NotEmpty(t, c.SeasonID)
	}
}

func TestReturnsSeasonSpectatorSubSessionIDs(t *testing.T) {
	api := DefaultClient.Season()
	assert.NotNil(t, api)

	values, err := api.SpectatorSubSessionIDs()
	assert.NoError(t, err)
	assert.NotEmpty(t, values)
	assert.NotEmpty(t, values.SubsessionIds)
	assert.NotEmpty(t, values.EventTypes)
}

func TestReturnsSeasonSpectatorSubSessionIDsWithEventTypes(t *testing.T) {
	api := DefaultClient.Season()
	assert.NotNil(t, api)

	opts := []int{2, 4}
	values, err := api.SpectatorSubSessionIDs(&EventTypes{EventTypes: opts})
	assert.NoError(t, err)
	assert.NotEmpty(t, values)
	assert.True(t, values.Success)
	assert.NotEmpty(t, values.SubsessionIds)
	assert.NotEmpty(t, values.EventTypes)

	for _, c := range values.EventTypes {
		assert.Contains(t, opts, c)
	}
}
