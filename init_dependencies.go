package main

import (
	"database/sql"
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	user3 "github.com/Micheli97/crud-campeonato-golang/repository/user"
	"github.com/Micheli97/crud-campeonato-golang/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := user3.NewUserRespository(database)
	service := user.NewUserDomainService(repository)

	return user2.NewUserHandlerInterface(service)
}
