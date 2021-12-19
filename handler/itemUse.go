package handler

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

var Show = map[*player.Player]bool{}
var S *server.Server

func (h *PlayerHandler) HandleItemUse(*event.Context) {
	held, _ := h.P.HeldItems()
	Greydye := item.Dye{Colour: item.ColourGrey()}
	Greendye := item.Dye{Colour: item.ColourLime()}
	switch held.Item() {
	case Greendye:
		Show[h.P] = false
		h.P.Inventory().SetItem(8, item.NewStack(Greydye, 1).WithCustomName("§aShow Players"))
		for _, p := range S.Players() {
			h.P.HideEntity(p)
		}
	case Greydye:
		delete(Show, h.P)
		h.P.Inventory().SetItem(8, item.NewStack(Greendye, 1).WithCustomName("§cHide Players"))
		for _, p := range S.Players() {
			h.P.ShowEntity(p)
		}
	}
}
