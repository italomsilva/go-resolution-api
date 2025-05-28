package model

import "time"

type Solution struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	EstimatedCost float32 `json:"estimated_cost"`
	Approved bool `json:"approved"`
	CreatedAt time.Time `json:"created_at"`
	ProblemId string `json:"problem_id"`
	UserId string `json:"user_id"`
}

func NewSolution() Solution{
	return Solution{
		EstimatedCost: 0,
		Approved: false,
		CreatedAt: time.Now(),
	}
}