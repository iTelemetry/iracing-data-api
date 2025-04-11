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

func TestReturnsLeagueSeasons(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	seasons, err := api.Seasons(context.TODO(), 6529)
	assert.NoError(t, err)
	assert.NotEmpty(t, seasons)

	// Check that essential fields are not empty
	assert.NotEmpty(t, seasons.LeagueID)
	assert.NotEmpty(t, seasons.Seasons)
	assert.Equal(t, 6529, seasons.LeagueID)

	// Check that at least one season has essential fields
	if len(seasons.Seasons) > 0 {
		season := seasons.Seasons[0]
		assert.NotEmpty(t, season.LeagueID)
		assert.NotEmpty(t, season.SeasonID)
		assert.NotEmpty(t, season.SeasonName)
	}
}

func TestReturnsLeagueSeasonsWithRetired(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	seasons, err := api.Seasons(context.TODO(), 6529, &RetiredOption{Retired: true})
	assert.NoError(t, err)
	assert.NotEmpty(t, seasons)

	// Check that essential fields are not empty
	assert.NotEmpty(t, seasons.LeagueID)
	assert.NotEmpty(t, seasons.Seasons)
	assert.Equal(t, 6529, seasons.LeagueID)
}

func TestReturnsLeagueSeasonSessions(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	sessions, err := api.SeasonSessions(context.TODO(), 6529, 110898)
	assert.NoError(t, err)
	assert.NotEmpty(t, sessions)

	// Check that essential fields are not empty
	assert.NotEmpty(t, sessions.Sessions)

	// Check that at least one session has essential fields
	if len(sessions.Sessions) > 0 {
		session := sessions.Sessions[0]
		assert.NotEmpty(t, session.LeagueID)
		assert.NotEmpty(t, session.LeagueSeasonID)
		assert.NotEmpty(t, session.SessionID)
	}
}

func TestReturnsLeagueSeasonSessionsWithResultsOnly(t *testing.T) {
	api := DefaultClient.League()
	assert.NotNil(t, api)

	sessions, err := api.SeasonSessions(context.TODO(), 6529, 110898, &ResultsOnlyOption{ResultsOnly: true})
	assert.NoError(t, err)
	assert.NotEmpty(t, sessions)

	// Check that essential fields are not empty
	assert.NotEmpty(t, sessions.Sessions)

	// Check that all sessions have results
	for _, session := range sessions.Sessions {
		assert.True(t, session.HasResults)
	}
}
