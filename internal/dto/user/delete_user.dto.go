package dto

type DeleteUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DeleteUserResponse struct {
	Success bool `json:"success"`
}

