package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
	"github.com/Micheli97/crud-campeonato-golang/model/utils"
	"go.uber.org/zap"
)

func (user *userDomainService) CreateUser(
	userDomain user.UserDomainInterface,
) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser handler",
		zap.String("journey", "createUser"))
	utils.EncryptPassword(userDomain)

	logger.Info(userDomain.GetPassword())

	userDomainRepository, err := user.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil

}
