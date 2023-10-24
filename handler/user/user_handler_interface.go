package user

import "github.com/gin-gonic/gin"

type UserHandlerInterface interface {
	CreateUserHandler(context *gin.Context)
	DeleteUserHandler(context *gin.Context)
	FindUserByIDHandler(context *gin.Context)
	FindUserByEmailHandler(context *gin.Context)
	UpdateUserHandler(context *gin.Context)
}
