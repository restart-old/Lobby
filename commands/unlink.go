package commands

import (
	"time"

	"github.com/SGPractice/link"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/hako/durafmt"
)

var Linker *link.Linker

type UnlinkCommand struct {
}

func (l UnlinkCommand) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if r, ok, err := Linker.LinkedFromXUID(p.XUID()); ok {
			if err != nil {
				p.Message("§cAn error occurred, please notify staff!")
				return
			}
			if r.LinkedSince().Add(21600 * time.Minute).Before(time.Now()) {
				Linker.UnLink(src.Name())
				p.Message("§aYou have been unlinked from the ID: %s.", r.DiscordID())
			} else {
				p.Messagef("§cYou are on unlink cooldown, wait %s.", durafmt.Parse(time.Until(r.LinkedSince().Add(21600*time.Minute))).LimitFirstN(1))
			}
		} else {
			p.Message("§cYou are not linked to any discord account.")
		}
	}
}
