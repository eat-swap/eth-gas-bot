package models

type Gas struct {
	LastBlock       int64     `json:"LastBlock"`
	SafeGasPrice    int16     `json:"SafeGasPrice"`
	ProposeGasPrice int16     `json:"ProposeGasPrice"`
	FastGasPrice    int16     `json:"FastGasPrice"`
	SuggestBaseFee  float64   `json:"suggestBaseFee"`
	GasUsedRatio    []float64 `json:"gasUsedRatio"`
}
