package repository

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/user"
)

func NewUserRespository(database *sql.DB) UserRepository {
	return &userRespository{
		database,
	}
}

type userRespository struct {
	databaseConnection *sql.DB
}

type UserRepository interface {
	CreateUser(domainInterface user2.UserDomainInterface) (user2.UserDomainInterface, *rest_err.RestErr)
}
