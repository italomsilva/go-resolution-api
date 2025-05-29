package dto

type DeleteSolutionRequest struct {
	ID string `json:"id"`
}

type DeleteSolutionResponse struct {
	Success bool `json:"success"`
}
