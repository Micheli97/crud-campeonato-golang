package main

import (
	"database/sql"
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	repository2 "github.com/Micheli97/crud-campeonato-golang/model/repository/user"
	"github.com/Micheli97/crud-campeonato-golang/model/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := repository2.NewUserRespository(database)
	service := user.NewUserDomainService(repository)

	return user2.NewUserHandlerInterface(service)
}
