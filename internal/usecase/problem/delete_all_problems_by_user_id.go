package usecase

import (
	"fmt"
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) DeleteAllProblemsByUserId(ctx *gin.Context, userId string) (*dto.DeleteAllProblemsByUserIdResponse, error) {

	userIdToken, exists := usecase.tokenGateway.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return nil, fmt.Errorf("authentication required")
	}

	if userIdToken != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, fmt.Errorf("unauthorized user")
	}

	deletedCount, err := usecase.problemRepository.DeleteAllProblemsByUserId(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Error deleting")
		return nil, err

	}

	output := dto.DeleteAllProblemsByUserIdResponse{
		DeletedCounter: deletedCount,
		Success: true,
	}
	return &output, nil
}
