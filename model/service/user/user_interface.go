package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	UpdateUser(string, user.UserDomainInterface) *rest_err.RestErr
	CreateUser(user.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*user.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}