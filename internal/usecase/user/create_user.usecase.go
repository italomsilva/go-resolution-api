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

type CreateUserUsecase struct {
	userRepository     repository.UserRepository
	tokenGateway       gateway.TokenGateway
	idGeneratorGateway gateway.IDGeneratorGateway
	cryptorGateway     gateway.CryptorGateway
}

func NewCreateUserUsecase(
	userRepository repository.UserRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,
	cryptorGateway gateway.CryptorGateway,
) CreateUserUsecase {
	return CreateUserUsecase{
		userRepository:     userRepository,
		tokenGateway:       tokenGateway,
		idGeneratorGateway: idGeneratorGateway,
		cryptorGateway:     cryptorGateway,
	}
}

func (usecase *CreateUserUsecase) Execute(ctx *gin.Context, input *dto.CreateUserRequest) (*entity.User, error) {
	createUser := entity.NewUser()

	foundUserLogin, _ := usecase.userRepository.GetByLogin(input.Login)
	if foundUserLogin != nil {
		response.SendError(ctx, http.StatusConflict, "Login already in use")
		return nil, nil
	}
	createUser.Login = input.Login

	foundUserDocument, _ := usecase.userRepository.GetByDocument(input.Document)
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

	newUser, err := usecase.userRepository.Create(&createUser)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "SignUp Error")
		return nil, err
	}
	return newUser, nil
}
