package main

import (
	"database/sql"
	login2 "github.com/Micheli97/crud-campeonato-golang/handler/login"
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/Micheli97/crud-campeonato-golang/repository/login"
	user3 "github.com/Micheli97/crud-campeonato-golang/repository/user"
	login3 "github.com/Micheli97/crud-campeonato-golang/service/login"
	"github.com/Micheli97/crud-campeonato-golang/service/user"
)

func initDependencies(database *sql.DB) user2.UserHandlerInterface {
	repository := user3.NewUserRespository(database)
	service := user.NewUserDomainService(repository)

	return user2.NewUserHandlerInterface(service)
}

func initLoginDependencies(database *sql.DB) login2.LoginHandlerInterface {
	loginRepository := login.LoginRepository(database)
	loginService := login3.LoginService(loginRepository)
	return login2.LoginHandler(loginService)
}
