package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetSectorByIDUsecase struct {
	sectorRepository repository.SectorRepository
}

func NewGetSectorByIDUsecase(
	sectorRepository repository.SectorRepository,
) GetSectorByIDUsecase {
	return GetSectorByIDUsecase{
		sectorRepository: sectorRepository,
	}
}

func (usecase *GetSectorByIDUsecase) Execute(ctx *gin.Context, id int) (*entity.Sector, error) {
	result, err := usecase.sectorRepository.GetById(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Sector Not Found")
		return nil, err
	}
	return result, nil
}
