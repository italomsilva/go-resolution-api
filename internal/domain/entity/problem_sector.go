package entity

type ProblemSector struct {
	ID        int    `json:"id"`
	ProblemID string `json:"problem_id"`
	SectorID  int    `json:"sector_id"`
}
