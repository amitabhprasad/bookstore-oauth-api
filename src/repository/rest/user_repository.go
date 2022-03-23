package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/domain/users"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type userRepository struct {
}

func NewRepository() RestUserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User,
	rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("Invalid response when trying to login user", errors.New("restclient error"))
	}
	if response.StatusCode > 299 {
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("Invalid error interface when trying to login user", errors.New("json parsing error"))
		}
		return nil, restErr
	}
	var user users.User
	err := json.Unmarshal(response.Bytes(), &user)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("unable to marshal login response", errors.New("json parsing error"))
	}
	return &user, nil
}
