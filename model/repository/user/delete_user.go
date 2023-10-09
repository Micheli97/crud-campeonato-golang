package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
)

func (user *userRespository) DeleteUser(id string) *rest_err.RestErr {
	db := user.databaseConnection

	query := `DELETE FROM t_user WHERE id=$1`

	_, err := db.Exec(query, id)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar cadastrar usuário no banco de dados", err)
		return nil
	}

	return nil

}
