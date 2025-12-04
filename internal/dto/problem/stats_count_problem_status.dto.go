package dto

import "go-resolution-api/internal/domain/entity"

type StatsCountProblemStatusResponse struct {
	Count  int                  `json:"count"`
	Status entity.ProblemStatus `json:"status"`
}
