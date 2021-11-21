package main

import (
	"fmt"
	"time"

	"github.com/RestartFU/slapper"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/query"
)

var SkinGanni, _ = slapper.DecodePNGSkin("./data/slapper/ganni.png", slapper.CustomSlimGeometry)
var SkinRestart, _ = slapper.DecodePNGSkin("./data/slapper/restart.png", slapper.CustomGeometry)

var NAPractice = slapper.New("Slapper(NA)", "§7Loading...", SkinRestart, mgl64.Vec3{-8.5, 144, -98.5}).WithAction(func(p *player.Player) {
	p.Transfer("nitrofaction.fr:19132")
}).WithSpawnFunc(naPractice).WithYawAndPitch(60, 0)

func naPractice(s *slapper.Slapper) {
	go func() {
		for {
			var newTag string
			q, err := query.Do("na.sgpractice.tk:19132")
			if err != nil {
				newTag = "§9NA Practice\n§cOFFLINE"
			} else {
				newTag = fmt.Sprintf("§9Practice\n§a%v/%v", q["numplayers"], q["maxplayers"])
			}
			s.SetNameTag(newTag)
			time.Sleep(3 * time.Second)
		}
	}()
}
