package login

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/domain/login"
	login2 "github.com/Micheli97/crud-campeonato-golang/entity/login"
	user2 "github.com/Micheli97/crud-campeonato-golang/repository/convert/login"
)

func (login *loginRepository) LoginRepository(email, password string) (login.LoginDomainInterface, *rest_err.RestErr) {
	db := login.databaseConnection

	query := `SELECT id, name, email FROM t_user WHERE email=$1 AND password=$2 `

	loginEntity := &login2.LoginEntity{}

	err := db.QueryRow(query, email, password).Scan(&loginEntity.ID, &loginEntity.Name, &loginEntity.Email)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter login no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return user2.ConvertLoginEntityToDomain(*loginEntity), nil
}

func LoginRepository(database *sql.DB) LoginRepositoryInterface {
	return &loginRepository{
		databaseConnection: database,
	}
}
