package postgres

import (
	"database/sql"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	_ "github.com/lib/pq"
	"os"
)

func InitializeDatabasePostgres() {

	credentialDB := os.Getenv("POSTGRESS_CREDENTIAL")

	db, err := sql.Open("postgres", credentialDB)

	if err != nil {
		logger.Error("Erro ao abrir a conexão com o PostgreSQL:", err)
		return

	}

	defer db.Close()

	db.Ping()

	logger.Info("Banco de dados iniciado....")

}
