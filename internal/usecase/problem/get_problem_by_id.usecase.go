package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) GetProblemByID(ctx *gin.Context, id string) (*entity.Problem, error){
	result, err := usecase.problemRepository.GetProblemById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound,"Problem Not Found")
		return nil, err
	}
	return result, nil
}