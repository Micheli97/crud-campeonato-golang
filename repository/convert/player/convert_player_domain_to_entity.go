package player

import (
	"github.com/Micheli97/crud-campeonato-golang/domain/player"
	player2 "github.com/Micheli97/crud-campeonato-golang/entity/player"
)

func ConvertPlayerDomainToEntity(playerDomain player.PlayerDomainInterface) *player2.PlayerEntity {
	return &player2.PlayerEntity{
		Name:     playerDomain.GetName(),
		Photo:    playerDomain.GetPhoto(),
		Height:   playerDomain.GetHeight(),
		Weight:   playerDomain.GetWeight(),
		Age:      playerDomain.GetAge(),
		Position: playerDomain.GetPosition(),
		Number:   playerDomain.GetNumber(),
		TeamID:   playerDomain.GetTeamID(),
	}
}
