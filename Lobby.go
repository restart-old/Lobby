package main

import (
	"fmt"
	"time"

	"github.com/RestartFU/whitelist"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
)

var wl *whitelist.WhiteList
var config server.Config

func init() {
	wl, _ = whitelist.New("./whitelist.json")
	config = readConfig()
}

func main() {
	server := server.New(&config, logger())
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
			go handleJoin(p, wl, server)
		}
	}
}
func handleJoin(p *player.Player, wl *whitelist.WhiteList, server *server.Server) {
	fmt.Println(p.Name(), "is now connected with the ip:", p.Addr().String())
	if wl.Enabled {
		disconnectIfNotWhiteListed(wl, p, server)
	}
}

func disconnectIfNotWhiteListed(wl *whitelist.WhiteList, p *player.Player, server *server.Server) {
	if !wl.Whitelisted(p.Name()) {
		time.Sleep(500 * time.Millisecond)
		p.Disconnect("§9Server will be back soon\n§fhttp://sgpractice.tk/discord")

	}
}
