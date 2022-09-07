package main

import (
	"flag"
	"log"

	"github.com/Oybek-uzb/mana-taxi-bot-service/internal"
	"github.com/Oybek-uzb/mana-taxi-bot-service/internal/config"
	"github.com/Oybek-uzb/mana-taxi-bot-service/pkg/logging"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "configs/dev.yml", "config file path")
}

func main() {
	flag.Parse()

	log.Println("Config initializing")
	cfg := config.GetConfig(cfgPath)

	log.Println("Logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	logger.Println("Creating Application")
	app, err := internal.NewApp(logger, cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	app.Run()
}
