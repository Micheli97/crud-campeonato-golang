package router

import (
	login2 "github.com/Micheli97/crud-campeonato-golang/handler/login"
	user2 "github.com/Micheli97/crud-campeonato-golang/handler/user"
	"github.com/Micheli97/crud-campeonato-golang/router/login"
	"github.com/Micheli97/crud-campeonato-golang/router/user"
	"github.com/gin-gonic/gin"
)

// InitializeRouters inicializa as rotas
func InitializeRouters(router *gin.Engine, userHandler user2.UserHandlerInterface, loginHandler login2.LoginHandlerInterface) {

	v1 := router.Group("api/v1")

	user.UserRouters(v1, userHandler)
	login.LoginRouter(v1, loginHandler)

}
