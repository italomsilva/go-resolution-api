package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	dto "go-resolution-api/internal/dto/problem_sector"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProblemSectorUsecase struct {
	problemSectorRepository repository.ProblemSectorRepository
	problemRepository       repository.ProblemRepository
	sectorRepository        repository.SectorRepository
	tokenGateway            gateway.TokenGateway
}

func NewCreateProblemSectorUsecase(
	problemSectorRepository repository.ProblemSectorRepository,
	problemRepository repository.ProblemRepository,
	sectorRepository repository.SectorRepository,
	tokenGateway gateway.TokenGateway,
) CreateProblemSectorUsecase {
	return CreateProblemSectorUsecase{
		problemSectorRepository: problemSectorRepository,
		problemRepository:       problemRepository,
		sectorRepository:        sectorRepository,
		tokenGateway:            tokenGateway,
	}
}

func (usecase *CreateProblemSectorUsecase) Execute(ctx *gin.Context, input *dto.CreateProblemSectorRequest) (*entity.ProblemSector, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	problem, err := usecase.problemRepository.GetById(input.ProblemID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return nil, err
	}

	if problem.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, err
	}

	_, err = usecase.sectorRepository.GetById(input.SectorID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Sector not found")
		return nil, err
	}

	newProblemSector := entity.ProblemSector{
		ID:        0,
		ProblemID: input.ProblemID,
		SectorID:  input.SectorID,
	}

	result, err := usecase.problemSectorRepository.Create(&newProblemSector)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Create Problem Sector Failed")
		return nil, err
	}

	return result, nil
}
