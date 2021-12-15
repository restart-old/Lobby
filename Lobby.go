package main

import (
	"database/sql"
	"fmt"
	"lobby/handler"

	"github.com/RestartFU/whitelist"
	"github.com/SGPractice/link"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/dragonfly-on-steroids/moreHandlers"
	"github.com/go-sql-driver/mysql"
)

var wl *whitelist.WhiteList
var config server.Config

func init() {
	wl, _ = whitelist.New("./whitelist.json")
	config = readConfig()
}

func init() {
	config := mysql.NewConfig()
	config.DBName = "GlowHCF"
	config.User = "root"
	config.Addr = ":3306"
	config.Passwd = "f37JZEUm2QFexguhRuyscW{AdrKr86KajFGf%VT2h6BJUUF"
	config.Net = "tcp"

	connector, _ := mysql.NewConnector(config)
	db := sql.OpenDB(connector)

	storer := link.NewJSONStorer("/home/debian/link/")
	linker := link.NewLinker(db, storer)

	LINK := cmd.New("link", "idk", nil, &LinkCommand{linker: linker})
	UNLINK := cmd.New("unlink", "idk", nil, &UnlinkCommand{linker: linker})

	cmd.Register(LINK)
	cmd.Register(UNLINK)
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
			p.Handle(moreHandlers.New(&handler.PlayerHandler{P: p}))
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
