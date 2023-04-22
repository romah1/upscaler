package tg_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"upscaler/tg/tg_bot"
)

func SetupGinEngine(bot *tg_bot.Bot) *gin.Engine {
	engine := gin.Default()
	api := engine.Group("/api")

	api.POST("/upscaling_failed", func(context *gin.Context) {
		var body Error
		if err := context.ShouldBindJSON(&body); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		err := bot.SendMessage(body.ChatID, body.Reason)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, "success")
	})

	return engine
}
