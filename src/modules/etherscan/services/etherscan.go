package services

import (
	"encoding/json"
	"eth-gas-bot/config"
	"eth-gas-bot/modules/etherscan/models"
	"eth-gas-bot/modules/etherscan/models/response"
	"eth-gas-bot/utils"
	"fmt"
)

const (
	Endpoint = "https://api.etherscan.io/api"
)

func GetGas() (*models.Gas, error) {
	// println(config.EtherScanApiKey)
	resp, err := utils.HttpGetWithParams(Endpoint, nil, map[string]string{
		"module": "gastracker",
		"action": "gasoracle",
		"apikey": config.EtherScanApiKey,
	})
	if err != nil {
		return nil, err
	}

	gas := &response.GeneralResponse[response.GasInfoResult]{}
	err = json.Unmarshal(resp, gas)
	if err != nil {
		return nil, err
	}

	return gas.Result.ToGas()
}

func GetPrice() (*models.Price, error) {
	resp, err := utils.HttpGetWithParams(Endpoint, nil, map[string]string{
		"module": "stats",
		"action": "ethprice",
		"apikey": config.EtherScanApiKey,
	})
	if err != nil {
		return nil, err
	}

	price := &response.GeneralResponse[response.EthPriceResult]{}
	err = json.Unmarshal(resp, price)
	if err != nil {
		fmt.Println(string(resp))
		return nil, err
	}

	return price.Result.ToPrice()
}
