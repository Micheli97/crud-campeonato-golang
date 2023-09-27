package main

import (
	"github.com/Micheli97/crud-campeonato-golang/config/database/postgres"
	"github.com/Micheli97/crud-campeonato-golang/router"
)

func main() {
	postgres.NewDatabasePostgres()
	router.SetupRouter()

}
