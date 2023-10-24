package team

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	"github.com/Micheli97/crud-campeonato-golang/domain/team"
	team3 "github.com/Micheli97/crud-campeonato-golang/entity/team"
	team2 "github.com/Micheli97/crud-campeonato-golang/repository/convert/team"
)

func TeamRepository(database *sql.DB) TeamRepositoryInterface {
	return &teamRepository{
		databaseConnection: database,
	}
}

func (t *teamRepository) CreateTeam(team team.TeamDomainInterface) (team.TeamDomainInterface, *rest_err.RestErr) {

	db := t.databaseConnection

	value := team2.ConvertDomainToEntity(team)

	query := `INSERT INTO t_team(name, badgePhoto, city, coach, website)
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(query, value.Name, value.BadgePhoto, value.City, value.Coach, value.Website).Scan(&value.ID)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return team2.ConvertTeamEntityToDomain(*value), nil
}

func (t *teamRepository) FindTeamByID(id string) (team.TeamDomainInterface, *rest_err.RestErr) {
	db := t.databaseConnection

	query := `SELECT * from t_team where id=$1`

	teamEntity := &team3.TeamEntity{}

	err := db.QueryRow(query, id).Scan(&teamEntity.ID, &teamEntity.Name, &teamEntity.City, &teamEntity.Coach, &teamEntity.Website, &teamEntity.BadgePhoto)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return team2.ConvertTeamEntityToDomain(*teamEntity), nil
}

func (t *teamRepository) UpdateTeam(id string, team team.TeamDomainInterface) *rest_err.RestErr {
	db := t.databaseConnection

	query := `UPDATE t_team SET name=$1, badgePhoto=$2, city=$3, coach=$4, website=$5 where id=$6 `

	_, err := db.Exec(query, team.GetName(), team.GetBadgePhoto(), team.GetCity(), team.GetCoach(), team.GetWebSite(), id)

	if err != nil {
		logger.Error("Ocorreu um erro ao atualizar time no banco de dados", err)
		return rest_err.NewInternalServerError(err.Error())
	}
	return nil

}

func (t *teamRepository) DeleteTeam(id string) *rest_err.RestErr {
	db := t.databaseConnection

	query := `DELETE FROM t_team where id=$1`

	_, err := db.Exec(query, id)

	if err != nil {
		logger.Error("Ocorreu um erro ao Deletar time no banco de dados", err)

		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}

func (t *teamRepository) FindTotalRegisteredTeams() (int, *rest_err.RestErr) {
	db := t.databaseConnection

	query := `SELECT COUNT(*) FROM t_team`

	var totalRegistro int

	result, err := db.Query(query)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return 0, rest_err.NewInternalServerError(err.Error())
	}

	for result.Next() {
		var totalResult int

		err = result.Scan(&totalResult)

		if err != nil {
			break
		}

		totalRegistro = totalResult

	}

	return totalRegistro, nil
}

func (t *teamRepository) FindTotalTeams() ([]team.TeamDomainInterface, *rest_err.RestErr) {
	db := t.databaseConnection

	query := `SELECT * FROM t_team`

	var teams []team.TeamDomainInterface

	result, err := db.Query(query)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	for result.Next() {

		var team team3.TeamEntity
		err = result.Scan(&team.ID, &team.Name, &team.City, &team.Coach, &team.BadgePhoto, &team.Website)

		if err != nil {
			break
		}

		teamConvert := team2.ConvertTeamEntityToDomain(team)

		teams = append(teams, teamConvert)
	}

	return teams, nil

}
