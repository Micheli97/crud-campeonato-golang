package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"go.uber.org/zap"
)

func (user *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))
	user.EncryptPassword()
	return nil
}
