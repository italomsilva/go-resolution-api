package controller

import (
	"go-resolution-api/internal/domain/gateway"
	dto "go-resolution-api/internal/dto/problem_sector"
	"go-resolution-api/internal/dto/response"
	usecase "go-resolution-api/internal/usecase/problem/problem_sector"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemSectorController struct {
	tokenGateway                    gateway.TokenGateway
	createProblemSectorUsecase      usecase.CreateProblemSectorUsecase
	deleteProblemSectorUsecase      usecase.DeleteProblemSectorUsecase
	getAllSectorsByProblemIdUsecase usecase.GetAllSectorsByProblemIdUsecase
}

func NewProblemSectorController(
	tokenGateway gateway.TokenGateway,
	createProblemSectorUsecase usecase.CreateProblemSectorUsecase,
	deleteProblemSectorUsecase usecase.DeleteProblemSectorUsecase,
	getAllSectorsByProblemIdUsecase usecase.GetAllSectorsByProblemIdUsecase,
) ProblemSectorController {
	return ProblemSectorController{
		tokenGateway:                    tokenGateway,
		createProblemSectorUsecase:      createProblemSectorUsecase,
		deleteProblemSectorUsecase:      deleteProblemSectorUsecase,
		getAllSectorsByProblemIdUsecase: getAllSectorsByProblemIdUsecase,
	}
}

func (controller *ProblemSectorController) CreateProblemSector(ctx *gin.Context) {
	body := dto.CreateProblemSectorRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	result, _ := controller.createProblemSectorUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "")
	}
}

func (controller *ProblemSectorController) DeleteProblemSector(ctx *gin.Context) {
	body := dto.DeleteProblemSectorRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	result, _ := controller.deleteProblemSectorUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemSectorController) GetSectorsByProblemID(ctx *gin.Context) {
	problemID := ctx.Param("problemId")
	result, _ := controller.getAllSectorsByProblemIdUsecase.Execute(ctx, problemID)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}
