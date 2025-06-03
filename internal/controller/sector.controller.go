package controller

import (
	"go-resolution-api/internal/dto/response"
	usecase "go-resolution-api/internal/usecase/sector"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SectorController struct {
	getAllSectorsUsecase usecase.GetAllSectorsUsecase
	getSectorByIdUsecase usecase.GetSectorByIDUsecase
}

func NewSectorController(
	getAllSectorsUsecase usecase.GetAllSectorsUsecase,
	getSectorByIdUsecase usecase.GetSectorByIDUsecase,
) SectorController {
	return SectorController{
		getAllSectorsUsecase: getAllSectorsUsecase,
		getSectorByIdUsecase: getSectorByIdUsecase,
	}
}

func (controller *SectorController) GetAllSectors(ctx *gin.Context) {
	result, _ := controller.getAllSectorsUsecase.Execute(ctx)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *SectorController) GetSectorById(ctx *gin.Context) {
	sectorIdString := ctx.Param("sectorId")
	sectorId, err := strconv.Atoi(sectorIdString)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	
	result, _ := controller.getSectorByIdUsecase.Execute(ctx, sectorId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}
