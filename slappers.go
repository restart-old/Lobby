package main

import (
	"fmt"
	"time"

	"github.com/RestartFU/mcbe"
	"github.com/RestartFU/slapper"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/go-gl/mathgl/mgl64"
)

func NewSlapperTransfer(pos mgl64.Vec3, skin skin.Skin, addr, name string) *slapper.Slapper {
	f := func(s *slapper.Slapper) {
		ticker := time.NewTicker(3 * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					var newTag string
					q, err := mcbe.Query(addr)
					if err != nil {
						newTag = name + "\n§cOFFLINE"
					} else {
						newTag = fmt.Sprintf(name+"\n§a%v/%v", q["online_players"], q["max_players"])
					}
					s.SetNameTag(newTag)
				}
			}
		}()
	}

	s := slapper.New("Slapper("+name+")", "§7Loading...", SkinRestart, mgl64.Vec3{-8.5, 144, -98.5}).WithAction(func(p *player.Player) {
		p.Transfer(addr)
	}).WithSpawnFunc(f)

	return s
}

var SkinGanni, _ = slapper.DecodePNGSkin("./data/slapper/ganni.png", slapper.CustomSlimGeometry)
var SkinRestart, _ = slapper.DecodePNGSkin("./data/slapper/restart.png", slapper.CustomGeometry)

var NAPractice = NewSlapperTransfer(mgl64.Vec3{-8.5, 144, -98.5}, SkinRestart, "glowhcf.net:19134", "§9HCF").WithYawAndPitch(60, 0)
