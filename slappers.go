package main

import (
	"github.com/RestartFU/slapper"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

var SkinGanni, _ = slapper.DecodePNGSkin("./data/slapper/ganni.png", slapper.CustomSlimGeometry)
var SkinRestart, _ = slapper.DecodePNGSkin("./data/slapper/restart.png", slapper.CustomGeometry)

var SlapperPractice = slapper.New("Slapper(Practice)", "§7Loading...", SkinGanni, mgl64.Vec3{0, 10, 0}, 0, 0).WithAction(func(p *player.Player) {
})
