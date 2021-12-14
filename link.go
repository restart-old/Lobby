package main

import (
	"time"

	"github.com/SGPractice/link"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/hako/durafmt"
)

type LinkCommand struct {
	linker *link.Linker
}

func (l LinkCommand) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		var code link.Code
		if r, ok := l.linker.LinkedFromGamerTag(src.Name()); !ok {
			if code, ok = l.linker.LoadByUser(src.Name()); !ok {
				code = link.NewCode(7)
				l.linker.Storer.Store(src.Name(), code)
			}
			until := time.Until(code.Expiration)
			p.Messagef("Your code is %s and it will expire in %s.", code.Code, durafmt.Parse(until).LimitFirstN(2))
		} else {
			p.Messagef("You are already linked with the ID: %s, use /unlink if you wish to link again.", r.DiscordID())
		}
	}
}

type UnlinkCommand struct {
	linker *link.Linker
}

func (l UnlinkCommand) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if r, ok := l.linker.LinkedFromGamerTag(src.Name()); ok {
			if r.LinkedSince().Add(21600 * time.Minute).Before(time.Now()) {
				l.linker.UnLink(src.Name())
				p.Message("You have been unlinked from the ID: %s.", r.DiscordID())
			} else {
				p.Messagef("You are on unlink cooldown, wait %s.", durafmt.Parse(time.Until(r.LinkedSince().Add(21600*time.Minute))).LimitFirstN(1))
			}
		}
	}
}