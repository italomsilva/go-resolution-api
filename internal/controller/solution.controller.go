package controller

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution"
	usecase "go-resolution-api/internal/usecase/solution"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SolutionController struct {
	tokenGateway                      gateway.TokenGateway
	createSolutionUsecase             usecase.CreateSolutionUsecase
	getAllSolutionsByProblemIdUsecase usecase.GetAllSolutionsByProblemIdUsecase
}

func NewSolutionController(
	tokenGateway gateway.TokenGateway,
	createSolutionUsecase usecase.CreateSolutionUsecase,
	getAllSolutionsByProblemIdUsecase usecase.GetAllSolutionsByProblemIdUsecase,

) SolutionController {
	return SolutionController{
		tokenGateway:                      tokenGateway,
		createSolutionUsecase:             createSolutionUsecase,
		getAllSolutionsByProblemIdUsecase: getAllSolutionsByProblemIdUsecase,
	}
}

func (controller *SolutionController) CreateSolution(ctx *gin.Context) {
	body := dto.CreateSolutionRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.createSolutionUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "")
		return
	}

}

func (controller *SolutionController) GetAllSolutionsByProblemId(ctx *gin.Context) {
	problemId := ctx.Param("problemId")
	result, _ := controller.getAllSolutionsByProblemIdUsecase.Execute(ctx, problemId)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "")
		return
	}

}
