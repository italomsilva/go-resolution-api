package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *UserUseCase) CreateUser(ctx *gin.Context, input *dto.CreateUserRequest) (*entity.User, error) {
	createUser := entity.NewUser()

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

	createUser.ID = usecase.idGeneratorGateway.Generate()

	hashedPassword, err := usecase.cryptorGateway.HashPassword(input.Password)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: hash")
		return nil, err
	}
	createUser.Password = hashedPassword

	token, err := usecase.tokenGateway.Generate(createUser.ID)
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
