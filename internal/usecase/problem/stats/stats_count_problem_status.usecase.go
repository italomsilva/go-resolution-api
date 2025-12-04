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

type StatsCountProblemStatusUsecase struct {
	problemRepository repository.ProblemRepository
	tokenGateway      gateway.TokenGateway
}

func NewStatsCountProblemStatusUsecase(
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
) StatsCountProblemStatusUsecase {
	return StatsCountProblemStatusUsecase{
		problemRepository: problemRepository,
		tokenGateway:      tokenGateway,
	}
}

func (usecase *StatsCountProblemStatusUsecase) Execute(ctx *gin.Context, userId string) ([]dto.StatsCountProblemStatusResponse, error) {
	stats := []dto.StatsCountProblemStatusResponse{}

	userIdToken, _ := usecase.tokenGateway.GetUserId(ctx)
	if userIdToken != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized User")
		return stats, fmt.Errorf("unauthorized user")
	}

	stats, err := usecase.problemRepository.GetStatsCountProblemStatus(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Fetch Users Failed")
		return stats, err
	}
	return stats, nil

}
