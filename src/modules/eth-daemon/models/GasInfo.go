package models

import (
	"eth-gas-bot/modules/etherscan/models"
	"time"
)

type GasInfo struct {
	Valid      bool        `json:"valid"`
	Gas        *models.Gas `json:"gas,omitempty"`
	ObtainedAt time.Time   `json:"obtained_at,omitempty"`
}

type HistoricalGas struct {
	Base  float64
	Low   float64
	Avg   float64
	High  float64
	Count uint64
}
