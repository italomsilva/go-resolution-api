package repository

import "go-resolution-api/internal/domain/entity"

type ProblemSectorRepository interface {
	GetAll() ([]entity.ProblemSector, error)
	GetAllByProblemId(problemID string) ([]entity.ProblemSector, error)
	GetById(id int) (*entity.ProblemSector, error)
	Create(data *entity.ProblemSector) (*entity.ProblemSector, error)
	Delete(id int) (bool, error)
	DeleteAllByProblemId(problemID string) (int, error)
}
