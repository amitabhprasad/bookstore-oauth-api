package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}
type AccessTokenRequest struct {
	GrantType string `json:grant_type`
	Scope     string `json:"scope"`

	// used for grant_type as password
	Username string `json:"username"`
	Password string `json:"password"`

	// used for client_credentials grant_type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at AccessTokenRequest) Validate() rest_errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameters ")
	}
	return nil
}

func (at *AccessToken) Validate() rest_errors.RestErr {
	tokenId := strings.TrimSpace(at.AccessToken)
	if len(tokenId) == 0 {
		return rest_errors.NewBadRequestError("Invalid token id, size of token should be greater then 1")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("Invalid User ID")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("Invalid Client ID")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}
func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:   userId,
		ClientId: 1,
		Expires:  time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).UTC().Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires)
}
