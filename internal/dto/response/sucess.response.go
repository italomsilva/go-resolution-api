package response

import "github.com/gin-gonic/gin"

func SendSucess(ctx *gin.Context, statusCode int, data any, message string) bool {
	if message == "" {
		message = "successful operation"
	}
	ctx.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
	})
	return true
}
