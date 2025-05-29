package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteUserUsecase struct {
	userRepository repository.UserRepository
	tokenGateway   gateway.TokenGateway
	loginUsecase   LoginUsecase
}

func NewDeleteUserUsecase(
	userRepository repository.UserRepository,
	tokenGateway gateway.TokenGateway,
	loginUsecase LoginUsecase,

) DeleteUserUsecase {
	return DeleteUserUsecase{
		userRepository: userRepository,
		tokenGateway:   tokenGateway,
		loginUsecase:   loginUsecase,
	}
}

func (usecase *DeleteUserUsecase) Execute(ctx *gin.Context, input *dto.DeleteUserRequest) (*dto.DeleteUserResponse, error) {
	output := dto.DeleteUserResponse{
		Success: false,
	}

	inputLogin := dto.LoginRequest{
		Login:    input.Login,
		Password: input.Password,
	}

	userLogin, err := usecase.loginUsecase.Execute(ctx, &inputLogin)
	if userLogin == nil {
		return &output, err
	}

	userId, _ := usecase.tokenGateway.GetUserId(ctx)
	if userId != userLogin.ID {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return &output, err
	}

	deleteUser, err := usecase.userRepository.DeleteUser(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "user deletion failed")
	}
	output.Success = deleteUser
	return &output, nil

}
