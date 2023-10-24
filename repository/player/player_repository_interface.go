package player

import (
	"database/sql"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/player"
	player3 "github.com/Micheli97/crud-campeonato-golang/entity/player"
	player2 "github.com/Micheli97/crud-campeonato-golang/repository/convert/player"
)

func PlayerRepository(database *sql.DB) PlayerRepositoryInterface {
	return &playerRepository{database: database}
}

func (player *playerRepository) CreatePlayer(playerDomain player.PlayerDomainInterface) (player.PlayerDomainInterface, *rest_err.RestErr) {
	db := player.database

	value := player2.ConvertPlayerDomainToEntity(playerDomain)

	query := `INSERT INTO t_player(name, photo, height, weight, age, position, number, teamID)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := db.QueryRow(query, value.Name, value.Photo, value.Height, value.Weight, value.Age, value.Position, value.Number, value.TeamID).Scan(&value.Id)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return player2.ConverPlayerEntityToDomain(*value), nil
}

func (player *playerRepository) FindPlayers() ([]player.PlayerDomainInterface, *rest_err.RestErr) {
	//db := player.database
	//
	//var players []player.PlayerDomainInterface
	//
	//query := `SELECT * FROM t_player`
	//
	//rows, err := db.Query(query)
	//
	//if err != nil {
	//	return nil, rest_err.NewInternalServerError(err.Error())
	//}
	//
	//for rows.Next() {
	//	var p player3.PlayerEntity
	//
	//	err := rows.Scan(&p.Age, &p.Name, &p.Height, &p.Id, &p.Position, &p.Weight, &p.Number, &p.TeamID, &p.Photo)
	//
	//	if err != nil {
	//		break
	//	}
	//
	//	players = append(players, p)
	//
	//}
	//
	//return players, nil

	return nil, nil

}

func (player *playerRepository) FindTotalPlayer() (int32, *rest_err.RestErr) {
	db := player.database

	query := `SELECT COUNT(*) FROM t_player`

	var totalRegistro int32

	result, err := db.Query(query)

	if err != nil {
		return 0, rest_err.NewInternalServerError(err.Error())
	}

	for result.Next() {
		var totalResult int32

		err = result.Scan(&totalResult)

		if err != nil {
			break
		}

		totalRegistro = totalResult

	}

	return totalRegistro, nil

}

func (player *playerRepository) FindPlayerByID(id string) (player.PlayerDomainInterface, *rest_err.RestErr) {
	db := player.database

	query := `SELECT * FROM t_player WHERE id=$1`

	playerEntity := &player3.PlayerEntity{}

	err := db.QueryRow(query, id).Scan(&playerEntity.Id, &playerEntity.Name, &playerEntity.Photo, &playerEntity.Height,
		&playerEntity.Weight, &playerEntity.Age, &playerEntity.Position, &playerEntity.Number, &playerEntity.TeamID)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return player2.ConverPlayerEntityToDomain(*playerEntity), nil
}

func (player *playerRepository) UpdatePlayer(id string, playerDomain player.PlayerDomainInterface) *rest_err.RestErr {
	return nil
}

func (player *playerRepository) DeletePlayer(id string) *rest_err.RestErr {
	db := player.database

	query := `DELETE FROM t_player WHERE id=$1`

	_, err := db.Exec(query, id)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
