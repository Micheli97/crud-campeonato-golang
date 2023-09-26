package user

import "github.com/gin-gonic/gin"

// UserRouters rotas usuário
func UserRouters(router *gin.RouterGroup) {

	routerGroup := router.Group("/user")

	routerGroup.GET("/")

	routerGroup.POST("/")

}
