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
	price := global.GetCurrentPrice().Price.Usd
	historyGas := global.GetHistoryGas()
	maxCount := []float64{20, 240, 1440, 5760}
	estimateMap := map[string]float64{
		"Transfer ETH":     21000,
		"Transfer USDT":    54128,
		"Uniswap V2: Swap": 152809,
	}

	var text string

	text += fmt.Sprintf("`%3d", gas.Gas.SafeGasPrice)
	for i, v := range historyGas {
		if i == 0 {
			text += fmt.Sprintf(" | %.2f", v.Low/math.Min(float64(v.Count), maxCount[i]))
		} else {
			text += fmt.Sprintf(" | %.3f", v.Low/math.Min(float64(v.Count), maxCount[i]))
		}
	}
	text += "`  _Low_\n"

	text += fmt.Sprintf("`%3d", gas.Gas.ProposeGasPrice)
	for i, v := range historyGas {
		if i == 0 {
			text += fmt.Sprintf(" | %.2f", v.Avg/math.Min(float64(v.Count), maxCount[i]))
		} else {
			text += fmt.Sprintf(" | %.3f", v.Avg/math.Min(float64(v.Count), maxCount[i]))
		}
	}
	text += "`  _Avg_\n"

	text += fmt.Sprintf("`%3d", gas.Gas.FastGasPrice)
	for i, v := range historyGas {
		if i == 0 {
			text += fmt.Sprintf(" | %.2f", v.High/math.Min(float64(v.Count), maxCount[i]))
		} else {
			text += fmt.Sprintf(" | %.3f", v.High/math.Min(float64(v.Count), maxCount[i]))
		}
	}
	text += "`  _High_\n"

	text += "Current | 5 Minutes | 1 Hour | 6 Hours | 24 Hours\n"

	text += fmt.Sprintf("`%.6f", gas.Gas.SuggestBaseFee)
	for i, v := range historyGas {
		text += fmt.Sprintf(" | %.3f", v.Base/math.Min(float64(v.Count), maxCount[i]))
	}
	text += "`  Base\n\nEstimated transaction fees:\n"

	for k, v := range estimateMap {
		text += fmt.Sprintf("`$%6.3f | $%6.3f | $%6.3f`  %s\n", float64(gas.Gas.SafeGasPrice)*v*1e-9*price, float64(gas.Gas.ProposeGasPrice)*v*1e-9*price, float64(gas.Gas.FastGasPrice)*v*1e-9*price, k)
	}

	text += appendTimestamp(&gas.ObtainedAt)

	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      utils.WrapForMarkdownWorse(text),
		ParseMode: params.MessageParseModeMarkdown,
	})
	if err != nil || r == nil {
		log.Printf("Error sending message: %s", err)
	}
}

func sendPriceInfo(message *entities.Message) {
	price := global.GetCurrentPrice()
	historyPrice := global.GetHistoryPrice()

	text := fmt.Sprintf("`%.3f`  Current Ethereum price\n", price.Price.Usd)
	text += fmt.Sprintf("`%.3f`  5 Minutes Average\n", historyPrice[0])
	text += fmt.Sprintf("`%.3f`  1 Hour Average\n", historyPrice[1])
	text += fmt.Sprintf("`%.3f`  6 Hours Average\n", historyPrice[2])
	text += fmt.Sprintf("`%.3f`  24 Hours Average\n", historyPrice[3])

	text += appendTimestamp(&price.Price.UsdTimestamp)

	r, err := services.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:    message.Chat.Id,
		Text:      utils.WrapForMarkdownWorse(text),
		ParseMode: params.MessageParseModeMarkdown,
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
	return fmt.Sprintf("\n*%s* | Server Time\n*%s* | Data Time\n*%s* | Next Update",
		time.Now().Format("15:04:05.000"),
		obtained.Format("15:04:05.000"),
		global.NextUpdate.Format("15:04:05.000"),
	)
}
