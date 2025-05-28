package pkg

import (
	"go-resolution-api/internal/domain/gateway"

	"github.com/google/uuid"
)

type UUIDGateway struct{}

func NewUUIDGateway() gateway.IDGeneratorGateway {
	return &UUIDGateway{}
}

func (uuidGateway *UUIDGateway) Generate() string {
	return uuid.New().String()
}
