package team

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/team"
)

type TeamRepositoryInterface interface {
	CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr)

	FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr)

	UpdateTeam(id string, team team.TeamDomainInterface) *rest_err.RestErr
	DeleteTeam(id string) *rest_err.RestErr

	FindTotalRegisteredTeams() (int, *rest_err.RestErr)

	FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr)
}
