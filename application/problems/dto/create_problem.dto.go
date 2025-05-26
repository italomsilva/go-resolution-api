package dto

import "go-resolution-api/application/problems/model"

type ReqCreateProblem struct {
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Location    string               `json:"location"`
	Status      *model.ProblemStatus `json:"status"`
}

func NewReqCreateProblem() ReqCreateProblem {
	return ReqCreateProblem{}
}
