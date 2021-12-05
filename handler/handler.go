package handler

import "github.com/df-mc/dragonfly/server/player"

type PlayerHandler struct {
	player.NopHandler
	P *player.Player
}

func (*PlayerHandler) Name() string {
	return "PlayerHandler"
}
