package team

import (
	"github.com/Micheli97/crud-campeonato-golang/handler/team"
	"github.com/gin-gonic/gin"
)

// TeamRouter rotas para team
func TeamRouter(router *gin.RouterGroup, team team.TeamHandlerInterface) {
	teamGroup := router.Group("/team")

	// Obter o total de times cadastrados
	teamGroup.GET("/teamCount", team.FindTotalTeams)

	// Listar os times cadastrados
	teamGroup.GET("/", team.FindTeams)

	// Cadastra um novo time
	teamGroup.POST("/createTeam", team.CreateTeamHandler)

	// Editar/Atualizar time
	teamGroup.PUT("/updateTeam/:id", team.UpdateTeamHandler)

	// Deleta um time
	teamGroup.DELETE("/deleteTeam/:id", team.DeleteTeamHandler)

	// Obtem um time pelo id
	teamGroup.GET("/:id", team.FindTeamByIDHandler)
}
