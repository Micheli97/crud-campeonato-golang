package user

import (
	user3 "github.com/Micheli97/crud-campeonato-golang/domain/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/entity/user"
)

func ConvertEntityToDomain(userEntity user2.UserEntity) user3.UserDomainInterface {
	domain := user3.NewUserDomain(
		userEntity.Email, userEntity.Password, userEntity.Name)

	domain.SetID(userEntity.ID)
	return domain
}
