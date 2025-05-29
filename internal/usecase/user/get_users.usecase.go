package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type GetUsersUsecase struct {
	userRepository     repository.UserRepository
}

func NewGetUsersUsecase(
	userRepository repository.UserRepository,
) GetUsersUsecase {
	return GetUsersUsecase{
		userRepository:     userRepository,
	}
}

func (usecase *GetUsersUsecase) Execute() ([]entity.User, error) {
	return usecase.userRepository.GetUsers()
}
