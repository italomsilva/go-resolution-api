package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatsCountSolutionsReactionsUsecase struct {
	solutionReactionRepository repository.SolutionReactionRepository
	solutionRepository         repository.SolutionRepository
	tokenGateway               gateway.TokenGateway
}

func NewStatsCountSolutionsReactionsUsecase(
	solutionReactionRepository repository.SolutionReactionRepository,
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,

) StatsCountSolutionsReactionsUsecase {
	return StatsCountSolutionsReactionsUsecase{
		solutionReactionRepository: solutionReactionRepository,
		solutionRepository:         solutionRepository,
		tokenGateway:               tokenGateway,
	}
}

func (usecase *StatsCountSolutionsReactionsUsecase) Execute(ctx *gin.Context) (*[]dto.StatsCountSolutionsReactionsResponseDto, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	solutions, err := usecase.solutionRepository.GetAllByUserId(userId)
	if err != nil || len(solutions) == 0 {
		response.SendError(ctx, http.StatusNotFound, "Solutions not found")
		return nil, err
	}

	result := []dto.StatsCountSolutionsReactionsResponseDto{}

	for _, solution := range solutions {
		likes, dislikes, _, err := usecase.solutionReactionRepository.GetReactionsBySolutionIdAndUserId(solution.ID, userId)
		if err == nil {
			solutionDto := dto.StatsCountSolutionsReactionsResponseDto{
				SolutionTitle: solution.Title,
				Likes: likes,
				Dislikes: dislikes,
				CreatedAt: solution.CreatedAt,
			}
			result = append(result, solutionDto)
		}
	}

	return &result, nil
}
