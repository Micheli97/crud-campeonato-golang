package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/repository/convert/user"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userRespository) CreateUser(userDomain user.UserDomainInterface) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	db := user.databaseConnection

	value := user2.ConvertDomainToEntity(userDomain)

	query := `INSERT INTO t_user(name, email, password)
	VALUES($1, $2, $3) RETURNING id`

	err := db.QueryRow(query, value.Name, value.Email, value.Password).Scan(&value.ID)
	if err != nil {
		logger.Error("Ocorreu um erro ao tentar cadastrar usuário no banco de dados", err)
	}

	return user2.ConvertEntityToDomain(*value), nil
}
