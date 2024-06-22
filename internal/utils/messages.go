package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, op string, data interface{}, status int) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(status, gin.H{
		"message": fmt.Sprintf("operation from handler: %s, successful.", op),
		"data":    data,
	})
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
