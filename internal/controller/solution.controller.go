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
	tokenGateway                         gateway.TokenGateway
	createSolutionUsecase                usecase.CreateSolutionUsecase
	getAllSolutionsByProblemIdUsecase    usecase.GetAllSolutionsByProblemIdUsecase
	getSolutionByIdUsecase               usecase.GetSolutionByIdUsecase
	deleteSolutionUsecase                usecase.DeleteSolutionUsecase
	deleteAllSolutionsByProblemIdUsecase usecase.DeleteAllSolutionsByProblemIdUsecase
	deleteAllSolutionsByUserIdUsecase    usecase.DeleteAllSolutionsByUserIdUsecase
	updateSolutionUsecase                usecase.UpdateSolutionUsecase
}

func NewSolutionController(
	tokenGateway gateway.TokenGateway,
	createSolutionUsecase usecase.CreateSolutionUsecase,
	getAllSolutionsByProblemIdUsecase usecase.GetAllSolutionsByProblemIdUsecase,
	getSolutionByIdUsecase usecase.GetSolutionByIdUsecase,
	deleteSolutionUsecase usecase.DeleteSolutionUsecase,
	deleteAllSolutionsByProblemIdUsecase usecase.DeleteAllSolutionsByProblemIdUsecase,
	deleteAllSolutionsByUserIdUsecase usecase.DeleteAllSolutionsByUserIdUsecase,
	updateSolutionUsecase usecase.UpdateSolutionUsecase,

) SolutionController {
	return SolutionController{
		tokenGateway:                         tokenGateway,
		createSolutionUsecase:                createSolutionUsecase,
		getAllSolutionsByProblemIdUsecase:    getAllSolutionsByProblemIdUsecase,
		getSolutionByIdUsecase:               getSolutionByIdUsecase,
		deleteSolutionUsecase:                deleteSolutionUsecase,
		deleteAllSolutionsByProblemIdUsecase: deleteAllSolutionsByProblemIdUsecase,
		deleteAllSolutionsByUserIdUsecase:    deleteAllSolutionsByUserIdUsecase,
		updateSolutionUsecase:                updateSolutionUsecase,
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
		response.SendSucess(ctx, http.StatusOK, result, "")
		return
	}

}

func (controller *SolutionController) GetSolutionById(ctx *gin.Context) {
	solutionId := ctx.Param("solutionId")
	result, _ := controller.getSolutionByIdUsecase.Execute(ctx, solutionId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
		return
	}
}

func (controller *SolutionController) DeleteSolution(ctx *gin.Context) {
	body := dto.DeleteSolutionRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.deleteSolutionUsecase.Execute(ctx, body.ID)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
		return
	}

}

func (controller *SolutionController) DeleteAllSolutionsByProblemId(ctx *gin.Context) {
	body := dto.DeleteAllSolutionsByProblemIdRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.deleteAllSolutionsByProblemIdUsecase.Execute(ctx, body.ProblemId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
		return
	}
}

func (controller *SolutionController) DeleteAllSolutionsByUserId(ctx *gin.Context) {
	userId, _ := controller.tokenGateway.GetUserId(ctx)
	result, _ := controller.deleteAllSolutionsByUserIdUsecase.Execute(ctx, userId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
		return
	}
}

func (controller *SolutionController) UpdateSolution(ctx *gin.Context) {
	body := dto.UpdateSolutionRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.updateSolutionUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
		return
	}
}
