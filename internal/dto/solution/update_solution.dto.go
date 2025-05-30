package dto

type UpdateSolutionRequest struct {
	ID            *string    `json:"id"`
	Title         *string    `json:"title"`
	Description   *string    `json:"description"`
	EstimatedCost *float32   `json:"estimated_cost"`
}
