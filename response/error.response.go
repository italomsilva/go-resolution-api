package response

import "github.com/gin-gonic/gin"

func SendError(ctx *gin.Context, statusCode int, message string) bool {
	if message == "" {
		message = "operation failed"
	}
	ctx.JSON(statusCode, gin.H{
		"StatusCode": statusCode,
		"message":    message,
	})
	return true
}
