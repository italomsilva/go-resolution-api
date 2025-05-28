package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) UpdateUser(ctx *gin.Context, input *dto.UpdateUserRequest) (*entity.User, error) {
	userId, exists := usecase.tokenGateway.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required.")
		return nil, nil
	}

	user, err := usecase.userRepository.GetUserById(userId)
	if user == nil {
		response.SendError(ctx, http.StatusNotFound, "User Not Found")
		return nil, err
	}

	if input.Name == nil || *input.Name != "" {
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
