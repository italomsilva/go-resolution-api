package pkg

import (
	"go-resolution-api/internal/domain/gateway"

	"golang.org/x/crypto/bcrypt"
)

type BcryptGateway struct{}

func NewBcryptGateway() gateway.CryptorGateway {
	return &BcryptGateway{}
}

func (cryptorGateway *BcryptGateway) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (cryptorGateway *BcryptGateway) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
