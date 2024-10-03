package bot

import (
	"encoding/json"
	"fmt"
	"github.com/04Akaps/Scalable_Discord_bot/bots/utils"
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
	utils.RunWork.Done()

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

	b.handleRequest(m.Content, s, m)
}

func (b *Bot) handleRequest(msg string, s *discordgo.Session, m *discordgo.MessageCreate) {

	info, ok := b.handler[msg]

	if !ok {
		return
	}

	switch info.Type {
	case MESSAGE:
		b.handleMessageType(info, s, m)
		return

	case MESSAGE_COMPLEX:
		b.handleComplexType(info, s, m)
		return
	}

}

func (b *Bot) handleMessageType(info *BotHandler, s *discordgo.Session, m *discordgo.MessageCreate) {
	switch info.ContentMatch {
	case "!hello":
		s.ChannelMessageSend(m.ChannelID, info.Message)
		return
	}
}

func (b *Bot) handleComplexType(info *BotHandler, s *discordgo.Session, m *discordgo.MessageCreate) {
	switch info.ContentMatch {
	case "!test":
		var data TestCallType

		if err := json.Unmarshal([]byte(info.Message), &data); err != nil {
			b.log.Error("Failed to unmarshal", zap.Error(err))
			return
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Example Title",
			Description: data.Content,
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
				Content:    embed.Description,
				Embed:      embed,
				Components: []discordgo.MessageComponent{discordgo.ActionsRow{Components: messageComponents}}}); err != nil {
			b.log.Error("Failed to send complex message", zap.Error(err))
		}

	}
}
