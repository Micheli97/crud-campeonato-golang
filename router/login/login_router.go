package login

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/login"
	"github.com/gin-gonic/gin"
)

func LoginRouter(router *gin.RouterGroup, login login.LoginHandlerInterface) {

	loginGroup := router.Group("/login")

	loginGroup.POST("/", login.LoginHandler)

}
