package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) GetUserById(ctx *gin.Context, id string) (*entity.User, error) {
	user, err := usecase.userRepository.GetUserById(id)
	if err != nil || user == nil {
		response.SendError(ctx, http.StatusNotFound, "User not found")
		return nil, err
	}
	return user, nil
}
