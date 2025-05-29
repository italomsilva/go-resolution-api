package dto

type DeleteAllSolutionsByProblemIdRequest struct {
	ProblemId string `json:"problem_id"`
}

type DeleteAllSolutionsByProblemIdResponse struct {
	Success        bool `json:"success"`
	DeletedCounter int  `json:"deleted_counter"`
}
