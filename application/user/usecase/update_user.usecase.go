package usecase

import (
	"fmt"
	"go-resolution-api/application/user/dto"
	"go-resolution-api/application/user/model"
	"go-resolution-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) UpdateUser(ctx *gin.Context, input *dto.ReqUpdateUser) (*model.User, error) {
	userIdToken, exists := ctx.Get("userId")
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required.")
		return nil, nil
	}
	userId := fmt.Sprintf("%v", userIdToken)

	user, err := usecase.userRepository.GetUserById(userId)
	if user == nil {
		response.SendError(ctx, http.StatusNotFound, "User Not Found")
		return nil, err
	}

	if input.Name == nil ||*input.Name != ""  {
		user.Name = *input.Name
	}

	if input.Login == nil || *input.Login != "" {
		foundUserByLogin, _ := usecase.userRepository.GetUserByLogin(*input.Login)
		if foundUserByLogin != nil {
			response.SendError(ctx, http.StatusConflict, "Login already exists")
			return nil, nil
		}
		user.Login = *input.Login
	}
	return usecase.userRepository.UpdateUser(userId, user)
}
