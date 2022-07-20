package response

import (
	"eth-gas-bot/modules/etherscan/models"
	"strconv"
	"time"
)

type EthPriceResult struct {
	EthBtc          string `json:"ethbtc"`
	EthBtcTimestamp string `json:"ethbtc_timestamp"`
	EthUsd          string `json:"ethusd"`
	EthUsdTimestamp string `json:"ethusd_timestamp"`
}

func (t *EthPriceResult) ToPrice() (*models.Price, error) {
	btc, err := strconv.ParseFloat(t.EthBtc, 64)
	if err != nil {
		return nil, err
	}

	usd, err := strconv.ParseFloat(t.EthUsd, 64)
	if err != nil {
		return nil, err
	}

	btcTimestampInt, err := strconv.ParseInt(t.EthBtcTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	btcTimestamp := time.Unix(btcTimestampInt, 0)

	usdTimestampInt, err := strconv.ParseInt(t.EthUsdTimestamp, 10, 64)
	if err != nil {
		return nil, err
	}
	usdTimestamp := time.Unix(usdTimestampInt, 0)

	return &models.Price{
		Btc:          btc,
		Usd:          usd,
		BtcTimestamp: btcTimestamp,
		UsdTimestamp: usdTimestamp,
	}, nil
}
