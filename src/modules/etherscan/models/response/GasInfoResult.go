package response

import (
	"eth-gas-bot/modules/etherscan/models"
	"strconv"
	"strings"
)

type GasInfoResult struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

func (t *GasInfoResult) toGas() (*models.Gas, error) {
	lastBlock, err := strconv.ParseInt(t.LastBlock, 10, 64)
	if err != nil {
		return nil, err
	}

	safeGasPrice, err := strconv.ParseInt(t.SafeGasPrice, 10, 16)
	if err != nil {
		return nil, err
	}

	proposeGasPrice, err := strconv.ParseInt(t.ProposeGasPrice, 10, 16)
	if err != nil {
		return nil, err
	}

	fastGasPrice, err := strconv.ParseInt(t.FastGasPrice, 10, 16)
	if err != nil {
		return nil, err
	}

	suggestBaseFee, err := strconv.ParseFloat(t.SuggestBaseFee, 64)
	if err != nil {
		return nil, err
	}

	usedRatioStr := strings.Split(t.GasUsedRatio, ",")
	usedRatio := make([]float64, len(usedRatioStr))
	for i, v := range usedRatioStr {
		usedRatio[i], err = strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
	}

	return &models.Gas{
		LastBlock:       lastBlock,
		SafeGasPrice:    int16(safeGasPrice),
		ProposeGasPrice: int16(proposeGasPrice),
		FastGasPrice:    int16(fastGasPrice),
		SuggestBaseFee:  suggestBaseFee,
		GasUsedRatio:    usedRatio,
	}, nil
}
