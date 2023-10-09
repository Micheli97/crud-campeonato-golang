package user

import (
	"github.com/Micheli97/crud-campeonato-golang/model/service/user"
	"github.com/gin-gonic/gin"
)

func NewUserHandlerInterface(serviceInterface user.UserDomainService) UserHandlerInterface {
	return &userHandlerInterface{
		service: serviceInterface,
	}
}

type UserHandlerInterface interface {
	CreateUserHandler(context *gin.Context)
	DeleteUserHandler(context *gin.Context)
	FindUserByIDHandler(context *gin.Context)
	FindUserByEmailHandler(context *gin.Context)
	UpdateUserHandler(context *gin.Context)
}

type userHandlerInterface struct {
	service user.UserDomainService
}
