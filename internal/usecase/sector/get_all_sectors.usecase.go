package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllSectorsUsecase struct {
	sectorRepository repository.SectorRepository
}

func NewGetAllSectorsUsecase(
	sectorRepository repository.SectorRepository,
) GetAllSectorsUsecase {
	return GetAllSectorsUsecase{
		sectorRepository: sectorRepository,
	}
}

func (usecase *GetAllSectorsUsecase) Execute(ctx *gin.Context) ([]entity.Sector, error) {
	result, err := usecase.sectorRepository.GetAll()
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Get Sectors Failed")
		return nil, err
	}
	return result, nil
}
