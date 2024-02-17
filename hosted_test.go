package irdata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsHostedSessions(t *testing.T) {
	api := DefaultClient.Hosted()
	assert.NotNil(t, api)

	r, err := api.GetSessions()
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	assert.True(t, r.Success)

	values := r.Sessions
	assert.NotNil(t, values)
	assert.NotEmpty(t, values)

	for _, c := range values {
		assert.NotEmpty(t, c.SessionID)
		assert.NotEmpty(t, c.SessionName)
		assert.NotEmpty(t, c.SessionTypes)
		assert.NotEmpty(t, c.CarTypes)
	}
}

func TestReturnsCombinedHostedSessions(t *testing.T) {
	api := DefaultClient.Hosted()
	assert.NotNil(t, api)

	tests := []struct {
		name                 string
		includePackageFilter bool
		packageID            int
		skipIfEmpty          bool
	}{
		{
			name:                 "GetCombinedSessionsWithBathurstPackageID",
			includePackageFilter: true,
			packageID:            146, // Bathurst
			skipIfEmpty:          true,
		},
		{
			name:                 "GetCombinedSessionsWithSpaPackageID",
			includePackageFilter: true,
			packageID:            103, // Spa (165: Endurance, 163: Grand Prix)
			skipIfEmpty:          true,
		},
		{
			name:                 "GetCombinedSessionsWithOkayamaPackageID",
			includePackageFilter: true,
			packageID:            109, // Okayama (166: Full Course)
			skipIfEmpty:          true,
		},
		{
			name:                 "GetCombinedSessionsWithDaytonaPackageID",
			includePackageFilter: true,
			packageID:            120, // Daytona (191: Oval, 192: Road Course)
			skipIfEmpty:          true,
		},
		{
			name:                 "GetCombinedSessionsWithNoPackageID",
			includePackageFilter: false,
			skipIfEmpty:          false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var opts []HostedCombinedSessionsOption
			if test.includePackageFilter {
				opts = append(opts, &PackageIDOption{PackageID: test.packageID})
			}

			r, err := api.GetCombinedSessions(opts...)
			assert.NoError(t, err)
			assert.True(t, r.Success)

			if test.skipIfEmpty && len(r.Sessions) == 0 {
				t.Skip("No sessions found for track")
				return
			} else if !test.skipIfEmpty && len(r.Sessions) == 0 {
				t.Error("No sessions found for track")
				return
			}

			values := r.Sessions
			assert.NotNil(t, values)
			assert.NotEmpty(t, values)

			for _, c := range values {
				assert.NotEmpty(t, c.SessionID)
				assert.NotEmpty(t, c.SessionName)
				assert.NotEmpty(t, c.SessionTypes)
				assert.NotEmpty(t, c.CarTypes)
			}
		})
	}
}
