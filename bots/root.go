package bots

import (
	"github.com/04Akaps/Scalable_Discord_bot/bots/utils"
	"time"

	"github.com/04Akaps/Scalable_Discord_bot/bots/bot"
	"github.com/04Akaps/Scalable_Discord_bot/config"
	"github.com/04Akaps/Scalable_Discord_bot/repository/db"
	"go.uber.org/zap"

	. "github.com/04Akaps/Scalable_Discord_bot/type/bot"
)

type Bots struct {
	cfg config.Config

	db   db.DatabaseImpl
	bots map[string]*BotInfo
}

func NewBots(cfg config.Config, db db.DatabaseImpl) Bots {
	bs := Bots{cfg: cfg, db: db, bots: make(map[string]*BotInfo)}

	bots, err := bs.db.GetBotTotalInfo()

	if err != nil {
		bs.cfg.Logger.Error("Failed to get all bot info", zap.Error(err))
		panic(err)
	}

	utils.RunWork.Add(len(bots))

	for _, info := range bots {
		// Left Join을 통해서 Handler 정보도 모두 가져 올 수 있지만, 데이터 양 자체가 그렇게 많지 않고,
		// Null 조건에 대해서 검증하는 로직보다는 그냥 쿼리를 한번 더 날리는 것이 깔끔하다 생각해서, 다음과 같이 구성
		// 필요하다면 Left Join 결과에 대한 Null 처리를 통해서 처리하는 방식도 가능
		handler, err := bs.db.GetBotHandler(info.BotName)

		if err != nil {
			bs.cfg.Logger.Error("Failed to get handler", zap.Error(err))
			panic(err)
		}

		bs.bots[info.BotName] = info
		go bot.NewBot(info, handler, cfg.Logger)
	}

	if len(bs.bots) == 0 {
		panic("no bots for running")
	}

	utils.RunWork.Wait()

	return bs
}

func (bs Bots) Run() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// TODO 쓸만한 로그 추가
		}
	}

}
