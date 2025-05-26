package model

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Document string `json:"document"`
	Profile Profile `json:"profile"`
	Login string `json:"login"`
	Password string `json:"password"`
	Token string `json:"token"`
}