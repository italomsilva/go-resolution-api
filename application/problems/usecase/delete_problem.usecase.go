package usecase

import (
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) DeleteProblem(ctx *gin.Context, id string) (*dto.ResDeleteProblem, error) {
	output := dto.NewResDeleteProblem()

	problem, err := usecase.problemRepository.GetProblemById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return nil, err
	}

	userIdToken, exists := utils.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return nil, err
	}

	if problem.UserID != userIdToken {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, err
	}

	isDeleted, err := usecase.problemRepository.DeleteProblem(id)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Error deleting")
		return nil, err
	}
	output.Success = isDeleted

	return &output, nil
}
