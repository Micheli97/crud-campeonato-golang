package gerador_campeonato

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/config/logger"
	gerador_campeonato "github.com/Micheli97/crud-campeonato-golang/domain/gerador-campeonato"
	gerador_campeonato2 "github.com/Micheli97/crud-campeonato-golang/entity/gerador-campeonato"
	gerador_campeonato3 "github.com/Micheli97/crud-campeonato-golang/repository/convert/gerador-campeonato"
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

	queryNamePlayers := `SELECT name FROM t_player`

	resultPlayerName, err := db.Query(queryNamePlayers)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	//  Crian função para gerar o campeonato

	return nil

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
