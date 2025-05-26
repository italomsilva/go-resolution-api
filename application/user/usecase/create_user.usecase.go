package usecase

import (
	"go-resolution-api/application/user/dto"
	"go-resolution-api/application/user/model"
	"go-resolution-api/gateway"
	"go-resolution-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) CreateUser(ctx *gin.Context, input *dto.ReqCreateUser) (*model.User, error) {
	createUser := model.NewUser()

	foundUserLogin, _ := usecase.userRepository.GetUserByLogin(input.Login)
	if foundUserLogin != nil {
		response.SendError(ctx, http.StatusConflict, "Login already in use")
		return nil, nil
	}
	createUser.Login = input.Login

	foundUserDocument, _ := usecase.userRepository.GetUserByDocument(input.Document)
	if foundUserDocument != nil {
		response.SendError(ctx, http.StatusConflict, "Document already in use")
		return nil, nil
	}
	createUser.Document = input.Document

	createUser.ID = gateway.GenerateUUID()

	hashedPassword, err := gateway.HashPassword(input.Password)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: hash")
		return nil, err
	}
	createUser.Password = hashedPassword

	token, err := gateway.GenerateJWT(createUser.ID)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: jwt")
		return nil, err
	}
	createUser.Token = token

	newUser, err := usecase.userRepository.CreateUser(&createUser)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "SignUp Error")
		return nil, err
	}
	return newUser, nil
}
