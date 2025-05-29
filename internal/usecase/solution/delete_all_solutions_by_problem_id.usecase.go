package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteAllSolutionsByProblemIdUsecase struct {
	solutionRepository repository.SolutionRepository
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
}

func NewDeleteAllSolutionsByProblemIdUsecase(
	solutionRepository repository.SolutionRepository,
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,

) DeleteAllSolutionsByProblemIdUsecase {
	return DeleteAllSolutionsByProblemIdUsecase{
		solutionRepository: solutionRepository,
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
	}
}

func (usecase *DeleteAllSolutionsByProblemIdUsecase) Execute(ctx *gin.Context, problemId string) (*dto.DeleteAllSolutionsByProblemIdResponse, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	problem, err := usecase.problemRepository.GetById(problemId)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "problem not found")
		return nil, err
	}

	if problem.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, err
	}

	result, err := usecase.solutionRepository.DeleteAllByProblemId(problemId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Delete Solutions error")
		return nil, err

	}

	output := dto.DeleteAllSolutionsByProblemIdResponse{
		Success:        true,
		DeletedCounter: result,
	}
	return &output, nil
}
