package db

import (
	"fmt"

	"github.com/gocql/gocql"

	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/clients/cassandra"
	"github.com/amitabhprasad/bookstore-app/bookstore-oauth-api/src/domain/access_token"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
)

const (
	queryGetAccessToken          = "SELECT access_token, user_id, client_id, expires FROM access_tokens where access_token=?;"
	queryCreateAccessToken       = "INSERT into access_tokens (access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateAccessTokenExpiry = "UPDATE access_tokens set expires=? where access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
	CassandraLocalFunction()
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no access token found with the given id %s ", id))
		}
		return nil, rest_errors.NewInternalServerError("error when getting access_token", err)
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).
		Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when saving access_token", err)
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateAccessTokenExpiry,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when updating cexpiration time", err)
	}
	return nil
}

func (r *dbRepository) CassandraLocalFunction() {
	fmt.Println("All good")
}
