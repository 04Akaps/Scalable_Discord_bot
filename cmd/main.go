package main

import (
	"flag"
	"github.com/04Akaps/Scalable_Discord_bot/cmd/app"
	"github.com/04Akaps/Scalable_Discord_bot/config"
	"go.uber.org/zap"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*configFlag)
	cfg.Logger = zap.Must(zap.NewProduction())
	app.NewApp(*cfg)
}
