package main

import (
	"database/sql"
	userHandler "github.com/Micheli97/crud-campeonato-golang/handler/user"
	repository2 "github.com/Micheli97/crud-campeonato-golang/model/repository"
	"github.com/Micheli97/crud-campeonato-golang/model/service/user"
)

func initDependencies(database *sql.DB) userHandler.UserHandlerInterface {
	repository := repository2.NewUserRespository(database)
	service := user.NewUserDomainService(repository)
	return userHandler.NewUserHandlerInterface(service)
}
