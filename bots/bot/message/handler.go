package message

import (
	"github.com/04Akaps/Scalable_Discord_bot/type/bot"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"strings"
)

type Handler struct {
}

type handler func(
	info *bot.BotHandler,
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	log *zap.Logger,
)

var mapper = map[string]handler{
	"!hello": helloHandler,
}

func NewMessageHandler() Handler {
	return Handler{}
}

func (h Handler) HandleMessage(
	log *zap.Logger,
	info *bot.BotHandler,
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	match := strings.Fields(info.ContentMatch)

	if len(match) > 0 {
		content := match[0]

		if function, ok := mapper[content]; ok {
			function(info, s, m, log)
		}
	}
}

func helloHandler(info *bot.BotHandler, s *discordgo.Session, m *discordgo.MessageCreate, log *zap.Logger) {
	s.ChannelMessageSend(m.ChannelID, info.Message)
}
