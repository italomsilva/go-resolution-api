package middleware

import (
	"go-resolution-api/internal/domain/gateway"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenGateway gateway.TokenGateway
}

func NewAuthMiddleware(tokenGateway gateway.TokenGateway) *AuthMiddleware {
	return &AuthMiddleware{tokenGateway: tokenGateway}
}

func (middleware *AuthMiddleware) Apply() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("req-token")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := middleware.tokenGateway.Validate(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set("userId", userID)
		c.Next()
	}
}
