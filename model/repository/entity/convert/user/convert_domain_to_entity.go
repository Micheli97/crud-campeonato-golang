package user

import (
	user2 "github.com/Micheli97/crud-campeonato-golang/model/repository/entity/user"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func ConvertDomainToEntity(userDomain user.UserDomainInterface) *user2.UserEntity {
	return &user2.UserEntity{
		ID:       userDomain.GetID(),
		Name:     userDomain.GetName(),
		Password: userDomain.GetPassword(),
		Email:    userDomain.GetEmail(),
	}

}
