package mysql

import (
	"github.com/04Akaps/Scalable_Discord_bot/config"
	"github.com/04Akaps/Scalable_Discord_bot/repository/db/mysql/bot"
	"github.com/04Akaps/Scalable_Discord_bot/repository/db/mysql/botHandler"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"

	. "github.com/04Akaps/Scalable_Discord_bot/type/bot"
)

// For Remove Full Group By
const sqlMode = "'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ALLOW_INVALID_DATES,ERROR_FOR_DIVISION_BY_ZERO'"

type sql struct {
	sess db.Session
	cfg  config.Config

	botInfo    bot.BotInfoTable
	botHandler botHandler.BotHandlerTable

	utils utils
}

type SqlImpl interface {
	GetBotTotalInfo() ([]*BotInfo, error)
	GetBotHandler(name string) (map[string]*BotHandler, error)
}

var _ SqlImpl = (*sql)(nil)

func NewSql(config config.Config) (SqlImpl, error) {
	s := &sql{cfg: config}
	var err error

	cfg := config.MySQL["discord"]

	if s.sess, err = mysql.Open(mysql.ConnectionURL{
		Database: cfg.Database,
		Host:     cfg.Host,
		User:     cfg.User,
		Password: cfg.Password,
		//Options: map[string]string{    		// if need add change sql mode in session
		//	"sql_mode": sqlMode,
		//},
	}); err != nil {
		return nil, err
	} else if err = s.sess.Ping(); err != nil {
		return nil, err
	} else {
		s.botInfo = bot.NewBotInfoTable(s.sess, s.sess.Collection(cfg.Collections["bot_info"]))
		s.botHandler = botHandler.NewBotHandlerTable(s.sess, s.sess.Collection(cfg.Collections["bot_handler"]))

		s.utils = NewSqlUtils(s.cfg)
	}

	return s, nil
}

func (s sql) GetBotTotalInfo() ([]*BotInfo, error) {
	return s.botInfo.GetBotTotalInfo()
}

func (s sql) GetBotHandler(name string) (map[string]*BotHandler, error) {
	return s.botHandler.GetBotHandler(name)
}
