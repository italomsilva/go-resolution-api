package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/dto/response"
	user_dto "go-resolution-api/internal/dto/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) Login(ctx *gin.Context, input *user_dto.LoginRequest) (*entity.User, error) {
	user, _ := usecase.userRepository.GetUserByLogin(input.Login)
	if user == nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return nil, nil
	}

	comparePasswords := usecase.cryptorGateway.CheckPasswordHash(input.Password, user.Password)
	if !comparePasswords {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return nil, nil
	}

	newToken, err := usecase.tokenGateway.Generate(user.ID)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: jwt")
		return nil, err
	}
	user.Token = newToken

	return usecase.userRepository.UpdateUser(user.ID, user)
}
