package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserUsecase struct {
	userRepository     repository.UserRepository
	tokenGateway       gateway.TokenGateway
}

func NewUpdateUserUsecase(
	userRepository repository.UserRepository,
	tokenGateway gateway.TokenGateway,
) UpdateUserUsecase {
	return UpdateUserUsecase{
		userRepository:     userRepository,
		tokenGateway:       tokenGateway,
	}
}


func (usecase *UpdateUserUsecase) Execute(ctx *gin.Context, input *dto.UpdateUserRequest) (*entity.User, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)
	
	user, err := usecase.userRepository.GetById(userId)
	if user == nil {
		response.SendError(ctx, http.StatusNotFound, "User Not Found")
		return nil, err
	}

	if input.Name == nil || *input.Name != "" {
		user.Name = *input.Name
	}

	if input.Login == nil || *input.Login != "" {
		foundUserByLogin, _ := usecase.userRepository.GetByLogin(*input.Login)
		if foundUserByLogin != nil {
			response.SendError(ctx, http.StatusConflict, "Login already exists")
			return nil, nil
		}
		user.Login = *input.Login
	}
	return usecase.userRepository.Update(userId, user)
}
