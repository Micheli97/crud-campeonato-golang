package player

import (
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/Micheli97/crud-campeonato-golang/domain/player"
)

type PlayerRepositoryInterface interface {
	CreatePlayer(playerDomain player.PlayerDomainInterface) (player.PlayerDomainInterface, *rest_err.RestErr)

	FindPlayers() ([]player.PlayerDomainInterface, *rest_err.RestErr)

	FindTotalPlayer() (int32, *rest_err.RestErr)

	FindPlayerByID(id string) (player.PlayerDomainInterface, *rest_err.RestErr)

	UpdatePlayer(id string, playerDomain player.PlayerDomainInterface) *rest_err.RestErr

	DeletePlayer(id string) *rest_err.RestErr
}
