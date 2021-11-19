package main

import (
	"fmt"
	"lobby/handler"
	"time"

	"github.com/df-mc/dragonfly/server"
	"github.com/sandertv/gophertunnel/query"
)

func main() {
	go func() {
		for {
			var newTag string
			q, err := query.Do("nitrofaction.fr:19132")
			if err != nil {
				newTag = "§9Practice\n\uE300\n§cOFFLINE"
			} else {
				newTag = fmt.Sprintf("§9Practice\n\uE300\n§a%v/%v", q["numplayers"], q["maxplayers"])
			}
			SlapperPractice.SetNameTag(newTag)
			time.Sleep(1 * time.Second)
		}
	}()

	config := readConfig()
	log := logger()

	server := server.New(&config, log)
	server.Start()
	server.CloseOnProgramEnd()

	defaultWorld := server.World()

	defaultWorld.StopTime()
	defaultWorld.StopWeatherCycle()
	defaultWorld.SetTime(0)

	SlapperPractice.AddToWorld(defaultWorld)

	for {
		if p, err := server.Accept(); err != nil {
			return
		} else {
			p.Handle(&handler.PlayerHandler{P: p})
		}
	}
}
