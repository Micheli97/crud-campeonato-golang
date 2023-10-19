package player

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewPlayerHandler(service player.PlayerServiceInterface) PlayerHandlerInterface {
	return &playerHandler{playerService: service}
}

func (player *playerHandler) CreatePlayer(context *gin.Context) {
	var playerRequest request.PlayerRequest

	if err := context.ShouldBindJSON(&playerRequest); err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao cadastrar jogador. Erro interno!")
		return
	}

	playerDomain := player2.NewPlayerDomain(
		playerRequest.Name,
		playerRequest.Photo,
		playerRequest.Position,
		playerRequest.TeamID,
		playerRequest.Number,
		playerRequest.Age,
		playerRequest.Weight,
		playerRequest.Height,
	)

	result, err := player.playerService.CreatePlayer(playerDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao cadastrar jogador. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, view.ConvertPlayerDomainToResponse(result))
}

func (player *playerHandler) FindPlayers(context *gin.Context) {
	result, err := player.playerService.FindPlayers()

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao buscar jogadores. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, result)
}

func (player *playerHandler) FindTotalPlayer(context *gin.Context) {
	result, err := player.playerService.FindTotalPlayer()

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao buscar total de jogadores. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, result)

}

func (player *playerHandler) FindPlayerByID(context *gin.Context) {
	idPlayer := context.Param("id")

	result, err := player.playerService.FindPlayerByID(idPlayer)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao buscar total de jogadores. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, view.ConvertPlayerDomainToResponse(result))

}

func (player *playerHandler) UpdatePlayer(context *gin.Context) {
	idPlayer := context.Param("id")

	var playerRequest request.PlayerRequest

	if err := context.ShouldBindJSON(&playerRequest); err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao atualizar dados do jogador. Erro interno!")
		return
	}

	domain := player2.NewPlayerDomain(
		playerRequest.Name,
		playerRequest.Photo,
		playerRequest.Position,
		playerRequest.TeamID,
		playerRequest.Number,
		playerRequest.Age,
		playerRequest.Weight,
		playerRequest.Height,
	)

	err := player.playerService.UpdatePlayer(idPlayer, domain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao atualizar dados do jogador. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, "Dados do jogador atualizados com sucesso!")
}

func (player *playerHandler) DeletePlayer(context *gin.Context) {
	idPlayer := context.Param("id")

	err := player.playerService.DeletePlayer(idPlayer)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao atualizar dados do jogador. Erro interno!")
		return
	}

	context.JSON(http.StatusOK, "Jogador excluído com sucesso!")

}
