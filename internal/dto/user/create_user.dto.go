package dto

import "go-resolution-api/internal/domain/entity"

type CreateUserRequest struct {
	Name     string              `json:"name"`
	Email    string              `json:"email"`
	Document string              `json:"document"`
	Profile  *entity.ProfileType `json:"profile"`
	Login    string              `json:"login"`
	Password string              `json:"password"`
}