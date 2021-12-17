package main

import (
	"database/sql"
	"fmt"
	"lobby/commands"
	"lobby/handler"

	"github.com/RestartFU/gophig"
	"github.com/RestartFU/whitelist"
	"github.com/SGPractice/link"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/dragonfly-on-steroids/moreHandlers"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var wl *whitelist.WhiteList
var logger *logrus.Logger

func init() {
	wl, _ = whitelist.New("./whitelist.json")
	logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{ForceColors: true}
}

func init() {
	var config mysql.Config
	gophig.GetConf("./mysql", "toml", &config)
	connector, _ := mysql.NewConnector(&config)
	db := sql.OpenDB(connector)

	storer := link.NewJSONStorer("/home/debian/link/")
	commands.Linker = link.NewLinker(db, storer)

	LINK := cmd.New("link", "link this account with your discord account!", nil, &commands.LinkCommand{})
	UNLINK := cmd.New("unlink", "unlink this account from your currently linked discord account!", nil, &commands.UnlinkCommand{})

	cmd.Register(LINK)
	cmd.Register(UNLINK)
}

func main() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	var config server.Config
	gophig.GetConf("./config", "toml", &config)

	server := server.New(&config, logger)
	server.Start()
	fmt.Println()
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
			p.Handle(moreHandlers.NewPlayerHandler(&handler.PlayerHandler{P: p}))
			go handleJoin(p, wl, server)
		}
	}
}
func handleJoin(p *player.Player, wl *whitelist.WhiteList, server *server.Server) {
	logger.Println(p.Name(), "is now connected with the ip:", p.Addr().String())
	if wl.Enabled && !wl.Whitelisted(p.Name()) {
		p.Disconnect("§9Server will be back soon\n§fhttp://sgpractice.tk/discord")
		return
	}
	p.Inventory().Handle(&invHandler{})
	for _, pl := range server.Players() {
		pl.SendTip("§a[+]§f " + p.Name())
	}
	p.SetNameTag("§7" + p.Name())
	p.SetGameMode(LobbyGm{})
}

type invHandler struct {
	inventory.NopHandler
}

func (*invHandler) HandlePlace(ctx *event.Context, slot int, it item.Stack) {
	ctx.Cancel()
}
func (*invHandler) HandleDrop(ctx *event.Context, slot int, it item.Stack) {
	ctx.Cancel()
}
func (*invHandler) HandleTake(ctx *event.Context, slot int, it item.Stack) {
	ctx.Cancel()
}

type LobbyGm struct {
}

func (LobbyGm) AllowsEditing() bool      { return false }
func (LobbyGm) AllowsTakingDamage() bool { return false }
func (LobbyGm) CreativeInventory() bool  { return false }
func (LobbyGm) HasCollision() bool       { return true }
func (LobbyGm) AllowsFlying() bool       { return false }
func (LobbyGm) AllowsInteraction() bool  { return true }
func (LobbyGm) Visible() bool            { return true }
