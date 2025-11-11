package usecase

import (
	"go-resolution-api/internal/domain/repository"
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProblemByIdToAppUsecase struct {
	problemRepository  repository.ProblemRepository
	userRepository     repository.UserRepository
	solutionRepository repository.SolutionRepository
}

func NewGetProblemByIdToAppUsecase(
	problemRepository repository.ProblemRepository,
	userRepository repository.UserRepository,
	solutionRepository repository.SolutionRepository,

) GetProblemByIdToAppUsecase {
	return GetProblemByIdToAppUsecase{
		problemRepository:  problemRepository,
		userRepository:     userRepository,
		solutionRepository: solutionRepository,
	}
}

func (usecase *GetProblemByIdToAppUsecase) Execute(ctx *gin.Context, id string) (*dto.GetProblemToAppResponse, error) {
	problem, err := usecase.problemRepository.GetById(id)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Internal Server Error")
		return nil, err
	}

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

	return &problemDto, nil
}
