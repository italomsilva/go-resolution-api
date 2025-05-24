package controller

import (
	"go-resolution-api/application/user/dto"
	"go-resolution-api/application/user/model"
	"go-resolution-api/application/user/usecase"
	"go-resolution-api/response"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return UserController{
		userUseCase: useCase,
	}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	result, _ := controller.userUseCase.GetUsers()
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, _ := controller.userUseCase.GetUserById(ctx, id)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	input := dto.NewReqCreateUser()
	err := ctx.BindJSON(&input)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, _ := controller.userUseCase.CreateUser(ctx, &input)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "signUp successfully")
	}
}

func (controller *UserController) Login(ctx *gin.Context) {
	input := dto.NewReqLogin()
	err := ctx.BindJSON(&input)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, _ := controller.userUseCase.Login(ctx, &input)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "signIn successfully")
	}
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	input := dto.NewReqUpdateUser(&model.User{})
	err := ctx.BindJSON(&input)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, _ := controller.userUseCase.UpdateUser(ctx, &input)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "user updated successfully")
	}
}

func (controller *UserController) DeleteAccount(ctx *gin.Context) {
	input := dto.NewReqDeleteUser()
	err := ctx.BindJSON(&input)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, _ := controller.userUseCase.DeleteUser(ctx, &input)
	if result.Success {
		response.SendSucess(ctx, http.StatusOK, result, "acount deleted successfully")
	}
}
