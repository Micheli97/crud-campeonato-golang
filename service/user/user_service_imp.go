package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/repository/user"
)

func NewUserDomainService(userRepository user2.UserRepository) UserDomainService {
	return &userService{userRepository}
}

func (user *userService) CreateUser(
	userDomain user.UserDomainInterface,
) (user.UserDomainInterface, *rest_err.RestErr) {
	return user.repository.CreateUser(userDomain)

}

func (user *userService) DeleteUser(userId string) *rest_err.RestErr {
	return user.repository.DeleteUser(userId)

}

func (user *userService) FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	return user.repository.FindUserByEmail(email)
}

func (user *userService) FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	return user.repository.FindUserByID(id)
}

func (user *userService) UpdateUser(id string, userDomain user.UserDomainInterface) *rest_err.RestErr {
	return user.repository.UpdateUser(id, userDomain)

}
