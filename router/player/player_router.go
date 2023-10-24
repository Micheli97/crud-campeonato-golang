package player

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/player"
	"github.com/gin-gonic/gin"
)

func PLayerRouter(router *gin.RouterGroup, playerHandler player.PlayerHandlerInterface) {

	playerGroup := router.Group("/player")

	// Listar jogadores cadastrados
	playerGroup.GET("/", playerHandler.FindPlayers)

	// Lista a quantidade de jogadores cadastrados
	playerGroup.GET("/countPlayer", playerHandler.FindTotalPlayer)

	// Cadastra um novo jogador
	playerGroup.POST("/", playerHandler.CreatePlayer)

	// Atualiza os dados do jogador
	playerGroup.PUT("/:id", playerHandler.UpdatePlayer)

	// Deleta jogador
	playerGroup.DELETE("/:id", playerHandler.DeletePlayer)

	// Lista jogador pelo ID
	playerGroup.GET("/:id", playerHandler.FindPlayerByID)

}
