package dto

type DeleteAllSolutionsByUserIdResponse struct {
	Success        bool `json:"success"`
	DeletedCounter int  `json:"deleted_counter"`
}
