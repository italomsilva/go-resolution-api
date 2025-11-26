package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApproveSolutionUsecase struct {
	solutionRepository repository.SolutionRepository
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
}

func NewApproveSolutionUsecase(
	solutionRepository repository.SolutionRepository,
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,

) ApproveSolutionUsecase {
	return ApproveSolutionUsecase{
		solutionRepository: solutionRepository,
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
	}
}

func (usecase *ApproveSolutionUsecase) Execute(ctx *gin.Context, input *dto.ApproveSolutionRequest) (bool, error) {
	problem, err := usecase.problemRepository.GetById(input.ProblemID)
	if err != nil || problem == nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return false, err
	}

	solution, err := usecase.solutionRepository.GetById(input.SolutionID)
	if err != nil || solution == nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return false, err
	}

	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	if problem.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "You are not authorized to approve this solution")
		return false, err
	}

	problem.Status = entity.ProblemStatusInProgress
	_, err = usecase.problemRepository.Update(problem.ID, problem)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Update Problem error")
		return false, err
	}

	solution.Approved = true
	_, err = usecase.solutionRepository.Update(solution.ID, solution)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Update Solution error")
		return false, err
	}

	return true, nil
}
