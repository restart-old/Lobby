package handler

import (
	"math"

	"github.com/df-mc/dragonfly/server/event"
	"github.com/go-gl/mathgl/mgl64"
)

func (h *PlayerHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	p := h.P
	if p.Name() == "RestartFU" {
		p.SendTip("pos:", math.Round(newPos.X()), math.Round(newPos.Y()), math.Round(newPos.Z()), "yaw:", math.Round(newYaw), "pitch:", math.Round(newPitch))
	}
}
