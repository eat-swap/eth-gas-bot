package services

import (
	"eth-gas-bot/modules/eth-daemon/global"
	"eth-gas-bot/modules/eth-daemon/models"
	"eth-gas-bot/modules/etherscan/services"
	"log"
	"time"
)

func Daemon() {
	for {
		go gas()
		go price()
		log.Printf("Next scheduled update at %s", time.Now().Add(global.RefreshInterval).Format(time.RFC1123))
		time.Sleep(global.RefreshInterval)
	}
}

func gas() {
	gas, err := services.GetGas()
	if err != nil {
		log.Printf("Cannot update gas info: %s", err.Error())
		return
	}
	global.GasMutex.Lock()
	global.GasHead = (global.GasHead + 1) % global.CacheLimit
	global.GasCounter++
	global.HistoryGas[global.GasHead] = models.GasInfo{
		Valid:      true,
		Gas:        gas,
		ObtainedAt: time.Now(),
	}
	global.GasMutex.Unlock()

	// Print gas info
	log.Printf("Successfully obtained gas info. Base price: %.6f", gas.SuggestBaseFee)
}

func price() {
	price, err := services.GetPrice()
	if err != nil {
		log.Printf("Cannot update price info: %s", err.Error())
		return
	}
	global.PriceMutex.Lock()
	global.PriceHead = (global.PriceHead + 1) % global.CacheLimit
	global.PriceCounter++
	global.HistoryPrice[global.PriceHead] = models.PriceInfo{
		Valid:      true,
		Price:      price,
		ObtainedAt: time.Now(),
	}
	global.PriceMutex.Unlock()

	// Print price info
	log.Printf("Successfully obtained price info. Price: %.3f", price.Usd)
}
