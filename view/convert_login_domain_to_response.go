package view

import (
	"github.com/Micheli97/crud-campeonato-golang/domain/login"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/response"
)

func ConverteToLoginDomainResponse(login login.LoginDomainInterface) response.LoginResponse {
	return response.LoginResponse{
		ID:    login.GetID(),
		Name:  login.GetName(),
		Email: login.GetEmail(),
	}
}
