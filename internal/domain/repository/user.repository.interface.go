package repository

import "go-resolution-api/internal/domain/entity"

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUserById(id string) (*entity.User, error)
	GetUserByLogin(login string) (*entity.User, error)
	GetUserByDocument(document string) (*entity.User, error)
	CreateUser(data *entity.User) (*entity.User, error)
	UpdateUser(id string, data *entity.User) (*entity.User, error)
	DeleteUser(id string) (bool, error)
}
