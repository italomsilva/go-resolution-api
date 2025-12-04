package repository

import (
	"go-resolution-api/internal/domain/entity"
	dto "go-resolution-api/internal/dto/problem"
)

type ProblemRepository interface {
	GetAll() ([]entity.Problem, error)
	GetAllByUserId(userID string) ([]entity.Problem, error)
	GetById(id string) (*entity.Problem, error)
	Create(data *entity.Problem) (*entity.Problem, error)
	Update(id string, data *entity.Problem) (*entity.Problem, error)
	Delete(id string) (bool, error)
	DeleteAllByUserId(userId string) (int, error)
	GetStatsCountProblemStatus(userID string) ([]dto.StatsCountProblemStatusResponse, error)
}
