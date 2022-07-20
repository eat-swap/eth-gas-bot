package response

type GeneralResponse[T ConcreteResult] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}

type ConcreteResult interface {
	GasInfoResult | EthPriceResult
}
