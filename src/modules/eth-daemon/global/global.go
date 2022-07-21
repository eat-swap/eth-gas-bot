package global

import (
	"eth-gas-bot/modules/eth-daemon/models"
	"sync"
	"time"
)

const (
	CacheLimit             = 10000
	RefreshIntervalSeconds = 15
	RefreshInterval        = RefreshIntervalSeconds * time.Second
)

var (
	HistoryPrice [CacheLimit]models.PriceInfo
	PriceHead           = 0
	PriceMutex          = &sync.RWMutex{}
	PriceCounter uint64 = 0

	HistoryGas [CacheLimit]models.GasInfo
	GasHead           = 0
	GasMutex          = &sync.RWMutex{}
	GasCounter uint64 = 0
)

func GetCurrentGas() models.GasInfo {
	GasMutex.RLock()
	defer GasMutex.RUnlock()
	return HistoryGas[GasHead]
}

func GetCurrentPrice() models.PriceInfo {
	PriceMutex.RLock()
	defer PriceMutex.RUnlock()
	return HistoryPrice[PriceHead]
}
