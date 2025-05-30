package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteSolutionUsecase struct {
	solutionRepository repository.SolutionRepository
	tokenGateway       gateway.TokenGateway
}

func NewDeleteSolutionUsecase(
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,

) DeleteSolutionUsecase {
	return DeleteSolutionUsecase{
		solutionRepository: solutionRepository,
		tokenGateway:       tokenGateway,
	}
}

func (usecase *DeleteSolutionUsecase) Execute(ctx *gin.Context, id string) (*dto.DeleteSolutionResponse, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	solution, err := usecase.solutionRepository.GetById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return nil, err
	}

	if solution.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, err
	}

	result, err := usecase.solutionRepository.Delete(id)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Delete Solution error")
		return nil, err

	}

	output := dto.DeleteSolutionResponse{
		Success: result,
	}
	return &output, nil
}
