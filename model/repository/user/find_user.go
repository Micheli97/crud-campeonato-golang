package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user3 "github.com/Micheli97/crud-campeonato-golang/model/repository/convert/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/repository/entity/user"

	"github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userRespository) FindUserByEmail(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	db := user.databaseConnection

	query := `SELECT id, name, email FROM t_user WHERE email=$1`

	userEntity := &user2.UserEntity{}

	err := db.QueryRow(query, email).Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return nil, nil
	}

	return user3.ConvertEntityToDomain(*userEntity), nil
}

func (user *userRespository) FindUserByID(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	db := user.databaseConnection

	query := `SELECT id, name, email FROM t_user WHERE id=$1`

	userEntity := &user2.UserEntity{}

	err := db.QueryRow(query, id).Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return nil, nil
	}

	return user3.ConvertEntityToDomain(*userEntity), nil
}
