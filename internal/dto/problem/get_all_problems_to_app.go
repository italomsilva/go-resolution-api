package dto

import (
	"go-resolution-api/internal/domain/entity"
	"time"
)


type GetAllProblemsToAppResponse struct {
	ID             string               `json:"id"`
	Title          string               `json:"title"`
	Description    string               `json:"description"`
	Location       string               `json:"location"`
	Status         entity.ProblemStatus `json:"status"`
	CreatedAt      time.Time            `json:"created_at"`
	UserID         string               `json:"user_id"`
	UserLogin      string               `json:"user_login"`
	SolutionsCount int                  `json:"solutions_count"`
}
