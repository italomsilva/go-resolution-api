package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
)

type UserUseCase struct {
	userRepository     repository.UserRepository
	tokenGateway       gateway.TokenGateway
	cryptorGateway     gateway.CryptorGateway
	idGeneratorGateway gateway.IDGeneratorGateway
}

func NewUserUseCase(
	userRepository repository.UserRepository,
	tokenGateway gateway.TokenGateway,
	cryptorGateway gateway.CryptorGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,
) UserUseCase {
	return UserUseCase{
		userRepository:     userRepository,
		tokenGateway:       tokenGateway,
		cryptorGateway:     cryptorGateway,
		idGeneratorGateway: idGeneratorGateway,
	}
}
