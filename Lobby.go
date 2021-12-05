package main

import (
	"fmt"
	"image/png"
	"lobby/handler"
	"os"

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
			p.Handle(&handler.PlayerHandler{P: p})
			go handleJoin(p, wl, server)
			p.StopFlying()
		}
	}
}
func handleJoin(p *player.Player, wl *whitelist.WhiteList, server *server.Server) {
	fmt.Println(p.Name(), "is now connected with the ip:", p.Addr().String())
	if wl.Enabled && !wl.Whitelisted(p.Name()) {
		p.Disconnect("§9Server will be back soon\n§fhttp://sgpractice.tk/discord")
		return
	}
	p.SetGameMode(LobbyGm{})
	f, _ := os.Create("./data/skins/" + p.Name() + ".png")
	png.Encode(f, p.Skin())
}

type LobbyGm struct {
}

func (LobbyGm) AllowsEditing() bool      { return false }
func (LobbyGm) AllowsTakingDamage() bool { return false }
func (LobbyGm) CreativeInventory() bool  { return true }
func (LobbyGm) HasCollision() bool       { return false }
func (LobbyGm) AllowsFlying() bool       { return true }
func (LobbyGm) AllowsInteraction() bool  { return true }
func (LobbyGm) Visible() bool            { return true }
