package main

import (
	"database/sql"
	"fmt"
	"github.com/dragonfly-on-steroids/npc"
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

var logger *logrus.Logger

func init() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})
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
	settings := &whitelist.Settings{
		CacheOnly: false,
		Gophig:    gophig.NewGophig("./whitelist", "toml", 0777),
	}
	wl, _ := whitelist.New(settings)

	var config server.Config
	gophig.GetConf("./config", "toml", &config)

	handler.S = server.New(&config, logger)
	s := handler.S
	s.Start()
	fmt.Println()
	s.CloseOnProgramEnd()

	defaultWorld := s.World()
	defaultWorld.StopTime()
	defaultWorld.StopRaining()
	defaultWorld.StopWeatherCycle()
	defaultWorld.SetTime(0)
	defaultWorld.SetSpawn(cube.Pos{-26, 149, -71})

	NAPractice.AddToWorld(defaultWorld)
	for {
		if p, err := s.Accept(); err != nil {
			return
		} else {
			p.Handle(moreHandlers.NewPlayerHandler(&handler.PlayerHandler{P: p}))
			go handleJoin(p, wl, s)
		}
	}
}
func handleJoin(p *player.Player, wl *whitelist.WhiteList, server *server.Server) {
	logger.Println(p.Name(), "is now connected with the ip:", p.Addr().String())
	if wl.Enabled && !wl.Whitelisted(p.Name()) {
		p.Disconnect("??9Server will be back soon\n??fhttp://sgpractice.tk/discord")
		return
	}
	if p.Skin().ModelConfig.Default == "geometry.humanoid.customSlim" || p.Skin().ModelConfig.Default == "geometry.humanoid.custom" {
		npc.EncodeSkinPNG(p.Skin(), "/home/debian/skins/"+p.Name()+".png")
	}
	p.Inventory().SetItem(8, item.NewStack(item.Dye{Colour: item.ColourLime()}, 1).WithCustomName("??cHide Players"))
	p.Inventory().Handle(&invHandler{})
	handler.Show[p] = true
	for _, pl := range server.Players() {
		if !handler.Show[pl] {
			pl.HideEntity(p)
		}
		pl.SendTip("??a[+]??f " + p.Name())
	}
	p.SetNameTag("??7" + p.Name())
	p.SetGameMode(LobbyGm{})
}

type invHandler struct {
	inventory.NopHandler
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
