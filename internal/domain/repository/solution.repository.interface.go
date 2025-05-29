package repository

import "go-resolution-api/internal/domain/entity"

type SolutionRepository interface {
	GetAllSolutionsByProblemId(problemId string) ([]entity.Solution, error)
	GetSolutionById(id string) (*entity.Solution, error)
	CreateSolution(data *entity.Solution) (*entity.Solution, error)
}