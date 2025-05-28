package usecase

import (
	"fmt"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) GetAllProblemsByUserId(ctx *gin.Context, userId string) ([]entity.Problem, error) {
	problems := []entity.Problem{}

	userIdToken, exists := usecase.tokenGateway.GetUserId(ctx)
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
