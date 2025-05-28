package repository

import "go-resolution-api/internal/domain/entity"

type ProblemRepository interface {
	GetAllProblems() ([]entity.Problem, error)
	GetAllProblemsByUserId(userID string) ([]entity.Problem, error)
	GetProblemById(id string) (*entity.Problem, error)
	CreateProblem(data *entity.Problem) (*entity.Problem, error)
	UpdateProblem(id string, data *entity.Problem) (*entity.Problem, error)
	DeleteProblem(id string) (bool, error)
	DeleteAllProblemsByUserId(userId string) (int, error)
}
