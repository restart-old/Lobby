package commands

import (
	"time"

	"github.com/SGPractice/link"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/hako/durafmt"
)

type LinkCommand struct {
}

func (l LinkCommand) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		var code link.Code
		if r, ok, err := Linker.LinkedFromXUID(p.XUID()); !ok {
			if err != nil {
				p.Message("§cAn error occurred, please notify staff!")
				return
			}
			if code, _, ok = Linker.LoadByUser(src.Name()); !ok {
				code = link.NewCode(7, p.XUID())
				Linker.Storer.Store(src.Name(), code)
			}
			until := time.Until(code.Expiration)
			p.Messagef("§aYour code is %s and it will expire in %s.\ngo on discord.gg/glowhcf and run the command /link <code>", code.Code, durafmt.Parse(until).LimitFirstN(2))
		} else {
			p.Messagef("§cYou are already linked with the ID: %s, use /unlink if you wish to link again.", r.DiscordID())
		}
	}
}
