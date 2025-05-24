package usecase

import (
	"fmt"
	"go-resolution-api/application/user/dto"
	"go-resolution-api/response"
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

	userIdToken, _ := ctx.Get("userId")
	userId := fmt.Sprintf("%v", userIdToken)
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
