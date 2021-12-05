package main

import (
	"fmt"
	"time"

	"github.com/RestartFU/mcbe"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/dragonfly-on-steroids/npc"
	"github.com/go-gl/mathgl/mgl64"
)

func NewSlapperTransfer(pos mgl64.Vec3, skin skin.Skin, addr, name string) *npc.NPC {
	f := func(s *npc.NPC) {
		ticker := time.NewTicker(5 * time.Second)
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

	s := npc.New("Slapper("+name+")", "§7Loading...", SkinRestart, mgl64.Vec3{-8.5, 144, -98.5}).WithAction(func(p *player.Player) {
		p.Transfer(addr)
	}).WithSpawnFunc(f)

	return s
}

var SkinGanni, _ = npc.DecodePNGSkin("./data/slapper/ganni.png", npc.CustomSlimGeometry)
var SkinRestart, _ = npc.DecodePNGSkin("./data/slapper/restart.png", npc.CustomGeometry)

var NAPractice = NewSlapperTransfer(mgl64.Vec3{-8.5, 144, -98.5}, SkinRestart, "glowhcf.net:19134", "§9HCF").WithYawAndPitch(60, 0)
