package userDto

type ReqLogin struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func NewReqLogin() ReqLogin{
	return ReqLogin{}
}