package gerador_campeonato

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/domain/gerador-campeonato"
	gerador_campeonato2 "github.com/Micheli97/crud-campeonato-golang/entity/gerador-campeonato"
	gerador_campeonato3 "github.com/Micheli97/crud-campeonato-golang/repository/convert/gerador-campeonato"
	"math/rand"
)

func (repository *geradorCampeonatoRepository) GerarCampeonatoHandler() *rest_err.RestErr {
	db := repository.database

	query := `SELECT COUNT(*) FROM t_team`

	var totalRegistro int

	result, err := db.Query(query)

	if err != nil {
		logger.Error("Ocorreu um erro ao criar time no banco de dados", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	for result.Next() {
		var totalResult int

		err = result.Scan(&totalResult)

		if err != nil {
			break
		}

		totalRegistro = totalResult

	}

	if totalRegistro < 8 {
		logger.Error("Ocorreu um erro ao gerar Campeonato", err)
		return rest_err.NewInternalServerError("É necessário ter 8 times cadastrados para gerar um campeonato.")

	}

	queryTotalPlayer := `SELECT COUNT(*) FROM t_player`

	var totalRegistroPlayer int32

	resultPlayer, err := db.Query(queryTotalPlayer)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	for resultPlayer.Next() {
		var totalResult int32

		err = resultPlayer.Scan(&totalResult)

		if err != nil {
			break
		}

		totalRegistroPlayer = totalResult

	}

	if totalRegistroPlayer < 176 {
		logger.Error("Ocorreu um erro ao gerar Campeonato", err)
		return rest_err.NewInternalServerError("Quantidade inválida de jogadores. Cada time deve conter 22 jogadores cadastrados.")

	}

	queryNameTeam := `SELECT name FROM t_team`

	resultNameTeam, err := db.Query(queryNameTeam)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	var teams []string

	for resultNameTeam.Next() {

		var team string

		err = resultPlayer.Scan(&team)

		if err != nil {
			break
		}

		teams = append(teams, team)
	}

	timesSorteados := geraCampeonato(teams)

	// Como adicionar no banco?

	return nil

}

func geraCampeonato(times []string) []string {
	var timesRegistrados []string

	timesDisponiveis := times

	for len(timesDisponiveis) > 0 && len(timesRegistrados) < 8 {
		indiceSorteado := rand.Int() * len(timesDisponiveis)

		timesRegistrados = append(timesRegistrados, timesRegistrados[indiceSorteado])
	}
	return timesRegistrados

}

func (repository *geradorCampeonatoRepository) ListarTimesCampeonatoHandler() ([]gerador_campeonato.GeradorCampeonatoDomainInterface, *rest_err.RestErr) {
	db := repository.database

	query := `SELECT * FROM t_campeonato`

	result, err := db.Query(query)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	var jogosCampeonato []gerador_campeonato.GeradorCampeonatoDomainInterface

	for result.Next() {

		var campeonato gerador_campeonato2.GeradorCampeonatoEntity

		err = result.Scan(&campeonato.TimeA, &campeonato.TimeB)

		if err != nil {
			break
		}

		teamConvert := gerador_campeonato3.ConvertGeradorCampeonatoToDomain(campeonato)

		jogosCampeonato = append(jogosCampeonato, teamConvert)
	}

	return jogosCampeonato, nil
}

func NewGeradorCampeonatoRepository(databse *sql.DB) GeradorCampeonatoRepositoryInterface {
	return &geradorCampeonatoRepository{
		database: databse,
	}
}
