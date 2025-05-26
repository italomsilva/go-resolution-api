package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserIdByToken(ctx *gin.Context) (string, bool) {
	userIdToken, exists := ctx.Get("userId")
	userId := fmt.Sprintf("%v", userIdToken)
	if !exists || userId == "" {
		return "", false
	}

	return userId, true

}