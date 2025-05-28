package dto

import "go-resolution-api/internal/domain/entity"

type UpdateProblemRequest struct {
	ID          string                `json:"id"`
	Title       *string               `json:"title"`
	Description *string               `json:"description"`
	Location    *string               `json:"location"`
	Status      *entity.ProblemStatus `json:"status"`
}
