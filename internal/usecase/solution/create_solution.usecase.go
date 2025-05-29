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

type CreateSolutionUsecase struct {
	solutionRepository repository.SolutionRepository
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
	idGeneratorGateway gateway.IDGeneratorGateway
}

func NewCreateSolutionUsecase(
	solutionRepository repository.SolutionRepository,
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,

) CreateSolutionUsecase {
	return CreateSolutionUsecase{
		solutionRepository: solutionRepository,
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
		idGeneratorGateway: idGeneratorGateway,
	}
}

func (usecase *CreateSolutionUsecase) Execute(ctx *gin.Context, input *dto.CreateSolutionRequest) (*entity.Solution, error) {
	problem, err := usecase.problemRepository.GetById(input.ProblemId)
	if err != nil || problem == nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return nil, err
	}

	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	newSolution := entity.NewSolution()
	newSolution.ID = usecase.idGeneratorGateway.Generate()
	newSolution.Title = input.Title
	newSolution.Description = input.Description
	newSolution.EstimatedCost = input.Estimated_cost
	newSolution.ProblemId = input.ProblemId
	newSolution.UserId = userId

	result, err := usecase.solutionRepository.Create(&newSolution)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Create Solution error")
		return nil, err

	}
	return result, nil
}
