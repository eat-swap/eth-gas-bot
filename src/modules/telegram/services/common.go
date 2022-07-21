package services

import (
	"encoding/json"
	"eth-gas-bot/config"
	"eth-gas-bot/utils"
)

func GenericCall(token, path string, params interface{}) ([]byte, error) {
	url := config.GetApiRootByToken(token) + path
	s, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	return utils.HttpPost(url, s, map[string][]string{
		"Content-Type": {"application/json"},
	})
}
