package user

import (
	"fmt"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
	"go.uber.org/zap"
)

func (user *userDomainService) CreateUser(
	userDomain user.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
