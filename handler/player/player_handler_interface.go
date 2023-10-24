package player

import "github.com/gin-gonic/gin"

type PlayerHandlerInterface interface {
	CreatePlayer(context *gin.Context)
	FindPlayers(context *gin.Context)
	FindTotalPlayer(context *gin.Context)
	FindPlayerByID(context *gin.Context)
	UpdatePlayer(context *gin.Context)
	DeletePlayer(context *gin.Context)
}
