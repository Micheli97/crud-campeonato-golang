package main

import (
	"database/sql"
	login2 "github.com/Micheli97/crud-campeonato-golang/handler/login"
	team2 "github.com/Micheli97/crud-campeonato-golang/handler/team"
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/Micheli97/crud-campeonato-golang/repository/login"
	"github.com/Micheli97/crud-campeonato-golang/repository/team"
	user3 "github.com/Micheli97/crud-campeonato-golang/repository/user"
	login3 "github.com/Micheli97/crud-campeonato-golang/service/login"
	team3 "github.com/Micheli97/crud-campeonato-golang/service/team"
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

func initTeamDependencies(database *sql.DB) team2.TeamHandlerInterface {
	teamRepository := team.TeamRepository(database)
	teamService := team3.TeamService(teamRepository)
	return team2.NewTeamHandler(teamService)

}
