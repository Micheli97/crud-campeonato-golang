package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user3 "github.com/Micheli97/crud-campeonato-golang/model/repository/entity/convert/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/model/user"
)

func (user *userRespository) CreateUser(userDomain user2.UserDomainInterface) (user2.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	db := user.databaseConnection

	value := user3.ConvertDomainToEntity(userDomain)

	query := `INSERT INTO user(name, email, password)
	VALUES(%s, %s, %s)`

	_, err := db.Query(query, value.Name, value.Email, value.Password)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar cadastrar usuário no banco de dados", err)
	}

	//rows, err := db.Query(`SELECT * from TABLE ( user )`)
	//
	//for rows.Next() {
	//	var id string
	//	err = rows.Scan(&id)
	//
	//	if err != nil {
	//		break
	//	}
	//
	//	value.ID = id
	//
	//}
	//
	//return user3.ConvertEntityToDomain(*value), nil

	return nil, nil

}
