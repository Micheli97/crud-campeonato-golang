package user

import (
	userHandler "github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/gin-gonic/gin"
)

// UserRouters rotas usuário
func UserRouters(router *gin.RouterGroup, userHandler userHandler.UserHandlerInterface) {

	routerGroup := router.Group("/user")

	routerGroup.GET("/", userHandler.GetUserHandler)

	routerGroup.POST("/", userHandler.CreateUserHandler)

	routerGroup.PUT("/", userHandler.UpdateUserHandler)

	routerGroup.DELETE("/", userHandler.DeleteUserHandler)

}
