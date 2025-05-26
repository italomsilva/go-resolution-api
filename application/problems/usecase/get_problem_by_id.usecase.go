package usecase

import (
	"go-resolution-api/application/problems/model"
	"go-resolution-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) GetProblemByID(ctx *gin.Context, id string) (*model.Problem, error){
	result, err := usecase.problemRepository.GetProblemById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound,"Problem Not Found")
		return nil, err
	}
	return result, nil
}