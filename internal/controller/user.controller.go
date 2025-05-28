package controller

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/user"
	usecase "go-resolution-api/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase  usecase.UserUseCase
	tokenGateway gateway.TokenGateway
}

func NewUserController(useCase usecase.UserUseCase, tokenGateway gateway.TokenGateway) UserController {
	return UserController{
		userUseCase: useCase,
		tokenGateway: tokenGateway,
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
	input := dto.CreateUserRequest{}
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
	input := dto.LoginRequest{}
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
	input := dto.UpdateUserRequest{}
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
	input := dto.DeleteUserRequest{}
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
