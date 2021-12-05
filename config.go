package main

import (
	"encoding/json"
	"os"

	"github.com/df-mc/dragonfly/server"
	"github.com/sirupsen/logrus"
)

func readConfig() server.Config {
	config := server.DefaultConfig()
	content, err := os.ReadFile("./config.json")
	if err != nil {
		b, _ := json.MarshalIndent(config, "", "\t")
		os.WriteFile("./config.json", b, 0777)
		return config
	}
	json.Unmarshal(content, &config)
	return config
}
func logger() *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel
	return log
}
