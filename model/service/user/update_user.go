package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
	"go.uber.org/zap"
)

func (user *userDomainService) UpdateUser(id string, userDomain user.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser handler",
		zap.String("journey", "createUser"))

	logger.Info(userDomain.GetPassword())
	err := user.userRepository.UpdateUser(id, userDomain)

	if err != nil {
		return err
	}

	return nil
}
