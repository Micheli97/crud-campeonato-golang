package login

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/login"
	login2 "github.com/Micheli97/crud-campeonato-golang/repository/login"
	"github.com/Micheli97/crud-campeonato-golang/utils"
)

func (loginService *loginService) LoginService(login login.LoginDomainInterface) (login.LoginDomainInterface, *rest_err.RestErr) {

	utils.EncryptPasswordLogin(login)

	user, err := loginService.loginRepository.LoginRepository(login.GetEmail(), login.GetPassword())

	if err != nil {
		return nil, err
	}

	//token, err := user.GenerateToken()
	//
	//if err != nil {
	//
	//	return nil, "", err
	//}
	//
	//println(token)
	//
	return user, nil

}

func LoginService(loginRepository login2.LoginRepositoryInterface) LoginServiceInterface {
	return &loginService{
		loginRepository: loginRepository,
	}
}
