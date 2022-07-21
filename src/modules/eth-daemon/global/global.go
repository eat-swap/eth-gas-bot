package global

import (
	"eth-gas-bot/modules/eth-daemon/models"
	"math"
	"sync"
	"time"
)

const (
	CacheLimit             = 10000
	RefreshIntervalSeconds = 15
	RefreshInterval        = RefreshIntervalSeconds * time.Second
)

var (
	HistoryPrice  [CacheLimit]models.PriceInfo
	PriceHead            = 0
	PriceMutex           = &sync.RWMutex{}
	PriceCounter  uint64 = 0
	AvgPriceMutex        = &sync.RWMutex{}
	AvgPrice5Min         = 0.0
	AvgPrice1h           = 0.0
	AvgPrice6h           = 0.0
	AvgPrice24h          = 0.0

	HistoryGas  [CacheLimit]models.GasInfo
	GasHead            = 0
	GasMutex           = &sync.RWMutex{}
	GasCounter  uint64 = 0
	AvgGasMutex        = &sync.RWMutex{}
	AvgGas5Min  models.HistoricalGas
	AvgGas1h    models.HistoricalGas
	AvgGas6h    models.HistoricalGas
	AvgGas24h   models.HistoricalGas

	NextUpdate time.Time
)

func GetCurrentGas() models.GasInfo {
	GasMutex.RLock()
	defer GasMutex.RUnlock()
	return HistoryGas[GasHead]
}

func GetHistoryGas() []models.HistoricalGas {
	GasMutex.RLock()
	defer GasMutex.RUnlock()
	return []models.HistoricalGas{
		AvgGas5Min,
		AvgGas1h,
		AvgGas6h,
		AvgGas24h,
	}
}

func GetCurrentPrice() models.PriceInfo {
	PriceMutex.RLock()
	defer PriceMutex.RUnlock()
	return HistoryPrice[PriceHead]
}

func GetHistoryPrice() []float64 {
	PriceMutex.RLock()
	defer PriceMutex.RUnlock()
	return []float64{
		AvgPrice5Min / math.Min(float64(PriceCounter), 20.0),
		AvgPrice1h / math.Min(float64(PriceCounter), 240.0),
		AvgPrice6h / math.Min(float64(PriceCounter), 1440.0),
		AvgPrice24h / math.Min(float64(PriceCounter), 5760.0),
	}
}
