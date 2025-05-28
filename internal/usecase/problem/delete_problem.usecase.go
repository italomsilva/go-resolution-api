package usecase

import (
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) DeleteProblem(ctx *gin.Context, id string) (*dto.DeleteProblemResponse, error) {
	
	problem, err := usecase.problemRepository.GetProblemById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return nil, err
	}
	
	userIdToken, exists := usecase.tokenGateway.GetUserId(ctx)
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
	output := dto.DeleteProblemResponse{
		Success: isDeleted,
	}
	return &output, nil
}
