package user

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user3 "github.com/Micheli97/crud-campeonato-golang/domain/user"
	user4 "github.com/Micheli97/crud-campeonato-golang/entity/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/repository/convert/user"
)

func (user *userRespository) CreateUser(userDomain user3.UserDomainInterface) (user3.UserDomainInterface, *rest_err.RestErr) {
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
func (user *userRespository) DeleteUser(id string) *rest_err.RestErr {
	db := user.databaseConnection

	query := `DELETE FROM t_user WHERE id=$1`

	_, err := db.Exec(query, id)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar cadastrar usuário no banco de dados", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil

}

func (user *userRespository) FindUserByEmail(email string) (user3.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	db := user.databaseConnection

	query := `SELECT id, name, email FROM t_user WHERE email=$1`

	userEntity := &user4.UserEntity{}

	err := db.QueryRow(query, email).Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return user2.ConvertEntityToDomain(*userEntity), nil
}

func (user *userRespository) FindUserByID(id string) (user3.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	db := user.databaseConnection

	query := `SELECT id, name, email FROM t_user WHERE id=$1`

	userEntity := &user4.UserEntity{}

	err := db.QueryRow(query, id).Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return user2.ConvertEntityToDomain(*userEntity), nil
}

func (user *userRespository) UpdateUser(id string, userDomain user3.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository")

	db := user.databaseConnection

	query := `UPDATE  t_user SET name=$1, email=$2 WHERE id=$3`

	_, err := db.Exec(query, userDomain.GetName(), userDomain.GetEmail(), id)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}

func NewUserRespository(database *sql.DB) UserRepository {
	return &userRespository{
		database,
	}
}
