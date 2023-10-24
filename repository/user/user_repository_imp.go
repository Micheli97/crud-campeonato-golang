package user

import (
	"database/sql"
	"fmt"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	user3 "github.com/Micheli97/crud-campeonato-golang/domain/user"
	user4 "github.com/Micheli97/crud-campeonato-golang/entity/user"
	user2 "github.com/Micheli97/crud-campeonato-golang/repository/convert/user"
	"strconv"
	"strings"
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
	consulta, valores := gerarConsultaAtualizacao(userDomain, id)

	_, err := db.Exec(consulta, valores...)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}

func gerarConsultaAtualizacao(user user3.UserDomainInterface, idUser string) (string, []interface{}) {
	var filtros = map[string]string{
		"name":  user.GetName(),
		"email": user.GetEmail(),
	}

	campos := make([]string, 0)
	valores := make([]interface{}, 0)

	for filtro, valor := range filtros {
		if valor != "" {
			campos = append(campos, filtro+" = $"+strconv.Itoa(len(valores)+1))
			valores = append(valores, valor)
		}
	}

	consulta := fmt.Sprintf("UPDATE t_user SET %s WHERE id = $%d", strings.Join(campos, ", "), len(valores)+1)

	id, err := strconv.ParseInt(idUser, 10, 32)
	if err != nil {
		logger.Error("Erro ao converter valor", err)
		return "", nil
	}
	valores = append(valores, int32(id))

	return consulta, valores
}

func NewUserRespository(database *sql.DB) UserRepository {
	return &userRespository{
		database,
	}
}
