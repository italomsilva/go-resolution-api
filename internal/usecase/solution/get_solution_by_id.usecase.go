package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetSolutionByIdUsecase struct {
	solutionRepository repository.SolutionRepository
}

func NewGetSolutionByIdUsecase(
	solutionRepository repository.SolutionRepository,

) GetSolutionByIdUsecase {
	return GetSolutionByIdUsecase{
		solutionRepository: solutionRepository,
	}
}

func (usecase *GetSolutionByIdUsecase) Execute(ctx *gin.Context, id string) (*entity.Solution, error) {
	solution, err := usecase.solutionRepository.GetById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return nil, err
	}
	return solution, nil
}
