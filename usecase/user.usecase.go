package usecase

import (
	"fmt"
	"go-resolution-api/dto/response"
	"go-resolution-api/dto/user"
	"go-resolution-api/gateway"
	"go-resolution-api/model"
	"go-resolution-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
	}
}

func (usecase *UserUseCase) GetUsers() ([]model.User, error) {
	return usecase.userRepository.GetUsers()
}

func (usecase *UserUseCase) GetUserById(ctx *gin.Context, id string) (*model.User, error) {
	user, err := usecase.userRepository.GetUserById(id)
	if err != nil || user == nil {
		response.SendError(ctx, http.StatusNotFound, "User not found")
		return nil, err
	}
	return user, nil
}

func (usecase *UserUseCase) CreateUser(ctx *gin.Context, input *userDto.ReqCreateUser) (*model.User, error) {
	foundUserLogin, _ := usecase.userRepository.GetUserByLogin(input.Login)
	if foundUserLogin != nil {
		response.SendError(ctx, http.StatusConflict, "Login already in use")
		return nil, nil
	}

	foundUserDocument, _ := usecase.userRepository.GetUserByDocument(input.Document)
	if foundUserDocument != nil {
		response.SendError(ctx, http.StatusConflict, "Document already in use")
		return nil, nil
	}

	id := gateway.GenerateUUID()
	input.ID = id

	hashedPassword, err := gateway.HashPassword(input.Password)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: hash")
		return nil, err
	}
	input.Password = hashedPassword

	token, err := gateway.GenerateJWT(input.ID)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: jwt")
		return nil, err
	}
	input.Token = token

	newUser, err := usecase.userRepository.CreateUser(input)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "SignUp Error")
		return nil, err
	}
	return newUser, nil
}

func (usecase *UserUseCase) Login(ctx *gin.Context, input *userDto.ReqLogin) (*model.User, error) {
	user, _ := usecase.userRepository.GetUserByLogin(input.Login)
	if user == nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return nil, nil
	}

	comparePasswords := gateway.CheckPasswordHash(input.Password, user.Password)
	if !comparePasswords {
		response.SendError(ctx, http.StatusBadRequest, "Invalid login or password")
		return nil, nil
	}

	newToken, err := gateway.GenerateJWT(user.ID)
	if err != nil {
		response.SendError(ctx, http.StatusBadGateway, "Gateway error: jwt")
		return nil, nil
	}

	user.Token = newToken
	updatedUser := userDto.NewReqUpdateUser(user)
	return usecase.userRepository.UpdateUser(user.ID, &updatedUser)

}

func (usecase *UserUseCase) UpdateUser(ctx *gin.Context, input *userDto.ReqUpdateUser) (*model.User, error) {
	userId, exists := ctx.Get("userId")
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required.")
		return nil, nil
	}
	userIdStr := fmt.Sprintf("%v", userId)
	foundUserById, err := usecase.userRepository.GetUserById(userIdStr)
	if foundUserById == nil {
		response.SendError(ctx, http.StatusNotFound, "User Not Found")
		return nil, err
	}

	userToUpdate := userDto.NewReqUpdateUser(foundUserById)

	if input.Name != nil {
		userToUpdate.Name = input.Name
	}

	if input.Login != nil {
		foundUserByLogin, _ := usecase.userRepository.GetUserByLogin(*input.Login)
		if foundUserByLogin != nil {
			response.SendError(ctx, http.StatusConflict, "Login already exists")
			return nil, nil
		}
		userToUpdate.Login = input.Login
	}
	return usecase.userRepository.UpdateUser(userIdStr, &userToUpdate)
}

func (usecase *UserUseCase) DeleteUser(ctx *gin.Context, input *userDto.ReqDeleteUser) (*userDto.ResDeleteUser, error) {
	responseDelete := userDto.NewResDeleteUser()
	responseDelete.Success = false

	inputLogin := userDto.NewReqLogin()
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
