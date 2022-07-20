package main

import (
	"encoding/json"
	"eth-gas-bot/base/http"
	"eth-gas-bot/modules/etherscan/services"
	"fmt"
)

func main() {
	// debug()

	http.RunServer()
}

func debug() {
	gas, err := services.GetGas()
	if err != nil {
		panic(err)
	}
	s, _ := json.Marshal(gas)
	fmt.Println(string(s))
}
