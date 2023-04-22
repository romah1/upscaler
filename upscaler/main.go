package main

import (
	"github.com/gin-gonic/gin"
	"upscaler/upscaler/controllers"
)

func main() {
	engine := setupGinEngine()

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func setupGinEngine() *gin.Engine {
	engine := gin.Default()
	api := engine.Group("/api")

	// upscale
	api.POST("/upscale", controllers.Upscale)

	return engine
}
