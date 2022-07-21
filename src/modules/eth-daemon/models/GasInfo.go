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
