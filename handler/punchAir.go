package handler

import "github.com/df-mc/dragonfly/server/event"

func (p *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	ctx.Cancel()
}
