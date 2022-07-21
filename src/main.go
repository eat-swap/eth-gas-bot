package main

import (
	"encoding/json"
	"eth-gas-bot/config"
	EtherscanServices "eth-gas-bot/modules/etherscan/services"
	"eth-gas-bot/modules/telegram/models/params"
	TelegramServices "eth-gas-bot/modules/telegram/services"
	"fmt"
)

func main() {
	debug()

	// http.RunServer()
}

func debug() {
	gas, err := EtherscanServices.GetGas()
	if err != nil {
		panic(err)
	}
	s, _ := json.MarshalIndent(gas, "", "  ")
	// st := "```\n" + string(s) + "\n```"

	ret, err := TelegramServices.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:              0,
		Text:                "+++***-*/*",
		ParseMode:           params.MessageParseModeMarkdown,
		DisableNotification: true,
		ProtectedContent:    true,
	})

	if err != nil {
		panic(err)
	}

	s, _ = json.MarshalIndent(ret, "", "  ")
	fmt.Println(string(s))
}
