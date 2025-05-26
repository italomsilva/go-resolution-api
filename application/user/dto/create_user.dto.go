package dto

import "go-resolution-api/application/user/model"


type ReqCreateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Document string `json:"document"`
	Profile *model.Profile `json:"profile"`
	Login string `json:"login"`
	Password string `json:"password"`
}

func NewReqCreateUser() ReqCreateUser {
	return ReqCreateUser{}
}