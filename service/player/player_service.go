package player

import "github.com/Micheli97/crud-campeonato-golang/repository/player"

type playerService struct {
	playerRepository player.PlayerRepositoryInterface
}
