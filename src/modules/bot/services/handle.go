package services

import (
	"eth-gas-bot/config"
	"eth-gas-bot/modules/eth-daemon/global"
	"eth-gas-bot/modules/telegram/models/entities"
	"eth-gas-bot/modules/telegram/models/params"
	"eth-gas-bot/modules/telegram/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func HandleIncomingMessage(message *entities.Message, ctx *gin.Context) {
	switch message.Text {
	default:
		sendError(message)
		fallthrough
	case "/help":
		sendHelp(message)
	case "/gas":
		sendGasInfo(message)
	case "/eth":
		sendPriceInfo(message)
	}
}

func sendGasInfo(message *entities.Message) {
	gas := global.GetCurrentGas()
	text := fmt.Sprintf("The current gas price is <b>%.6f</b>", gas.Gas.SuggestBaseFee)

	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      text,
		ParseMode: params.MessageParseModeHTML,
	})
	if err != nil || r == nil {
		log.Printf("Error sending message: %s", err)
	}
}

func sendPriceInfo(message *entities.Message) {
	price := global.GetCurrentPrice()
	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      fmt.Sprintf("The current Ethereum price is <b>%.3f</b>", price.Price.Usd),
		ParseMode: params.MessageParseModeHTML,
	})
	if err != nil || r == nil {
		log.Printf("Error sending message: %s", err)
	}
}

func sendHelp(message *entities.Message) {
	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      DefaultMessage,
		ParseMode: params.MessageParseModeMarkdown,
	})
	if err != nil || r == nil {
		log.Printf("Error sending message: %s", err)
	}
}

func sendError(message *entities.Message) {
	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:                   message.Chat.Id,
		Text:                     ErrorMessage,
		ParseMode:                params.MessageParseModeMarkdown,
		ReplyToMessageId:         message.MessageId,
		AllowSendingWithoutReply: true,
	})
	if err != nil || r == nil {
		log.Printf("Error sending message: %s", err)
	}
}

func appendTimestamp(text string, obtained *time.Time) string {
	return fmt.Sprintf("%s\nServer time: %s\nData time: %s\n", text, time.Now().Format(time.RFC1123))
}
