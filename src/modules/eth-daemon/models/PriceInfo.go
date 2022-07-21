package models

import (
	"eth-gas-bot/modules/etherscan/models"
	"time"
)

type PriceInfo struct {
	Valid      bool          `json:"valid"`
	Price      *models.Price `json:"price,omitempty"`
	ObtainedAt time.Time     `json:"obtained_at,omitempty"`
}
