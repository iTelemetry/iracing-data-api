package irdata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsLeague(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	league, err := api.Get(context.TODO(), 8244)
	assert.NoError(t, err)
	assert.NotEmpty(t, league)

	// Check that essential fields are not empty
	assert.NotEmpty(t, league.LeagueID)
	assert.NotEmpty(t, league.LeagueName)
}

func TestReturnsLeagueWithMultipleLeagueIds(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	leagueIds := []int{8244, 7198, 6529, 10654}

	for _, leagueId := range leagueIds {
		league, err := api.Get(context.TODO(), leagueId)
		assert.NoError(t, err)
		assert.NotEmpty(t, league)

		// Check that essential fields are not empty
		assert.NotEmpty(t, league.LeagueID)
		assert.NotEmpty(t, league.LeagueName)

		assert.Equal(t, leagueId, league.LeagueID)
	}
}

func TestReturnsLeagueWithIncludeLicenses(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	league, err := api.Get(context.TODO(), 7198, &IncludeLicensesOption{IncludeLicenses: true})
	assert.NoError(t, err)
	assert.NotEmpty(t, league)

	// Check that essential fields are not empty
	assert.NotEmpty(t, league.LeagueID)
	assert.NotEmpty(t, league.LeagueName)
}
