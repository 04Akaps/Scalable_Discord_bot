package bots

import (
	"github.com/04Akaps/Scalable_Discord_bot/bots/bot"
	"github.com/04Akaps/Scalable_Discord_bot/config"
	"github.com/04Akaps/Scalable_Discord_bot/repository/db"
	"go.uber.org/zap"
	"runtime"
	"time"
)

type Bots struct {
	cfg config.Config

	db db.DatabaseImpl

	log *zap.Logger
}

func RunBots(cfg config.Config, db db.DatabaseImpl) {
	bs := Bots{cfg: cfg, db: db, log: cfg.Logger}

	bots, err := bs.db.GetBotTotalInfo()

	if err != nil {
		bs.cfg.Logger.Error("Failed to get all bot info", zap.Error(err))
		panic(err)
	}

	for _, info := range bots {
		// Left Join을 통해서 Handler 정보도 모두 가져 올 수 있지만, 데이터 양 자체가 그렇게 많지 않고,
		// Null 조건에 대해서 검증하는 로직보다는 그냥 쿼리를 한번 더 날리는 것이 깔끔하다 생각해서, 다음과 같이 구성
		// 필요하다면 Left Join 결과에 대한 Null 처리를 통해서 처리하는 방식도 가능
		handler, err := bs.db.GetBotHandler(info.BotName)

		if err != nil {
			bs.cfg.Logger.Error("Failed to get handler", zap.Error(err))
			panic(err)
		}

		go bot.NewBot(info, handler, cfg.Logger)
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	bs.log.Info("Module is running.... ")

	for {
		select {
		case <-ticker.C:
			bs.log.Info("", zap.Int("CPUs", runtime.NumCPU()), zap.Int("Goroutine", runtime.NumGoroutine()))
		}
	}

}
