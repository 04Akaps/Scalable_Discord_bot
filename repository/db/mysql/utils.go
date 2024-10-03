package mysql

import "github.com/04Akaps/Scalable_Discord_bot/config"

type utils struct {
	cfg config.Config
}

func NewSqlUtils(cfg config.Config) utils {
	su := utils{cfg: cfg}

	return su
}
