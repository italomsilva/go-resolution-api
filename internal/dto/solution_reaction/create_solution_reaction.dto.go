package dto

import "go-resolution-api/internal/domain/entity"

type CreateSolutionReactionRequest struct {
	SolutionID   string              `json:"solution_id"`
	ReactionType entity.ReactionType `json:"reaction_type"`
}
