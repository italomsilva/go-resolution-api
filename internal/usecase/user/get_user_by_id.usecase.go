package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserByIdUsecase struct {
	userRepository     repository.UserRepository
}

func NewGetUserByIdUsecase(
	userRepository repository.UserRepository,
) GetUserByIdUsecase {
	return GetUserByIdUsecase{
		userRepository:     userRepository,
	}
}


func (usecase *GetUserByIdUsecase) Execute(ctx *gin.Context, id string) (*entity.User, error) {
	user, err := usecase.userRepository.GetUserById(id)
	if err != nil || user == nil {
		response.SendError(ctx, http.StatusNotFound, "User not found")
		return nil, err
	}
	return user, nil
}
