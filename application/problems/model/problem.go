package model

import "time"

type Problem struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Status ProblemStatus`json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UserID string `json:"user_id"`
}


func NewProblem() Problem{
	return Problem{
		CreatedAt: time.Now(),
		Status: Open,
	}
}