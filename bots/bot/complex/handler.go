package complex

import (
	"encoding/json"
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
	"!test": testHandler,
}

func NewComplexHandler() Handler {
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

func testHandler(
	info *bot.BotHandler,
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	log *zap.Logger,
) {
	var data bot.TestCallType

	if err := json.Unmarshal([]byte(info.Message), &data); err != nil {
		log.Error("Failed to unmarshal", zap.Error(err))
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Example Title",
		Description: data.Content,

		// 이미지가 있는 경우
		//Image: &discordgo.MessageEmbedImage{
		//	URL: "https://externlabs.com/blogs/wp-content/uploads/2023/04/discord-bot-1.jpg",
		//},

		//https://picsum.photos/500?random=<random>
	}

	var messageComponents []discordgo.MessageComponent

	for _, comp := range data.Components {
		for _, innerComp := range comp.Components {
			messageComponents = append(messageComponents, discordgo.Button{
				Label:    innerComp.Label,
				Style:    discordgo.ButtonStyle(innerComp.Style),
				CustomID: innerComp.CustomID,
			})
		}
	}

	if _, err := s.ChannelMessageSendComplex(
		m.ChannelID,
		&discordgo.MessageSend{
			Content: embed.Description,
			Embed:   embed,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: messageComponents}}},
	); err != nil {
		log.Error("Failed to send complex message", zap.Error(err))
	}

}
