package dto

type DeleteSolutionReactionRequest struct {
	ID string `json:"id"`
}

type DeleteSolutionReactionResponse struct {
	Success bool `json:"success"`
}
