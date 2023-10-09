package user

import (
	userHandler "github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/gin-gonic/gin"
)

// UserRouters rotas usuário
func UserRouters(router *gin.RouterGroup, userHandler userHandler.UserHandlerInterface) {

	routerGroup := router.Group("/user")

	routerGroup.GET("/getUserById/:userId", userHandler.FindUserByIDHandler)

	routerGroup.GET("/getUserByEmail/:userEmail", userHandler.FindUserByEmailHandler)

	routerGroup.POST("/createUser/", userHandler.CreateUserHandler)

	routerGroup.PUT("/updateUser/:userId", userHandler.UpdateUserHandler)

	routerGroup.DELETE("/:userId", userHandler.DeleteUserHandler)

}
