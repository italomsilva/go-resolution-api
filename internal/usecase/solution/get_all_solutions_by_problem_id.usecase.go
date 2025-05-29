package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllSolutionsByProblemIdUsecase struct {
	solutionRepository repository.SolutionRepository
	problemRepository  repository.ProblemRepository
}

func NewGetAllSolutionsByProblemIdUsecase(
	solutionRepository repository.SolutionRepository,
	problemRepository repository.ProblemRepository,

) GetAllSolutionsByProblemIdUsecase {
	return GetAllSolutionsByProblemIdUsecase{
		solutionRepository: solutionRepository,
		problemRepository: problemRepository,
	}
}

func (usecase *GetAllSolutionsByProblemIdUsecase) Execute(ctx *gin.Context, problemId string) ([]entity.Solution, error) {
	problem, err := usecase.problemRepository.GetById(problemId)
	if err != nil || problem == nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return nil, err
	}
	result, _ := usecase.solutionRepository.GetAllByProblemId(problemId)
	return result, nil
}
