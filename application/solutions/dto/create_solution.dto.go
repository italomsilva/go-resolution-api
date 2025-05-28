package dto

import "time"

type ReqCreateSolution struct {
	Title         string    `json:"title"`
	Description   *string    `json:"description"`
	EstimatedCost *float32   `json:"estimated_cost"`
	CreatedAt     *time.Time `json:"created_at"`
	ProblemId     *string    `json:"problem_id"`
}

func NewReqCreateSolution() ReqCreateSolution {
	return ReqCreateSolution{}
}
