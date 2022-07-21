package services

import (
	"eth-gas-bot/config"
	"eth-gas-bot/modules/eth-daemon/global"
	"eth-gas-bot/modules/telegram/models/entities"
	"eth-gas-bot/modules/telegram/models/params"
	"eth-gas-bot/modules/telegram/services"
	"eth-gas-bot/utils"
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

	text := fmt.Sprintf("Base gas price: <b>%.6f</b>\n", gas.Gas.SuggestBaseFee)
	text += fmt.Sprintf("Low: <b>%d</b>\n", gas.Gas.SafeGasPrice)
	text += fmt.Sprintf("Avg: <b>%d</b>\n", gas.Gas.ProposeGasPrice)
	text += fmt.Sprintf("High: <b>%d</b>\n", gas.Gas.FastGasPrice)

	text += appendTimestamp(&gas.ObtainedAt)

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
	text := fmt.Sprintf("Current Ethereum price: <b>%.3f</b>", price.Price.Usd)
	text += appendTimestamp(&price.Price.UsdTimestamp)

	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      text,
		ParseMode: params.MessageParseModeHTML,
	})
	if err != nil || r == nil {
		log.Printf("Error sending message: %s", err)
	}
}

func sendHelp(message *entities.Message) {
	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      fmt.Sprintf(DefaultMessage, utils.WrapForMarkdown(message.Chat.FirstName)),
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

func appendTimestamp(obtained *time.Time) string {
	return fmt.Sprintf("\n\nServer time: %s\nData time: %s\nNext update: %s",
		time.Now().Format(time.RFC1123),
		obtained.Format(time.RFC1123),
		global.NextUpdate.Format(time.RFC1123),
	)
}
