package usecase

import (
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMySolutionsToAppUsecase struct {
	solutionRepository         repository.SolutionRepository
	problemRepository          repository.ProblemRepository
	solutionReactionRepository repository.SolutionReactionRepository
}

func NewGetMySolutionsToAppUsecase(
	solutionRepository repository.SolutionRepository,
	problemRepository repository.ProblemRepository,
	solutionReactionRepository repository.SolutionReactionRepository,
) GetMySolutionsToAppUsecase {
	return GetMySolutionsToAppUsecase{
		solutionRepository:         solutionRepository,
		problemRepository:          problemRepository,
		solutionReactionRepository: solutionReactionRepository,
	}
}

func (usecase *GetMySolutionsToAppUsecase) Execute(ctx *gin.Context, userId string) ([]dto.GetMySolutionsToAppResponse, error) {
	result := []dto.GetMySolutionsToAppResponse{}
	solutions, _ := usecase.solutionRepository.GetAllByUserId(userId)
	for _, solution := range solutions {
		problem, err := usecase.problemRepository.GetById(solution.ProblemID)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "Problem Not found")
			return result, nil
		}
		likes, dislikes, _, err := usecase.solutionReactionRepository.GetReactionsBySolutionIdAndUserId(solution.ID, userId)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "Internal Server Error")
			return result, nil
		}
		solutionDto := dto.GetMySolutionsToAppResponse{
			ID:           solution.ID,
			Title:        solution.Title,
			Description:  solution.Description,
			Approved:     solution.Approved,
			CreatedAt:    solution.CreatedAt,
			ProblemID:    solution.ProblemID,
			ProblemTitle: problem.Title,
			Likes:        likes,
			Dislikes:     dislikes,
		}
		result = append(result, solutionDto)
	}
	return result, nil
}	
