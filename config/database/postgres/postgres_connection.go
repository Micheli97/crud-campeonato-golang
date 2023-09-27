package postgres

import (
	"database/sql"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	_ "github.com/lib/pq"
	"os"
)

var (
	PostgressCredential = "POSTGRESS_CREDENTIAL"
)

func NewDatabasePostgres() (*sql.DB, error) {

	credentialDB := os.Getenv(PostgressCredential)

	db, err := sql.Open("postgres", credentialDB)

	if err != nil {
		logger.Error("Erro ao abrir a conexão com o PostgreSQL:", err)
		return nil, err
	}

	defer db.Close()

	db.Ping()

	logger.Info("Banco de dados iniciado....")

	return db, nil

}
