package userDto

import "go-resolution-api/model"

type ReqCreateUser struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Document string `json:"document"`
	Profile model.Profile `json:"profile"`
	Login string `json:"login"`
	Password string `json:"password"`
	Token string `json:"token"`
}

func NewReqCreateUser() ReqCreateUser {
	return ReqCreateUser{
		Profile: model.Individual,
		Token:   "",  
	}
}