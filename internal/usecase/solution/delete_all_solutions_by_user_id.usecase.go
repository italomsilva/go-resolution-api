package usecase

import (
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteAllSolutionsByUserIdUsecase struct {
	solutionRepository repository.SolutionRepository
	userRepository     repository.UserRepository
}

func NewDeleteAllSolutionsByUserIdUsecase(
	solutionRepository repository.SolutionRepository,
	userRepository repository.UserRepository,

) DeleteAllSolutionsByUserIdUsecase {
	return DeleteAllSolutionsByUserIdUsecase{
		solutionRepository: solutionRepository,
		userRepository:     userRepository,
	}
}

func (usecase *DeleteAllSolutionsByUserIdUsecase) Execute(ctx *gin.Context, userId string) (*dto.DeleteAllSolutionsByUserIdResponse, error) {
	_, err := usecase.userRepository.GetById(userId)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "User not found")
		return nil, err
	}

	result, err := usecase.solutionRepository.DeleteAllByUserId(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Delete Solutions error")
		return nil, err

	}

	output := dto.DeleteAllSolutionsByUserIdResponse{
		Success:        true,
		DeletedCounter: result,
	}
	return &output, nil
}
