package services

import (
	"encoding/json"
	"eth-gas-bot/modules/telegram/models"
	"eth-gas-bot/modules/telegram/models/entities"
	"eth-gas-bot/modules/telegram/models/params"
	"fmt"
)

func SendMessage(token string, params *params.SendMessageParams) (ret *entities.Message, err error) {
	resp, err := GenericCall(token, "/sendMessage", params)
	var r models.Response[entities.Message]
	err = json.Unmarshal(resp, &r)
	if err != nil {
		// Just return
	} else if !r.Ok {
		err = fmt.Errorf("%s", r.Description)
	} else {
		ret = r.Result
	}
	return
}
