package dto

import "time"

type StatsCountSolutionsReactionsResponseDto struct {
	SolutionTitle string `json:"solution_title"`
	Likes         int `json:"likes"`
	Dislikes      int `json:"dislikes"`
	CreatedAt     time.Time `json:"created_at"`
}
