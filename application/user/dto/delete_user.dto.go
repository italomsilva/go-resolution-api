package dto

type ReqDeleteUser struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func NewReqDeleteUser() ReqDeleteUser{
	return ReqDeleteUser{}
}

type ResDeleteUser struct {
	Success bool `json:"success"`
}

func NewResDeleteUser() ResDeleteUser{
	return ResDeleteUser{}
}

