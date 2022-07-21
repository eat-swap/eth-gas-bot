package services

import (
	"eth-gas-bot/config"
	"eth-gas-bot/modules/telegram/models/entities"
	"eth-gas-bot/modules/telegram/models/params"
	"eth-gas-bot/modules/telegram/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleIncomingMessage(message *entities.Message, ctx *gin.Context) {
	var token = config.TelegramBotToken

	switch message.Text {
	default:
		r, err := services.SendMessage(token, &params.SendMessageParams{
			ChatId:                   message.Chat.Id,
			Text:                     ErrorMessage,
			ParseMode:                params.MessageParseModeMarkdown,
			ReplyToMessageId:         message.MessageId,
			AllowSendingWithoutReply: true,
		})
		if err != nil || r == nil {
			log.Printf("Error sending message: %s", err)
		}
		fallthrough
	case "/help":
		r, err := services.SendMessage(token, &params.SendMessageParams{
			ChatId:    message.Chat.Id,
			Text:      fmt.Sprintf(DefaultMessage, message.Chat.FirstName),
			ParseMode: params.MessageParseModeMarkdown,
		})
		if err != nil || r == nil {
			log.Printf("Error sending message: %s", err)
		}
	case "/gas":
		r, err := services.SendMessage(token, &params.SendMessageParams{
			ChatId:    message.Chat.Id,
			Text:      fmt.Sprintf("The current gas price is %s", "*Not implemented yet*"),
			ParseMode: params.MessageParseModeMarkdown,
		})
		if err != nil || r == nil {
			log.Printf("Error sending message: %s", err)
		}
	case "/eth":
		r, err := services.SendMessage(token, &params.SendMessageParams{
			ChatId:    message.Chat.Id,
			Text:      fmt.Sprintf("The current Ethereum price is %s", "*Not implemented yet*"),
			ParseMode: params.MessageParseModeMarkdown,
		})
		if err != nil || r == nil {
			log.Printf("Error sending message: %s", err)
		}
	}
}
