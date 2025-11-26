package dto

type ApproveSolutionRequest struct {
	ProblemID  string `json:"problem_id"`
	SolutionID string `json:"solution_id"`
}