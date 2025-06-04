package infra

import (
	"errors"
	"go-resolution-api/internal/domain/gateway"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthJWTGateway struct {
	secret []byte
}

func NewAuthJWTGateway() gateway.TokenGateway {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	return &AuthJWTGateway{secret: []byte(secret)}
}

func (authJWTGateway *AuthJWTGateway) Generate(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(authJWTGateway.secret)
}

func (authJWTGateway *AuthJWTGateway) Validate(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return authJWTGateway.secret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token")
	}

	return userID, nil
}

func (authJWTGateway *AuthJWTGateway) GetUserId(ctx *gin.Context) (string, bool) {
	tokenHeader := ctx.GetHeader("req-token")
	if tokenHeader == "" {
		return "", false
	}
	token := tokenHeader[7:]
	userId, err := authJWTGateway.Validate(token)
	if err != nil || userId == "" {
		println(err.Error())
		return "", false
	}
	return userId, true

}

