package usecase

import (
	"go-resolution-api/application/user/dto"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) DeleteUser(ctx *gin.Context, input *dto.ReqDeleteUser) (*dto.ResDeleteUser, error) {
	responseDelete := dto.NewResDeleteUser()
	responseDelete.Success = false

	inputLogin := dto.NewReqLogin()
	inputLogin.Login = input.Login
	inputLogin.Password = input.Password

	userLogin, err := usecase.Login(ctx, &inputLogin)
	if userLogin == nil {
		return &responseDelete, err
	}

	userId, _ := utils.GetUserId(ctx)
	if userId != userLogin.ID {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return &responseDelete, err
	}

	deleteUser, err := usecase.userRepository.DeleteUser(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "user deletion failed")
	}
	responseDelete.Success = deleteUser
	return &responseDelete, nil

}
