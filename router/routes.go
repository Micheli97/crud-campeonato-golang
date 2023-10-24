package router

import (
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/Micheli97/crud-campeonato-golang/router/user"
	"github.com/gin-gonic/gin"
)

// InitializeRouters inicializa as rotas
func InitializeRouters(router *gin.Engine, userHandler user2.UserHandlerInterface) {

	v1 := router.Group("api/v1")

	user.UserRouters(v1, userHandler)

}
