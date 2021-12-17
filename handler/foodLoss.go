package handler

import "github.com/df-mc/dragonfly/server/event"

func (*PlayerHandler) HandleFoodLoss(ctx *event.Context, from, to int) {
	ctx.Cancel()
}
