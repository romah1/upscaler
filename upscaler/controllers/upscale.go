package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upscale(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": "some result"})
}
