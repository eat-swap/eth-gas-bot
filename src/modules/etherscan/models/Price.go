package models

import "time"

type Price struct {
	Btc          float64   `json:"btc"`
	Usd          float64   `json:"usd"`
	BtcTimestamp time.Time `json:"btc_timestamp"`
	UsdTimestamp time.Time `json:"usd_timestamp"`
}
