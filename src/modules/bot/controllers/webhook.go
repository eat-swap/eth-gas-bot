package controllers

import (
	"eth-gas-bot/modules/bot/services"
	"eth-gas-bot/modules/telegram/models/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Webhook(c *gin.Context) {
	var u entities.Update
	err := c.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	go services.HandleIncomingMessage(u.ExtractMessage(), c)

	c.JSON(http.StatusOK, gin.H{})
}
