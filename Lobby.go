package main

import (
	"lobby/handler"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
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
			p.Move(mgl64.Vec3{0, 0, 0}, -90, 0)
			p.Handle(&handler.PlayerHandler{P: p})
		}
	}
}
