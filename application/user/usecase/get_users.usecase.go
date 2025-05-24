package usecase

import "go-resolution-api/application/user/model"

func (usecase *UserUseCase) GetUsers() ([]model.User, error) {
	return usecase.userRepository.GetUsers()
}
