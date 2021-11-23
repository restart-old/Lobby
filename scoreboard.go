package main

import (
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
)

func StartScoreboardTask(p *player.Player) {
	scoreboardTicker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-scoreboardTicker.C:
			scoreboard := scoreboard.New("hi")
			scoreboard.WriteString(p.Name())
			p.SendScoreboard(scoreboard)
		}
	}
}
