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

var SlapperPractice = slapper.New("Slapper(Practice)", "§7Loading...", SkinGanni, mgl64.Vec3{-8, 144, -99}).WithAction(func(p *player.Player) {
	p.Transfer("nitrofaction.fr:19132")
}).WithSpawnFunc(practiceSpawnFunc)

func practiceSpawnFunc(s *slapper.Slapper) {
	go func() {
		for {
			var newTag string
			q, err := query.Do("nitrofaction.fr:19132")
			if err != nil {
				fmt.Println(err)
				newTag = "§9Practice\n\uE300\n§cOFFLINE"
			} else {
				newTag = fmt.Sprintf("§9Practice\n\uE300\n§a%v/%v", q["numplayers"], q["maxplayers"])
			}
			s.SetNameTag(newTag)
			time.Sleep(3 * time.Second)
		}
	}()
}
