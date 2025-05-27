package dto

import "go-resolution-api/application/problems/model"

type ReqUpdateProblem struct {
	ID          string               `json:"id"`
	Title       *string              `json:"title"`
	Description *string              `json:"description"`
	Location    *string              `json:"location"`
	Status      *model.ProblemStatus `json:"status"`
}

func NewReqUpdateProblem() ReqUpdateProblem {
	return ReqUpdateProblem{}
}
