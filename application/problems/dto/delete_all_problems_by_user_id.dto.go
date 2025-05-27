package dto

type ResDeleteAllProblemsByUserId struct {
	Success bool `json:"success"`
	DeletedCounter int `json:"sucess"`
}

func NewResDeleteAllProblemsByUserId() ResDeleteAllProblemsByUserId{
 return ResDeleteAllProblemsByUserId{}
}