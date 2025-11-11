package usecase

import (
	"go-resolution-api/internal/domain/repository"
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProblemsToAppUsecase struct {
	problemRepository  repository.ProblemRepository
	userRepository     repository.UserRepository
	solutionRepository repository.SolutionRepository
}

func NewGetAllProblemsToAppUsecase(
	problemRepository repository.ProblemRepository,
	userRepository repository.UserRepository,
	solutionRepository repository.SolutionRepository,

) GetAllProblemsToAppUsecase {
	return GetAllProblemsToAppUsecase{
		problemRepository:  problemRepository,
		userRepository:     userRepository,
		solutionRepository: solutionRepository,
	}
}

func (usecase *GetAllProblemsToAppUsecase) Execute(ctx *gin.Context) ([]dto.GetProblemToAppResponse, error) {
	problems, err := usecase.problemRepository.GetAll()
	result := []dto.GetProblemToAppResponse{}
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Internal Server Error")
		return nil, err
	}

	for _, problem := range problems {
		user, err := usecase.userRepository.GetById(problem.UserID)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "User Not found")
			return nil, err
		}

		solutions, err := usecase.solutionRepository.GetAllByProblemId(problem.ID)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, err.Error())
			return nil, err
		}
		solutionsCount := 0
		solutionsCount = len(solutions)

		userLogin := ""
		userLogin = user.Login

		problemDto := dto.GetProblemToAppResponse{
			ID:             problem.ID,
			Title:          problem.Title,
			Description:    problem.Description,
			Location:       problem.Location,
			Status:         problem.Status,
			CreatedAt:      problem.CreatedAt,
			UserID:         problem.UserID,
			UserLogin:      userLogin,
			SolutionsCount: solutionsCount,
		}
		result = append(result, problemDto)
	}

	return result, nil
}
