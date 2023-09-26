package user

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/gin-gonic/gin"
)

// UserRouters rotas usuário
func UserRouters(router *gin.RouterGroup) {

	routerGroup := router.Group("/user")

	routerGroup.GET("/", user.GetUserHandler)

	routerGroup.POST("/", user.CreateUserHandler)

}
