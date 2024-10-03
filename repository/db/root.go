package db

import (
	"github.com/04Akaps/Scalable_Discord_bot/config"
	"github.com/04Akaps/Scalable_Discord_bot/repository/db/mysql"
	. "github.com/04Akaps/Scalable_Discord_bot/type/bot"
)

type database struct {
	cfg config.Config

	sql mysql.SqlImpl
}

type DatabaseImpl interface {
	GetBotTotalInfo() ([]*BotInfo, error)
	GetBotHandler(name string) (map[string]*BotHandler, error)
}

var _ DatabaseImpl = (*database)(nil)

func NewDatabase(cfg config.Config) (DatabaseImpl, error) {
	d := &database{cfg: cfg}
	var err error

	if d.sql, err = mysql.NewSql(cfg); err != nil {
		return nil, err
	}

	return d, nil
}

func (d database) GetBotTotalInfo() ([]*BotInfo, error) {
	return d.sql.GetBotTotalInfo()
}

func (d database) GetBotHandler(name string) (map[string]*BotHandler, error) {
	return d.sql.GetBotHandler(name)
}
