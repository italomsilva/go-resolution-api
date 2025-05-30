package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateSolutionUsecase struct {
	solutionRepository repository.SolutionRepository
	tokenGateway       gateway.TokenGateway
}

func NewUpdateSolutionUsecase(
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,

) UpdateSolutionUsecase {
	return UpdateSolutionUsecase{
		solutionRepository: solutionRepository,
		tokenGateway:       tokenGateway,
	}
}

func (usecase *UpdateSolutionUsecase) Execute(ctx *gin.Context, input *dto.UpdateSolutionRequest) (*entity.Solution, error) {
	solution, err := usecase.solutionRepository.GetById(*input.ID)
	if err != nil || solution == nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return nil, err
	}

	userId, _ := usecase.tokenGateway.GetUserId(ctx)
	if userId != solution.UserID {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, nil
	}

	if input.Title != nil || *input.Title != "" {
		solution.Title = *input.Title
	}

	if input.Description != nil {
		solution.Description = *input.Description
	}

	if input.EstimatedCost != nil {
		solution.EstimatedCost = *input.EstimatedCost
	}

	result, err := usecase.solutionRepository.Create(solution)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Update Solution error")
		return nil, err

	}
	return result, nil
}
