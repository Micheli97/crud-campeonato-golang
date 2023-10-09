package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/repository/user"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func NewUserDomainService(userRepository user2.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository user2.UserRepository
}

type UserDomainService interface {
	UpdateUser(string, user.UserDomainInterface) *rest_err.RestErr
	CreateUser(user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
