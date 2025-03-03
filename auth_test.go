package irdata

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var DefaultClient IRData

var ValidEmail = ""
var ValidPassword = ""

func init() {
	email, emailOk := os.LookupEnv("IRDATA_EMAIL")
	password, pwOk := os.LookupEnv("IRDATA_PASSWORD")
	if !emailOk && !pwOk {
		panic("IRDATA_EMAIL and IRDATA_PASSWORD environment variables not set")
	} else if !emailOk {
		panic("IRDATA_EMAIL environment variable not set")
	} else if !pwOk {
		panic("IRDATA_PASSWORD environment variable not set")
	}

	ValidEmail = email
	ValidPassword = password

	var err error
	DefaultClient, err = Login(ValidEmail, ValidPassword, WithRateLimitWait(60*time.Second), WithRateLimitLocking(true), WithRateLimitRetry(5))
	if err != nil {
		panic(err)
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		password      string
		options       []Options
		errExpected   bool
		expectedEmail string
	}{
		{
			name:          "ValidLogin",
			email:         ValidEmail,
			password:      ValidPassword,
			options:       []Options{}, // No extra options
			errExpected:   false,
			expectedEmail: ValidEmail,
		},
		{
			name:          "InvalidEmail",
			email:         "",
			password:      "password",
			options:       []Options{},
			errExpected:   true, // We expect error because of empty email
			expectedEmail: "",
		},
		{
			name:          "InvalidClient",
			email:         "email",
			password:      "password",
			options:       []Options{OptionsHttpClient{HttpClient: nil}},
			errExpected:   true, // Expect error because client is nil
			expectedEmail: "",
		},
		{
			name:          "InvalidPassword",
			email:         "email",
			password:      "",
			options:       []Options{},
			errExpected:   true, // Expect error because client is nil
			expectedEmail: "",
		},
		// place other tests scenarios here depending on what you need to test.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Login(tt.email, tt.password, tt.options...)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedEmail, res.Email())
			}
		})
	}
}

func TestEncodePassword(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		password string
		want     string
	}{
		{
			name:     "iRacing example email and password",
			email:    "CLunky@iracing.Com",
			password: "MyPassWord",
			want:     "xGKecAR27ALXNuMLsGaG0v5Q9pSs2tZTZRKNgmHMg+Q=",
		},
		{
			name:     "iRacing 2nd example email and password",
			email:    "john.smith@iracing.com",
			password: "SuperSecure123",
			want:     "3NkLTzZtITXmFo7HHDwbGEnZox4VfwLMHHaNZzdNuE4=",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := encodePassword(tc.email, tc.password)

			if got != tc.want {
				t.Errorf("encodePassword(%v, %v) = %v; want %v", tc.email, tc.password, got, tc.want)
			}
		})
	}
}

func TestReauthorize(t *testing.T) {
	c, err := Login(ValidEmail, ValidPassword, WithAutoReauthorize(true), WithAutoReauthorizeThreshold(10*time.Minute))
	assert.NoError(t, err)

	data, ok := c.(*irdata)
	assert.True(t, ok)

	data.expiration = time.Now().Add(data.reauthorizeThreshold * 2)
	assert.False(t, data.needsReauthorization(), "should not need reauthorization")

	data.expiration = time.Now().Add(data.reauthorizeThreshold / 2)
	assert.True(t, data.needsReauthorization(), "should need reauthorization")

	oldExpiration := data.expiration

	err = data.Reauthenticate()
	assert.NoError(t, err)

	assert.True(t, data.IsLoggedIn(), "should be logged in")
	assert.NotEqual(t, oldExpiration, data.expiration, "expiration should have changed")
}
