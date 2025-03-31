package irdata

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsResults(t *testing.T) {
	api := DefaultClient.Results()
	assert.NotNil(t, api)

	subsessionIds := []int{
		75902297,
		75902298,
		75902299,
		75902300,
		75902301,
		75902302,
	}

	for _, subsessionId := range subsessionIds {
		session, err := api.Get(context.TODO(), subsessionId)
		assert.NoError(t, err)
		assert.NotEmpty(t, session)
		assert.Equal(t, subsessionId, session.SubsessionID)

		assert.NotEmpty(t, session.CarClasses)
		assert.NotEmpty(t, session.RaceSummary)
		assert.NotEmpty(t, session.SessionResults)
	}
}
