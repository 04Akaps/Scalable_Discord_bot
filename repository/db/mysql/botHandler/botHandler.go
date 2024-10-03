package botHandler

import (
	"context"
	. "github.com/04Akaps/Scalable_Discord_bot/type/bot"
	"github.com/upper/db/v4"
)

type botHandler struct {
	table db.Collection
	sess  db.Session
}

type BotHandlerTable interface {
	GetBotHandler(name string) (map[string]*BotHandler, error)
}

var _ BotHandlerTable = (*botHandler)(nil)

func NewBotHandlerTable(sess db.Session, table db.Collection) BotHandlerTable {
	b := botHandler{table: table, sess: sess}
	return b
}

func (b botHandler) GetBotHandler(name string) (map[string]*BotHandler, error) {

	var res []*BotHandler

	err := b.table.Find(db.Cond{"bot_name": name}).
		Select("content_match", "type", "message").All(&res)

	if err != nil {
		return nil, err
	}

	result := make(map[string]*BotHandler, len(res))

	for _, h := range res {
		result[h.ContentMatch] = h
	}

	return result, nil
}

func (b botHandler) TxContext(fn func(sess db.Session) error) error {
	return b.session().TxContext(context.Background(), fn, nil)
}

func (b botHandler) session() db.Session {
	return b.table.Session()
}
