package team

import (
	team2 "github.com/Micheli97/crud-campeonato-golang/domain/team"
	"github.com/Micheli97/crud-campeonato-golang/entity/team"
)

func ConvertDomainToEntity(teamDomain team2.TeamDomainInterface) *team.TeamEntity {
	return &team.TeamEntity{
		ID:         teamDomain.GetID(),
		Name:       teamDomain.GetName(),
		Website:    teamDomain.GetWebSite(),
		City:       teamDomain.GetCity(),
		Coach:      teamDomain.GetCoach(),
		BadgePhoto: teamDomain.GetBadgePhoto(),
	}
}
