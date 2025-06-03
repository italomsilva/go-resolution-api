package dto

type DeleteProblemSectorRequest struct {
	ID int `json:"id"`
}


type DeleteProblemSectorResponse struct {
	Success bool `json:"success"`
}