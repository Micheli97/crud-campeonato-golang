package user

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/model/user"
	"go.uber.org/zap"
)

func (user *userDomainService) FindUserByEmailServices(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("Journey", "findUserByEmail"))
	return user.userRepository.FindUserByEmail(email)
}
func (user *userDomainService) FindUserByIDServices(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("Journey", "findUserByID"))
	return user.userRepository.FindUserByID(id)
}
