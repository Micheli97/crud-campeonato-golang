package main

import (
	"github.com/Micheli97/crud-campeonato-golang/config/database/postgres"
	"github.com/Micheli97/crud-campeonato-golang/router"
)

func main() {
	databasePostgres, err := postgres.NewDatabasePostgres()

	if err != nil {
		return
	}

	userHandler := initDependencies(databasePostgres)

	router.SetupRouter(userHandler)

}
