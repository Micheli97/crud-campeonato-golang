package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/user"
)

type UserRepository interface {
	CreateUser(domainInterface user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (user.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, domainInterface user.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
