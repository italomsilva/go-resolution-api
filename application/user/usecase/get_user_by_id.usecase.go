package usecase

import (
	"go-resolution-api/application/user/model"
	"go-resolution-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) GetUserById(ctx *gin.Context, id string) (*model.User, error) {
	user, err := usecase.userRepository.GetUserById(id)
	if err != nil || user == nil {
		response.SendError(ctx, http.StatusNotFound, "User not found")
		return nil, err
	}
	return user, nil
}
