package dto

type DeleteAllProblemsByUserIdResponse struct {
	Success        bool `json:"success"`
	DeletedCounter int  `json:"sucess"`
}
