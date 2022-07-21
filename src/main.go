package main

import (
	"eth-gas-bot/base/http"
	DaemonServices "eth-gas-bot/modules/eth-daemon/services"
)

func main() {
	go DaemonServices.Daemon()
	http.RunServer()
}
