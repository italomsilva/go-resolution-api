package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ApiKeyMiddleware struct{}

func NewApiKeyMiddleware() *ApiKeyMiddleware {
	return &ApiKeyMiddleware{}
}

func (middleware * ApiKeyMiddleware)Apply() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("go-api-key")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"statusCode": http.StatusUnauthorized,
					"message":    "Authorization header missing or malformed",
				})
			return
		}

		apikey := os.Getenv("API_KEY_VALUE")
		if authHeader != apikey {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"statusCode": http.StatusUnauthorized,
					"message":    "Invalid apikey",
				})
			return
		}
		ctx.Next()
	}
}
