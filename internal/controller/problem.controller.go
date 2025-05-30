package controller

import (
	"go-resolution-api/internal/domain/gateway"
	dto "go-resolution-api/internal/dto/problem"
	"go-resolution-api/internal/dto/response"
	usecase "go-resolution-api/internal/usecase/problem"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemController struct {
	tokenGateway                     gateway.TokenGateway
	createProblemUsecase             usecase.CreateProblemUsecase
	deleteAllProblemsByUserIdUsecase usecase.DeleteAllProblemsByUserIdUsecase
	deleteProblemUsecase             usecase.DeleteProblemUsecase
	getAllProblemsByUserIdUsecase    usecase.GetAllProblemsByUserIdUsecase
	getAllProblemsUsecase            usecase.GetAllProblemsUsecase
	getProblemByIDUsecase            usecase.GetProblemByIDUsecase
	updateProblemUsecase             usecase.UpdateProblemUsecase
}

func NewProblemController(
	tokenGateway gateway.TokenGateway,
	createProblemUsecase usecase.CreateProblemUsecase,
	deleteAllProblemsByUserIdUsecase usecase.DeleteAllProblemsByUserIdUsecase,
	deleteProblemUsecase usecase.DeleteProblemUsecase,
	getAllProblemsByUserIdUsecase usecase.GetAllProblemsByUserIdUsecase,
	getAllProblemsUsecase usecase.GetAllProblemsUsecase,
	getProblemByIDUsecase usecase.GetProblemByIDUsecase,
	updateProblemUsecase usecase.UpdateProblemUsecase,
) ProblemController {
	return ProblemController{
		tokenGateway:                     tokenGateway,
		createProblemUsecase:             createProblemUsecase,
		deleteAllProblemsByUserIdUsecase: deleteAllProblemsByUserIdUsecase,
		deleteProblemUsecase:             deleteProblemUsecase,
		getAllProblemsByUserIdUsecase:    getAllProblemsByUserIdUsecase,
		getAllProblemsUsecase:            getAllProblemsUsecase,
		getProblemByIDUsecase:            getProblemByIDUsecase,
		updateProblemUsecase:             updateProblemUsecase,
	}
}

func (controller *ProblemController) GetAllProblems(ctx *gin.Context) {
	result, _ := controller.getAllProblemsUsecase.Execute()
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) GetProblemById(ctx *gin.Context) {
	problemId := ctx.Param("problemId")
	result, _ := controller.getProblemByIDUsecase.Execute(ctx, problemId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) CreateProblem(ctx *gin.Context) {
	body := dto.CreateProblemRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.createProblemUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "")
	}
}

func (controller *ProblemController) UpdateProblem(ctx *gin.Context) {
	body := dto.UpdateProblemRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.updateProblemUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) GetAllProblemsByUserId(ctx *gin.Context) {
	userId, _ := controller.tokenGateway.GetUserId(ctx)
	result, err := controller.getAllProblemsByUserIdUsecase.Execute(ctx, userId)
	if err == nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) DeleteProblem(ctx *gin.Context) {
	body := dto.DeleteProblemRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	result, _ := controller.deleteProblemUsecase.Execute(ctx, body.ID)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) DeleteAllProblemsByUserId(ctx *gin.Context) {
	userId, _ := controller.tokenGateway.GetUserId(ctx)
	result, _ := controller.deleteAllProblemsByUserIdUsecase.Execute(ctx, userId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}
