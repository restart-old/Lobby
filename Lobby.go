package main

import (
	"lobby/handler"
	"time"

	"github.com/RestartFU/whitelist"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
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
			if !wl.Whitelisted(p.Name()) {
				go func() {
					time.Sleep(p.Latency() * 10)
					p.Disconnect("§cServer will be back soon\n§cdiscord.gg/tcY6bJv9nb")
				}()
			}
			p.Handle(&handler.PlayerHandler{P: p})
		}
	}
}
