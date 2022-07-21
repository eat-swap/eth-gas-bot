package http

import (
	"eth-gas-bot/config"
	BotControllers "eth-gas-bot/modules/bot/controllers"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func RegisterRouter() {
	Router = gin.Default()
	Router.TrustedPlatform = gin.PlatformCloudflare

	BotControllers.Register(Router)
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
