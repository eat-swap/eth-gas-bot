package controllers

import (
	"eth-gas-bot/config"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST(config.WebhookEntry, Webhook)
}
