package bot

import (
	"fmt"
	"github.com/04Akaps/Scalable_Discord_bot/bots/bot/complex"
	"github.com/04Akaps/Scalable_Discord_bot/bots/bot/message"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"time"

	. "github.com/04Akaps/Scalable_Discord_bot/type/bot"
)

const BOT = "Bot"

const (
	MESSAGE = iota
	MESSAGE_COMPLEX
)

type Bot struct {
	info    *BotInfo
	handler map[string]*BotHandler
	log     *zap.Logger

	messageHandler message.Handler
	complexHandler complex.Handler
}

func NewBot(
	info *BotInfo,
	handler map[string]*BotHandler,
	log *zap.Logger,
) *Bot {
	b := Bot{
		info:    info,
		handler: handler,
		log:     log,
	}

	b.messageHandler = message.NewMessageHandler()
	b.complexHandler = complex.NewComplexHandler()

	b.RunBot()

	return &b
}

func (b *Bot) RunBot() {

	name := b.info.BotName
	channel := b.info.ChannelName
	token := b.info.BotToken

	sess, err := discordgo.New(fmt.Sprintf("%s %s", BOT, token))

	if err != nil {
		b.log.Error("Failed to create bot session", zap.String("channel", channel), zap.String("name", name))
		return
	}

	sess.AddHandler(b.addHandler)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if err = sess.Open(); err != nil {
		b.log.Error("Failed to open bot session", zap.String("channel", channel), zap.String("name", name))
		return
	}

	defer sess.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.log.Info("BotRunning... ", zap.String("name", name), zap.String("channel", channel))
		}
	}
}

func (b *Bot) addHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	info, ok := b.handler[m.Content]

	if !ok {
		return
	}

	switch info.Type {
	case MESSAGE:
		b.messageHandler.HandleMessage(b.log, info, s, m)
		return
	case MESSAGE_COMPLEX:
		b.complexHandler.HandleMessage(b.log, info, s, m)
		return
	}

}
