package repository

import "go-resolution-api/internal/domain/entity"

type SolutionRepository interface {
	GetAllByProblemId(problemId string) ([]entity.Solution, error)
	GetAllByUserId(userId string) ([]entity.Solution, error)
	GetById(id string) (*entity.Solution, error)
	Create(data *entity.Solution) (*entity.Solution, error)
	Delete(id string) (bool, error)
	DeleteAllByProblemId(problemId string) (int, error)
	DeleteAllByUserId(userId string) (int, error)
	Update(id string, data *entity.Solution) (*entity.Solution, error)
}