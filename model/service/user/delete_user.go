package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
)

func (user *userDomainService) DeleteUser(userId string) *rest_err.RestErr {

	err := user.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
