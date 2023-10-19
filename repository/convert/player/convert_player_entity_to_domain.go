package player

import (
	"github.com/Micheli97/crud-campeonato-golang/domain/player"
	player2 "github.com/Micheli97/crud-campeonato-golang/entity/player"
)

func ConverPlayerEntityToDomain(playerEntity player2.PlayerEntity) player.PlayerDomainInterface {

	domain := player.NewPlayerDomain(
		playerEntity.Name,
		playerEntity.Photo,
		playerEntity.Position,
		playerEntity.TeamID,
		playerEntity.Number,
		playerEntity.Age,
		playerEntity.Weight,
		playerEntity.Height,
	)

	domain.SetID(playerEntity.Id)

	return domain
}
