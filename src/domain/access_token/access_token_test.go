package access_token

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokeConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, fmt.Sprintf("expirationTime should be set to 24 hrs, it is currently set at %d hrs", expirationTime))
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)
	assert.False(t, at.IsExpired(), "accesstoken shouldn't expire as soon as it is created")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
	assert.True(t, at.ClientId == 0, "new access token should not have an associated client id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring 3 hours from now should NOT be expired by default")
}
