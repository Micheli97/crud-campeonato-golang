package view

import (
	"github.com/Micheli97/crud-campeonato-golang/domain/player"
	"github.com/Micheli97/crud-campeonato-golang/handler/model/request"
)

func ConvertPlayerDomainToResponse(playerDomain player.PlayerDomainInterface) request.PlayerRequest {
	return request.PlayerRequest{
		Id:       playerDomain.GetID(),
		Weight:   playerDomain.GetWeight(),
		Name:     playerDomain.GetName(),
		Height:   playerDomain.GetHeight(),
		Number:   playerDomain.GetNumber(),
		Photo:    playerDomain.GetPhoto(),
		Position: playerDomain.GetPosition(),
		Age:      playerDomain.GetAge(),
		TeamID:   playerDomain.GetTeamID(),
	}
}
