package handler

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

var used = map[*player.Player]bool{}

func (p *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	if !p.P.OnGround() && !used[p.P] {
		p.P.PlaySound(sound.Pop{})
		v := entity.DirectionVector(p.P)
		newV := mgl64.Vec3{v.X(), 0.65, v.Z()}
		p.P.SetVelocity(newV)
		used[p.P] = true
		time.AfterFunc(1*time.Second/2, func() {
			delete(used, p.P)
		})
	}
	ctx.Cancel()
}
