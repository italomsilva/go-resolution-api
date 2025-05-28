package dto

type DeleteProblemRequest struct {
	ID string `json:"id"`
}

type DeleteProblemResponse struct {
	Success bool `json:"success"`
}

