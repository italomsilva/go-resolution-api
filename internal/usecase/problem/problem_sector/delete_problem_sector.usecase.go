package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	dto "go-resolution-api/internal/dto/problem_sector"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProblemSectorUsecase struct {
	problemSectorRepository repository.ProblemSectorRepository
	problemRepository       repository.ProblemRepository
	tokenGateway            gateway.TokenGateway
}

func NewDeleteProblemSectorUsecase(
	problemSectorRepository repository.ProblemSectorRepository,
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
) DeleteProblemSectorUsecase {
	return DeleteProblemSectorUsecase{
		problemSectorRepository: problemSectorRepository,
		problemRepository:       problemRepository,
		tokenGateway:            tokenGateway,
	}
}

func (usecase *DeleteProblemSectorUsecase) Execute(ctx *gin.Context, input *dto.DeleteProblemSectorRequest) (*dto.DeleteProblemSectorResponse, error) {
	output := dto.DeleteProblemSectorResponse{Success: false}

	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	problemSector, err := usecase.problemSectorRepository.GetById(input.ID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "ProblemSector not found")
		return &output, err
	}

	problem, err := usecase.problemRepository.GetById(problemSector.ProblemID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return &output, err
	}

	if problem.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return &output, err
	}

	result, err := usecase.problemSectorRepository.Delete(input.ID)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Delete Failed")
		return &output, err
	}
	output.Success = result
	return &output, nil
}
