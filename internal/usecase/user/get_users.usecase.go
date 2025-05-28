package usecase

import "go-resolution-api/internal/domain/entity"


func (usecase *UserUseCase) GetUsers() ([]entity.User, error) {
	return usecase.userRepository.GetUsers()
}
