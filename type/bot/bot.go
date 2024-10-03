package bot

type BotInfo struct {
	ChannelName string `json:"channel_name" db:"channel_name"`
	BotName     string `json:"bot_name" db:"bot_name"`
	BotToken    string `json:"bot_token" db:"bot_token"`
}

type BotHandler struct {
	ContentMatch string `json:"content_match" db:"content_match"`
	Type         int    `json:"type" db:"type"`
	Message      string `json:"message" db:"message"`
}
