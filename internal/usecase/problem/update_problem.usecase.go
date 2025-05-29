package usecase

import (
	"fmt"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateProblemUsecase struct {
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
}

func NewUpdateProblemUsecase(
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
) UpdateProblemUsecase {
	return UpdateProblemUsecase{
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
	}
}


func (usecase *UpdateProblemUsecase) Execute(ctx *gin.Context, input *dto.UpdateProblemRequest) (*entity.Problem, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)
	
	problem, err := usecase.problemRepository.GetById(input.ID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem Not Found")
		return nil, err
	}

	if userId !=  problem.UserID{
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized User")
		return nil, fmt.Errorf("unauthorized user")
	}

	if input.Title != nil && *input.Title != "" {
		problem.Title = *input.Title
	}

	if input.Description != nil {
		problem.Description = *input.Description
	}

	if input.Location != nil && *input.Location != "" {
		problem.Location = *input.Location
	}

	if input.Status != nil {
		problem.Status = *input.Status
	}

	result, err := usecase.problemRepository.Update(problem.ID, problem)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "update problem error")
		return nil, err
	}

	return result, nil

}
