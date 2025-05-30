package repository

import "go-resolution-api/internal/domain/entity"


type SectorRepository interface {
	GetAll() ([]entity.Sector, error)
	GetById(id int) (*entity.Sector, error)
	GetByIds(idList []int) ([]entity.Sector, error)
}
