package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProblemByIDUsecase struct {
	problemRepository  repository.ProblemRepository
}

func NewGetProblemByIDUsecase(
	problemRepository repository.ProblemRepository,
) GetProblemByIDUsecase {
	return GetProblemByIDUsecase{
		problemRepository:  problemRepository,
	}
}


func (usecase *GetProblemByIDUsecase) Execute(ctx *gin.Context, id string) (*entity.Problem, error){
	result, err := usecase.problemRepository.GetById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound,"Problem Not Found")
		return nil, err
	}
	return result, nil
}