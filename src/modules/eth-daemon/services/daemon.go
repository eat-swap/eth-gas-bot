package services

import (
	"eth-gas-bot/config"
	"eth-gas-bot/modules/eth-daemon/global"
	"eth-gas-bot/modules/eth-daemon/models"
	"eth-gas-bot/modules/etherscan/services"
	"eth-gas-bot/modules/telegram/models/params"
	TelegramServices "eth-gas-bot/modules/telegram/services"
	"eth-gas-bot/utils"
	"fmt"
	"log"
	"time"
)

func Daemon() {
	defer Daemon()
	gas()
	price()
	for {
		time.Sleep(global.RefreshInterval)
		go gas()
		go price()
		global.NextUpdate = time.Now().Add(global.RefreshInterval)
		log.Printf("Next scheduled update at %s", global.NextUpdate.Format(time.RFC1123))
		go monitor()
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

	global.AvgGasMutex.Lock()
	global.GasMutex.RLock()
	defer global.GasMutex.RUnlock()

	global.AvgGas5Min.Base += gas.SuggestBaseFee
	global.AvgGas5Min.Low += float64(gas.SafeGasPrice)
	global.AvgGas5Min.Avg += float64(gas.ProposeGasPrice)
	global.AvgGas5Min.High += float64(gas.FastGasPrice)
	global.AvgGas5Min.Count++
	if global.AvgGas5Min.Count > 20 {
		global.AvgGas5Min.Base -= global.HistoryGas[(global.GasHead+global.CacheLimit-20)%global.CacheLimit].Gas.SuggestBaseFee
		global.AvgGas5Min.Low -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-20)%global.CacheLimit].Gas.SafeGasPrice)
		global.AvgGas5Min.Avg -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-20)%global.CacheLimit].Gas.ProposeGasPrice)
		global.AvgGas5Min.High -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-20)%global.CacheLimit].Gas.FastGasPrice)
		global.AvgGas5Min.Count--
	}

	global.AvgGas1h.Base += gas.SuggestBaseFee
	global.AvgGas1h.Low += float64(gas.SafeGasPrice)
	global.AvgGas1h.Avg += float64(gas.ProposeGasPrice)
	global.AvgGas1h.High += float64(gas.FastGasPrice)
	global.AvgGas1h.Count++
	if global.AvgGas1h.Count > 240 {
		global.AvgGas1h.Base -= global.HistoryGas[(global.GasHead+global.CacheLimit-240)%global.CacheLimit].Gas.SuggestBaseFee
		global.AvgGas1h.Low -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-240)%global.CacheLimit].Gas.SafeGasPrice)
		global.AvgGas1h.Avg -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-240)%global.CacheLimit].Gas.ProposeGasPrice)
		global.AvgGas1h.High -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-240)%global.CacheLimit].Gas.FastGasPrice)
		global.AvgGas1h.Count--
	}

	global.AvgGas6h.Base += gas.SuggestBaseFee
	global.AvgGas6h.Low += float64(gas.SafeGasPrice)
	global.AvgGas6h.Avg += float64(gas.ProposeGasPrice)
	global.AvgGas6h.High += float64(gas.FastGasPrice)
	global.AvgGas6h.Count++
	if global.AvgGas6h.Count > 1440 {
		global.AvgGas6h.Base -= global.HistoryGas[(global.GasHead+global.CacheLimit-1440)%global.CacheLimit].Gas.SuggestBaseFee
		global.AvgGas6h.Low -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-1440)%global.CacheLimit].Gas.SafeGasPrice)
		global.AvgGas6h.Avg -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-1440)%global.CacheLimit].Gas.ProposeGasPrice)
		global.AvgGas6h.High -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-1440)%global.CacheLimit].Gas.FastGasPrice)
		global.AvgGas6h.Count--
	}

	global.AvgGas24h.Base += gas.SuggestBaseFee
	global.AvgGas24h.Low += float64(gas.SafeGasPrice)
	global.AvgGas24h.Avg += float64(gas.ProposeGasPrice)
	global.AvgGas24h.High += float64(gas.FastGasPrice)
	global.AvgGas24h.Count++
	if global.AvgGas24h.Count > 5760 {
		global.AvgGas24h.Base -= global.HistoryGas[(global.GasHead+global.CacheLimit-5760)%global.CacheLimit].Gas.SuggestBaseFee
		global.AvgGas24h.Low -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-5760)%global.CacheLimit].Gas.SafeGasPrice)
		global.AvgGas24h.Avg -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-5760)%global.CacheLimit].Gas.ProposeGasPrice)
		global.AvgGas24h.High -= float64(global.HistoryGas[(global.GasHead+global.CacheLimit-5760)%global.CacheLimit].Gas.FastGasPrice)
		global.AvgGas24h.Count--
	}

	global.AvgGasMutex.Unlock()

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

	global.AvgPriceMutex.Lock()
	global.PriceMutex.RLock()
	defer global.PriceMutex.RUnlock()

	global.AvgPrice5Min += price.Usd
	global.AvgPrice1h += price.Usd
	global.AvgPrice6h += price.Usd
	global.AvgPrice24h += price.Usd
	if global.PriceCounter > 20 {
		global.AvgPrice5Min -= global.HistoryPrice[(global.PriceHead+global.CacheLimit-20)%global.CacheLimit].Price.Usd
	}
	if global.PriceCounter > 240 {
		global.AvgPrice1h -= global.HistoryPrice[(global.PriceHead+global.CacheLimit-240)%global.CacheLimit].Price.Usd
	}
	if global.PriceCounter > 1440 {
		global.AvgPrice6h -= global.HistoryPrice[(global.PriceHead+global.CacheLimit-1440)%global.CacheLimit].Price.Usd
	}
	if global.PriceCounter > 5760 {
		global.AvgPrice24h -= global.HistoryPrice[(global.PriceHead+global.CacheLimit-5760)%global.CacheLimit].Price.Usd
	}
	global.AvgPriceMutex.Unlock()

	// Print price info
	log.Printf("Successfully obtained price info. Price: %.3f", price.Usd)
}

const (
	LowFeeThreshold = 0.3
)

var (
	previousLowFee = -LowFeeThreshold
)

func monitor() {
	gas := global.GetCurrentGas().Gas
	price := global.GetCurrentPrice().Price.Usd

	// Transfer ETH
	currentPrice := float64(gas.SafeGasPrice) * price * 1e-9 * 21000

	if (currentPrice-LowFeeThreshold)*previousLowFee >= 0 {
		return
	}

	var text string
	if currentPrice <= LowFeeThreshold {
		text = fmt.Sprintf("✅ Transferring price <= $%.3f, current *$%.3f*", LowFeeThreshold, currentPrice)
	} else {
		text = fmt.Sprintf("❌ Transferring price > $%.3f, current *$%.3f*", LowFeeThreshold, currentPrice)
	}
	previousLowFee = currentPrice - LowFeeThreshold
	_, err := TelegramServices.SendMessage(config.TelegramBotToken, &params.SendMessageParams{
		ChatId:           0,
		Text:             utils.WrapForMarkdownWorse(text),
		ParseMode:        params.MessageParseModeMarkdown,
		ProtectedContent: true,
	})
	if err != nil {
		log.Printf("Cannot send monitor message: %s", err.Error())
	}
}
