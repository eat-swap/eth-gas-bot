package http

import (
	"eth-gas-bot/config"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func RegisterRouter() {
	Router = gin.Default()
}

func RunServer() {
	if Router == nil {
		RegisterRouter()
	}
	err := Router.Run(config.BindAddress)
	if err != nil {
		panic(err)
	}
}
