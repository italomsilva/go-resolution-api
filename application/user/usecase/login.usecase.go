package usecase

import (
	"go-resolution-api/application/user/dto"
	"go-resolution-api/application/user/model"
	"go-resolution-api/gateway"
	"go-resolution-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) Login(ctx *gin.Context, input *dto.ReqLogin) (*model.User, error) {
	user, _ := usecase.userRepository.GetUserByLogin(input.Login)
	if user == nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return nil, nil
	}

	comparePasswords := gateway.CheckPasswordHash(input.Password, user.Password)
	if !comparePasswords {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return nil, nil
	}

	newToken, err := gateway.GenerateJWT(user.ID)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: jwt")
		return nil, nil
	}

	user.Token = newToken
	updatedUser := dto.NewReqUpdateUser(user)
	return usecase.userRepository.UpdateUser(user.ID, &updatedUser)

}
