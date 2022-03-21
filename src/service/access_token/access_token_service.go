package access_token

import (
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/repository/db"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/repository/rest"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/domain/access_token"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	dbRepository       db.DbRepository
	restUserRepository rest.RestUserRepository
}

func NewService(userRepo rest.RestUserRepository, dbRepo db.DbRepository) Service {
	return &service{
		dbRepository:       dbRepo,
		restUserRepository: userRepo,
	}
}

func (s *service) GetById(tokenId string) (*access_token.AccessToken, *errors.RestErr) {
	tokenId = strings.TrimSpace(tokenId)
	if len(tokenId) == 0 {
		return nil, errors.NewbadRequestError("Invalid token id, size of token should be greater then 1")
	}
	token, err := s.dbRepository.GetById(tokenId)
	if err != nil {
		return nil, err
	}
	return token, err
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, loginErr := s.restUserRepository.LoginUser(request.Username, request.Password)
	if loginErr != nil {
		return nil, loginErr
	}
	// Generate new accesstoken using user
	token := access_token.GetNewAccessToken(user.Id)
	token.Generate()
	if err := s.dbRepository.Create(token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepository.UpdateExpirationTime(at)
}
