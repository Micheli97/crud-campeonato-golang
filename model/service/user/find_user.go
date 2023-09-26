package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userDomainService) FindUser(userId string) (*user.UserDomainInterface, *rest_err.RestErr) {

	return nil, nil
}
