package userDto

import "go-resolution-api/model"

type ReqUpdateUser struct {
	Name     *string `json:"name"`
	Login    *string `json:"login"`
	Password *string `json:"password"`
	Token    *string `json:"token"`
}

func NewReqUpdateUser(user *model.User) ReqUpdateUser {
	return ReqUpdateUser{
		Name: &user.Name,
		Login: &user.Login,
		Password: &user.Password,
		Token: &user.Token,
	}
}
