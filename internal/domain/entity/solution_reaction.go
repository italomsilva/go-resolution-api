package entity

type SolutionReaction struct {
	ID           string       `json:"id"`
	UserID       string       `json:"user_id"`
	SolutionID   string       `json:"solution_id"`
	ReactionType ReactionType `json:"reaction_type"`
}
