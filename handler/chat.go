package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player/chat"
)

func (h *PlayerHandler) HandleChat(ctx *event.Context, s *string) {
	ctx.Cancel()
	chat.Global.WriteString("§7" + h.P.Name() + ": " + *s)
}
