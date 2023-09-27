package repository

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userRespository) CreateUser(domainInterface user2.UserDomainInterface) (user2.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")
	//TODO adicionar logica para criar o usuario no banco de dados

	return nil, nil
}
