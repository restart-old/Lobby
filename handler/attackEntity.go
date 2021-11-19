package handler

import (
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

func (p *PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64) {
	ctx.Cancel()
	if v, ok := e.(*player.Player); ok && v.XUID() == "" {
		v.Hurt(1, damage.SourceEntityAttack{Attacker: p.P})
	}

}
