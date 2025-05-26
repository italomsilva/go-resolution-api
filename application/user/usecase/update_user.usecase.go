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
	userId, exists := ctx.Get("userId")
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required.")
		return nil, nil
	}
	userIdStr := fmt.Sprintf("%v", userId)
	foundUserById, err := usecase.userRepository.GetUserById(userIdStr)
	if foundUserById == nil {
		response.SendError(ctx, http.StatusNotFound, "User Not Found")
		return nil, err
	}

	userToUpdate := dto.NewReqUpdateUser(foundUserById)

	if *input.Name != "" {
		userToUpdate.Name = input.Name
	}

	if *input.Login != "" {
		foundUserByLogin, _ := usecase.userRepository.GetUserByLogin(*input.Login)
		if foundUserByLogin != nil {
			response.SendError(ctx, http.StatusConflict, "Login already exists")
			return nil, nil
		}
		userToUpdate.Login = input.Login
	}
	return usecase.userRepository.UpdateUser(userIdStr, &userToUpdate)
}
