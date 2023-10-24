package view

import (
	"github.com/Micheli97/crud-campeonato-golang/domain/team"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/response"
)

func ConvertTeamDomainToResponse(team team.TeamDomainInterface) response.TeamResponse {
	return response.TeamResponse{
		Coach:      team.GetCoach(),
		Name:       team.GetName(),
		BadgePhoto: team.GetBadgePhoto(),
		City:       team.GetCity(),
		Website:    team.GetWebSite(),
		ID:         team.GetID(),
	}
}
