package user

import (
	user2 "github.com/Micheli97/crud-campeonato-golang/model/repository/entity/user"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func ConvertEntityToDomain(userEntity user2.UserEntity) user.UserDomainInterface {
	domain := user.NewUserDomain(
		userEntity.Email, userEntity.Password, userEntity.Name)

	domain.SetID(userEntity.ID)
	return domain
}
