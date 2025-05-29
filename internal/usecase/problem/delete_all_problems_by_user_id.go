package usecase

import (
	"fmt"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteAllProblemsByUserIdUsecase struct {
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
}

func NewDeleteAllProblemsByUserIdUsecase(
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
) DeleteAllProblemsByUserIdUsecase {
	return DeleteAllProblemsByUserIdUsecase{
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
	}
}


func (usecase *DeleteAllProblemsByUserIdUsecase) Execute(ctx *gin.Context, userId string) (*dto.DeleteAllProblemsByUserIdResponse, error) {

	userIdToken, exists := usecase.tokenGateway.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return nil, fmt.Errorf("authentication required")
	}

	if userIdToken != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, fmt.Errorf("unauthorized user")
	}

	deletedCount, err := usecase.problemRepository.DeleteAllByUserId(userId)
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
