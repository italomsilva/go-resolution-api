package usecase

import (
	"fmt"
	"go-resolution-api/application/problems/model"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) GetAllProblemsByUserId(ctx *gin.Context, userId string) ([]model.Problem, error) {
	problems := []model.Problem{}

	userIdToken, exists := utils.GetUserId(ctx)
	if !exists || userIdToken != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return problems, fmt.Errorf("authentication required")
	}

	problems, err := usecase.problemRepository.GetAllProblemsByUserId(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Fetch Users Failed")
		return problems, err
	}
	return problems, nil

}
