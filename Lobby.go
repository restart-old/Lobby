package main

import (
	"fmt"
	"lobby/handler"

	"github.com/RestartFU/whitelist"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

func main() {
	wl, _ := whitelist.New("./whitelist.json")
	config := readConfig()
	log := logger()

	server := server.New(&config, log)
	server.Start()
	server.CloseOnProgramEnd()

	defaultWorld := server.World()

	defaultWorld.StopTime()
	defaultWorld.StopRaining()
	defaultWorld.StopWeatherCycle()
	defaultWorld.SetTime(0)
	defaultWorld.SetSpawn(cube.Pos{-26, 149, -71})

	NAPractice.AddToWorld(defaultWorld)

	for {
		if p, err := server.Accept(); err != nil {
			return
		} else {
			fmt.Println(p.Name(), "is now connected with the ip:", p.Addr().String())
			if wl.Enabled {
				disconnectIfNotWhiteListed(wl, p, server)
			}
			p.Handle(&handler.PlayerHandler{P: p})
		}
	}
}

func disconnectIfNotWhiteListed(wl *whitelist.WhiteList, p *player.Player, server *server.Server) {
	if !wl.Whitelisted(p.Name()) {
		p.SetGameMode(world.GameModeSurvival)
		go func() {
			for _, ok := server.Player(p.UUID()); ok; {
				if p.OnGround() {
					p.Disconnect("§9Server will be back soon\n§fhttp://sgpractice.tk/discord")
				}
				continue
			}
		}()
	}
}
