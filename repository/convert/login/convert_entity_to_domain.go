package login

import (
	login2 "github.com/Micheli97/crud-campeonato-golang/domain/login"
	"github.com/Micheli97/crud-campeonato-golang/entity/login"
)

func ConvertLoginEntityToDomain(loginEntity login.LoginEntity) login2.LoginDomainInterface {
	domain := login2.NewLoginDomain(
		loginEntity.Email, loginEntity.Name, loginEntity.Password)

	domain.SetID(loginEntity.ID)
	return domain
}
