package bot

import (
	"context"
	. "github.com/04Akaps/Scalable_Discord_bot/type/bot"
	"github.com/upper/db/v4"
)

type botInfo struct {
	table db.Collection
	sess  db.Session
}

type BotInfoTable interface {
	GetBotTotalInfo() ([]*BotInfo, error)
}

var _ BotInfoTable = (*botInfo)(nil)

func NewBotInfoTable(sess db.Session, table db.Collection) BotInfoTable {
	b := botInfo{table: table, sess: sess}
	return b
}

func (b botInfo) GetBotTotalInfo() ([]*BotInfo, error) {

	var res []*BotInfo
	if err := b.table.Find().Select("channel_name", "bot_name", "bot_token").All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (b botInfo) TxContext(fn func(sess db.Session) error) error {
	return b.session().TxContext(context.Background(), fn, nil)
}

func (b botInfo) session() db.Session {
	return b.table.Session()
}
