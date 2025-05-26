package dto


type ReqUpdateUser struct {
	Name     *string `json:"name"`
	Login    *string `json:"login"`
	Password *string `json:"password"`
}

func NewReqUpdateUser() ReqUpdateUser {
	return ReqUpdateUser{}
}
