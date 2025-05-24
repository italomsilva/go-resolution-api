package usecase

import (
	"go-resolution-api/application/user/repository"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
	}
}
