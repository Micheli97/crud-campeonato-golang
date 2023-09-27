package repository

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userRespository) CreateUser(userDomain user2.UserDomainInterface) (user2.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	db := user.databaseConnection

	query := `INSERT INTO user(name, email, password)
	VALUES(%s, %s, %s)`

	rows, err := db.Query(query, userDomain.GetName(), userDomain.GetEmail(), userDomain.GetPassword())

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar cadastrar usuário no banco de dados", err)
	}

	var userID string

	for rows.Next() {
		var id string
		err = rows.Scan(&id)

		if err != nil {
			break
		}

		userID = id

	}

	userDomain.SetID(userID)

	return nil, nil
}
