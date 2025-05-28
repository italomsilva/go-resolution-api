package dto

import "go-resolution-api/internal/domain/entity"

type CreateProblemRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Location    string                `json:"location"`
	Status      *entity.ProblemStatus `json:"status"`
}
