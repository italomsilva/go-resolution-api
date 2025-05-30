package entity

import "time"

type Solution struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	EstimatedCost float32   `json:"estimated_cost"`
	Approved      bool      `json:"approved"`
	CreatedAt     time.Time `json:"created_at"`
	ProblemID     string    `json:"problem_id"`
	UserID        string    `json:"user_id"`
}

func NewSolution() Solution {
	return Solution{
		EstimatedCost: 0,
		Approved:      false,
		CreatedAt:     time.Now(),
	}
}
