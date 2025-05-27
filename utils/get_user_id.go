package utils

import (
	"go-resolution-api/gateway"

	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) (string, bool) {
	tokenHeader := ctx.GetHeader("req-token")
	if tokenHeader == "" {
		return "", false
	}
	token := tokenHeader[7:]
	userId, err := gateway.ValidateJWT(token)
	if err != nil || userId == "" {
		println(err.Error())
		return "", false
	}
	return userId, true

}
