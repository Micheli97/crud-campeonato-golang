package team

import (
	team2 "github.com/Micheli97/crud-campeonato-golang/domain/team"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
	"github.com/Micheli97/crud-campeonato-golang/service/team"
	"github.com/Micheli97/crud-campeonato-golang/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewTeamHandler(teamService team.TeamServiceInterface) TeamHandlerInterface {
	return &teamHandler{teamService: teamService}
}

// UpdateTeamHandler atualiza o time pelo id
func (team *teamHandler) UpdateTeamHandler(context *gin.Context) {

	var teamRequest request.TeamRequest

	teamID := context.Param("id")

	if err := context.ShouldBindJSON(&teamRequest); err != nil {

		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar dados do time")
		return
	}

	teamDomain := team2.NewTeamDomain(
		teamRequest.Name,
		teamRequest.Coach,
		teamRequest.Website,
		teamRequest.City,
		teamRequest.BadgePhoto,
	)

	err := team.teamService.UpdateTeam(teamID, teamDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar time.")
		return
	}

	context.JSON(http.StatusOK, "Time atualizado com sucesso!")

}

// CreateTeamHandler cadastra um novo time
func (team *teamHandler) CreateTeamHandler(context *gin.Context) {

	var teamRequest request.TeamRequest

	var totalRegisteredTeams = 8

	if err := context.ShouldBindJSON(&teamRequest); err != nil {

		context.JSON(http.StatusBadRequest, "Erro ao tentar criar time")
		return
	}

	totalTeams, err := team.teamService.FindTotalRegisteredTeams()

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar time.")
		return
	}

	if totalTeams > totalRegisteredTeams {
		context.JSON(http.StatusBadRequest, "O limite total de cadastro de times foi atingido.")
		return

	}

	teamDomain := team2.NewTeamDomain(
		teamRequest.Name,
		teamRequest.Coach,
		teamRequest.Website,
		teamRequest.City,
		teamRequest.BadgePhoto,
	)

	teamResponse, err := team.teamService.CreateTeam(teamDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar atualizar time.")
		return
	}

	context.JSON(http.StatusOK, view.ConvertTeamDomainToResponse(teamResponse))
}

// DeleteTeamHandler remove o time
func (team *teamHandler) DeleteTeamHandler(context *gin.Context) {

	teamID := context.Param("id")

	if err := team.teamService.DeleteTeam(teamID); err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar deletar time.")
		return

	}

	context.JSON(http.StatusOK, "Time excluído com sucesso!")
}

// GetTeamHandler obtém um time pelo id
func (team *teamHandler) FindTeamByIDHandler(context *gin.Context) {

	teamID := context.Param("id")

	teamResponse, err := team.teamService.FindTeamByID(teamID)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar obter time.")
		return
	}

	context.JSON(http.StatusOK, view.ConvertTeamDomainToResponse(teamResponse))
}

// GetTeamsHandler retorna a lista de times cadastrados
func (team *teamHandler) FindTeams(context *gin.Context) {

	teamResponse, err := team.teamService.FindTotalTeams()

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar obter times.")
		return
	}

	context.JSON(http.StatusOK, teamResponse)

}

// GetTotalRegisteredTeamsHandler obtem a quantidade total de times cadastrados
func (team *teamHandler) FindTotalTeams(context *gin.Context) {

	totalTeams, err := team.teamService.FindTotalRegisteredTeams()

	if err != nil {
		context.JSON(http.StatusBadRequest, "Erro ao tentar obter total de times.")
		return
	}

	context.JSON(http.StatusOK, totalTeams)
}
