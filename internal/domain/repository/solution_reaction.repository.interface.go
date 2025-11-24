package repository

import "go-resolution-api/internal/domain/entity"

type SolutionReactionRepository interface {
	GetByID(id string) (*entity.SolutionReaction, error)
	GetAllBySolutionId(solutionId string) ([]entity.SolutionReaction, error)
	Create(data *entity.SolutionReaction) (*entity.SolutionReaction, error)
	Delete(id string) (bool, error)
	Update(id string, reaction entity.SolutionReaction) (bool, error)
	DeleteAllBySolutionId(solutionId string) (int, error)
	GetBySolutionIdAndUserId(solutionId string, userId string) (*entity.SolutionReaction, error)
	GetReactionsBySolutionIdAndUserId(solutionId string, userId string) (int, int, entity.ReactionType, error)
}
