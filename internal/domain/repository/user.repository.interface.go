package repository

import "go-resolution-api/internal/domain/entity"

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetById(id string) (*entity.User, error)
	GetByLogin(login string) (*entity.User, error)
	GetByDocument(document string) (*entity.User, error)
	Create(data *entity.User) (*entity.User, error)
	Update(id string, data *entity.User) (*entity.User, error)
	Delete(id string) (bool, error)
}
