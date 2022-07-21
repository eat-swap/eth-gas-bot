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
	"math"
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
	historyGas := global.GetHistoryGas()
	maxCount := []float64{20, 240, 1440, 5760}

	text := fmt.Sprintf("Base gas price: <b>%.6f", gas.Gas.SuggestBaseFee)
	for i, v := range historyGas {
		text += fmt.Sprintf("/%.3f", v.Base/math.Min(float64(v.Count), maxCount[i]))
	}
	text += "</b>\n"

	text += fmt.Sprintf("Low: <b>%d", gas.Gas.SafeGasPrice)
	for i, v := range historyGas {
		text += fmt.Sprintf("/%.3f", v.Low/math.Min(float64(v.Count), maxCount[i]))
	}
	text += "</b>\n"

	text += fmt.Sprintf("Avg: <b>%d", gas.Gas.ProposeGasPrice)
	for i, v := range historyGas {
		text += fmt.Sprintf("/%.3f", v.Avg/math.Min(float64(v.Count), maxCount[i]))
	}
	text += "</b>\n"

	text += fmt.Sprintf("High: <b>%d", gas.Gas.FastGasPrice)
	for i, v := range historyGas {
		text += fmt.Sprintf("/%.3f", v.High/math.Min(float64(v.Count), maxCount[i]))
	}
	text += "</b>\n"
	text += "(Current/5 Minutes/1 Hour/6 Hours/24 Hours)"

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
	historyPrice := global.GetHistoryPrice()

	text := fmt.Sprintf("Current Ethereum price: <b>%.3f</b>\n", price.Price.Usd)
	text += fmt.Sprintf("5 Minutes Average: <b>%.3f</b>\n", historyPrice[0])
	text += fmt.Sprintf("1 Hour Average: <b>%.3f</b>\n", historyPrice[1])
	text += fmt.Sprintf("6 Hours Average: <b>%.3f</b>\n", historyPrice[2])
	text += fmt.Sprintf("24 Hours Average: <b>%.3f</b>\n", historyPrice[3])

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
	return fmt.Sprintf("\n\nServer time: <b>%s</b>\nData time: <b>%s</b>\nNext update: <b>%s</b>",
		time.Now().Format("15:04:05.000"),
		obtained.Format("15:04:05.000"),
		global.NextUpdate.Format("15:04:05.000"),
	)
}
