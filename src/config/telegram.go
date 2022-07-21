package config

import "os"

const (
	WebhookEntry = "/TQbqEEX9"
)

var (
	TelegramBotToken string
	TelegramApiRoot  string
)

func init() {
	TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	TelegramApiRoot = "https://api.telegram.org/bot" + TelegramBotToken
}

func GetApiRootByToken(token string) string {
	return "https://api.telegram.org/bot" + token
}
