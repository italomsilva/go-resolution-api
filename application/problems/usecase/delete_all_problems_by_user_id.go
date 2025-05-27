package usecase

import (
	"fmt"
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) DeleteAllProblemsByUserId(ctx *gin.Context, userId string) (*dto.ResDeleteAllProblemsByUserId, error) {

	userIdToken, exists := utils.GetUserId(ctx)
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

	output := dto.NewResDeleteAllProblemsByUserId()
	output.DeletedCounter = deletedCount
	output.Success = true

	return &output, nil
}
