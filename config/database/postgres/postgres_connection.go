package postgres

import (
	"database/sql"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	_ "github.com/lib/pq"
	"os"
)

var (
	PostgressCredential = "POSTGRES_CREDENTIAL"
)

func NewDatabasePostgres() (*sql.DB, error) {

	credentialDB := os.Getenv(PostgressCredential)

	db, err := sql.Open("postgres", credentialDB)

	if err != nil {
		logger.Error("Erro ao abrir a conexão com o PostgreSQL:", err)
		return nil, err
	}

	//defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error("Erro na conexão com o PostgreSQL:", err)
	}

	logger.Info("Banco de dados iniciado....")

	return db, nil

}
