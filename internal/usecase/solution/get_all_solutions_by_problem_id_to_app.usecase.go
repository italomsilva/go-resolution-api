package usecase

import (
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllSolutionsByProblemIdToAppUsecase struct {
	solutionRepository         repository.SolutionRepository
	problemRepository          repository.ProblemRepository
	solutionReactionRepository repository.SolutionReactionRepository
	userRepository             repository.UserRepository
}

func NewGetAllSolutionsByProblemIdToAppUsecase(
	solutionRepository repository.SolutionRepository,
	problemRepository repository.ProblemRepository,
	solutionReactionRepository repository.SolutionReactionRepository,
	userRepository repository.UserRepository,

) GetAllSolutionsByProblemIdToAppUsecase {
	return GetAllSolutionsByProblemIdToAppUsecase{
		solutionRepository:         solutionRepository,
		problemRepository:          problemRepository,
		solutionReactionRepository: solutionReactionRepository,
		userRepository:             userRepository,
	}
}

func (usecase *GetAllSolutionsByProblemIdToAppUsecase) Execute(ctx *gin.Context, problemId string) ([]dto.GetAllSolutionsByProblemIdToAAppResponse, error) {
	problem, err := usecase.problemRepository.GetById(problemId)
	if err != nil || problem == nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return nil, err
	}
	result := []dto.GetAllSolutionsByProblemIdToAAppResponse{}
	solutions, _ := usecase.solutionRepository.GetAllByProblemId(problemId)
	for _, solution := range solutions {
		user, err := usecase.userRepository.GetById(solution.UserID)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "User Not found")
			return result, nil
		}
		likes, dislikes, myReaction, err := usecase.solutionReactionRepository.GetReactionsBySolutionIdAndUserId(solution.ID, solution.UserID)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "Internal Server Error")
			return result, nil
		}
		solutionDto := dto.GetAllSolutionsByProblemIdToAAppResponse{
			ID:            solution.ID,
			Title:         solution.Title,
			Description:   solution.Description,
			EstimatedCost: float64(solution.EstimatedCost),
			Approved:      solution.Approved,
			CreatedAt:     solution.CreatedAt,
			ProblemID:     solution.ProblemID,
			ProblemTitle:  problem.Title,
			UserID:        solution.UserID,
			UserLogin:     user.Login,
			Likes:         likes,
			Dislikes:      dislikes,
			MyReaction:    myReaction,
		}
		result = append(result, solutionDto)

	}
	return result, nil
}
