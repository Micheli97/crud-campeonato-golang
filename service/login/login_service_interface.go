package login

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/login"
)

type LoginServiceInterface interface {
	LoginService(login login.LoginDomainInterface) (login.LoginDomainInterface, *rest_err.RestErr)
}
