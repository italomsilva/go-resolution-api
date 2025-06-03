package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllSectorsByProblemIdUsecase struct {
	problemSectorRepository repository.ProblemSectorRepository
	problemRepository       repository.ProblemRepository
	sectorRepository        repository.SectorRepository
	tokenGateway            gateway.TokenGateway
}

func NewGetAllSectorsByProblemIdUsecase(
	problemSectorRepository repository.ProblemSectorRepository,
	problemRepository repository.ProblemRepository,
	sectorRepository repository.SectorRepository,
	tokenGateway gateway.TokenGateway,
) GetAllSectorsByProblemIdUsecase {
	return GetAllSectorsByProblemIdUsecase{
		problemSectorRepository: problemSectorRepository,
		problemRepository:       problemRepository,
		sectorRepository:        sectorRepository,
		tokenGateway:            tokenGateway,
	}
}

func (usecase *GetAllSectorsByProblemIdUsecase) Execute(ctx *gin.Context, problemId string) ([]entity.Sector, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	problem, err := usecase.problemRepository.GetById(problemId)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem not found")
		return []entity.Sector{}, err
	}

	if problem.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return []entity.Sector{}, err
	}

	problemSectors, err := usecase.problemSectorRepository.GetAllByProblemId(problemId)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem Sectors not found")
		return []entity.Sector{}, err
	}

	var sectorIds []int
	for _, problemSector := range problemSectors {
		sectorIds = append(sectorIds, problemSector.SectorID)
	}

	result, err := usecase.sectorRepository.GetByIds(sectorIds)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Get Sectors Failed")
		return []entity.Sector{}, err
	}
	return result, nil
}
