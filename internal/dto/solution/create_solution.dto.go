package dto


type CreateSolutionRequest struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Estimated_cost float32   `json:"estimated_cost"`
	ProblemId      string    `json:"problem_id"`
}
