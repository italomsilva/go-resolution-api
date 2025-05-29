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

type CreateProblemUsecase struct {
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
	idGeneratorGateway gateway.IDGeneratorGateway
}

func NewCreateProblemUsecase(
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,
) CreateProblemUsecase {
	return CreateProblemUsecase{
		problemRepository:  problemRepository,
		idGeneratorGateway: idGeneratorGateway,
		tokenGateway:       tokenGateway,
	}
}

func (usecase *CreateProblemUsecase) Execute(ctx *gin.Context, input *dto.CreateProblemRequest) (*entity.Problem, error) {
	problem := entity.NewProblem()
	id := usecase.idGeneratorGateway.Generate()
	userId, exists := usecase.tokenGateway.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return nil, fmt.Errorf("authentication required")
	}
	problem.ID = id
	problem.UserID = userId
	problem.Title = input.Title
	problem.Description = input.Description
	problem.Location = input.Location
	if input.Status != nil {
		problem.Status = *input.Status
	}

	result, err := usecase.problemRepository.CreateProblem(&problem)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Create Problem Error")
		return nil, err
	}

	return result, nil

}
