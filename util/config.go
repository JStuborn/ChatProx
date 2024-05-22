package util

type Config struct {
	DiscordBotToken  string `json:"discord_bot_token"`
	TelegramBotToken string `json:"telegram_bot_token"`
	TelegramChatID   int64  `json:"telegram_chat_id"`
	DiscordChannelID string `json:"discord_channel_id"`
	TelegramTimeout  int    `json:"telegram_timeout"`
}
