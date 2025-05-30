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
	tokenGateway       gateway.TokenGateway
	createUserUsecase  usecase.CreateUserUsecase
	deleteUserUsecase  usecase.DeleteUserUsecase
	getUserByIdUsecase usecase.GetUserByIdUsecase
	getUsersUsecase    usecase.GetUsersUsecase
	loginUsecase       usecase.LoginUsecase
	updateUserUsecase  usecase.UpdateUserUsecase
}

func NewUserController(
	tokenGateway gateway.TokenGateway,
	createUserUsecase usecase.CreateUserUsecase,
	deleteUserUsecase usecase.DeleteUserUsecase,
	getUserByIdUsecase usecase.GetUserByIdUsecase,
	getUsersUsecase usecase.GetUsersUsecase,
	loginUsecase usecase.LoginUsecase,
	updateUserUsecase usecase.UpdateUserUsecase,

) UserController {
	return UserController{
		tokenGateway: tokenGateway,
		createUserUsecase: createUserUsecase,
		deleteUserUsecase: deleteUserUsecase,
		getUserByIdUsecase: getUserByIdUsecase,
		getUsersUsecase: getUsersUsecase,
		loginUsecase: loginUsecase,
		updateUserUsecase: updateUserUsecase,
	}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	result, _ := controller.getUsersUsecase.Execute()
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("userId")
	result, _ := controller.getUserByIdUsecase.Execute(ctx, id)
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

	result, _ := controller.createUserUsecase.Execute(ctx, &input)
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

	result, _ := controller.loginUsecase.Execute(ctx, &input)
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

	result, _ := controller.updateUserUsecase.Execute(ctx, &input)
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

	result, _ := controller.deleteUserUsecase.Execute(ctx, &input)
	if result.Success {
		response.SendSucess(ctx, http.StatusOK, result, "acount deleted successfully")
	}
}
