package dto
import (
	"time"
)

type GetMySolutionsToAppResponse struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Approved     bool      `json:"approved"`
	CreatedAt    time.Time `json:"created_at"`
	ProblemID    string    `json:"problem_id"`
	ProblemTitle string    `json:"problem_title"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
}