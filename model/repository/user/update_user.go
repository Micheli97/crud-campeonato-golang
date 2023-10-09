package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userRespository) UpdateUser(id string, userDomain user.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository")

	db := user.databaseConnection

	query := `UPDATE  t_user SET name=$1, email=$2 WHERE id=$3`

	_, err := db.Exec(query, userDomain.GetName(), userDomain.GetEmail(), id)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return nil
	}

	return nil
}
