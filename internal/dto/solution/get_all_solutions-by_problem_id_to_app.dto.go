package dto

import (
	"go-resolution-api/internal/domain/entity"
	"time"
)

// class GetAllSolutionsResponseDto
//   final String id;
//   final String title;
//   final String description;
//   final double estimatedCost;
//   final bool approved;
//   final DateTime createdAt;
//   final String problemId;
//   final String problemTitle;
//   final String userId;
//   final String userLogin;
//   final int likes;
//   final int dislikes;
//   SolutionReaction myReaction;

type GetAllSolutionsByProblemIdToAAppResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	EstimatedCost float64   `json:"estimated_cost"`
	Approved      bool      `json:"approved"`
	CreatedAt     time.Time `json:"created_at"`
	ProblemID     string    `json:"problem_id"`
	ProblemTitle  string    `json:"problem_title"`
	UserID        string    `json:"user_id"`
	UserLogin     string    `json:"user_login"`
	Likes         int       `json:"likes"`
	Dislikes      int       `json:"dislikes"`
	MyReaction    entity.ReactionType    `json:"my_reaction"`
}