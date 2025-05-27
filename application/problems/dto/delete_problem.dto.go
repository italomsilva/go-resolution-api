package dto

type ReqDeleteProblem struct {
	ID string `json:"id"`
}

func NewReqDeleteProblem() ReqDeleteProblem{
 return ReqDeleteProblem{}
}

type ResDeleteProblem struct {
	Success bool `json:"success"`
}

func NewResDeleteProblem() ResDeleteProblem{
 return ResDeleteProblem{}
}