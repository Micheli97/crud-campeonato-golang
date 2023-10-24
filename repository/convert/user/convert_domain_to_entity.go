package user

import (
	user3 "github.com/Micheli97/crud-campeonato-golang/domain/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/entity/user"
)

func ConvertDomainToEntity(userDomain user3.UserDomainInterface) *user2.UserEntity {
	return &user2.UserEntity{
		ID:       userDomain.GetID(),
		Name:     userDomain.GetName(),
		Password: userDomain.GetPassword(),
		Email:    userDomain.GetEmail(),
	}

}
