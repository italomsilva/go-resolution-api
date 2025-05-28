package gateway

import "github.com/gin-gonic/gin"

type TokenGateway interface {
	Generate(userID string) (string, error)
	Validate(tokenString string) (string, error)
	GetUserId(ctx *gin.Context) (string, bool)
}
