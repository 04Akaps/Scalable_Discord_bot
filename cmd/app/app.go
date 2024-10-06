package app

import (
	"github.com/04Akaps/Scalable_Discord_bot/bots"
	"github.com/04Akaps/Scalable_Discord_bot/config"
	"github.com/04Akaps/Scalable_Discord_bot/repository/db"
)

type App struct {
	cfg config.Config

	db db.DatabaseImpl
}

func NewApp(cfg config.Config) {
	a := App{cfg: cfg}
	var err error

	if a.db, err = db.NewDatabase(cfg); err != nil {
		panic(err)
	}

	bots.RunBots(cfg, a.db)
}
