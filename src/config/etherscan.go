package config

import "os"

var (
	EtherScanApiKey string
)

func init() {
	EtherScanApiKey = os.Getenv("ETHERSCAN_API_KEY")
}
