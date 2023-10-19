package team

import (
	team2 "github.com/Micheli97/crud-campeonato-golang/domain/team"
	"github.com/Micheli97/crud-campeonato-golang/entity/team"
)

func ConvertTeamEntityToDomain(teamEntity team.TeamEntity) team2.TeamDomainInterface {
	domain := team2.NewTeamDomain(
		teamEntity.Name,
		teamEntity.BadgePhoto,
		teamEntity.Coach,
		teamEntity.City,
		teamEntity.Website,
	)

	domain.SetID(teamEntity.ID)
	return domain
}
