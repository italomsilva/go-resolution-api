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
	usecase usecase.ProblemUseCase
	tokenGateway gateway.TokenGateway
}

func NewProblemController(usecase usecase.ProblemUseCase, tokenGateway gateway.TokenGateway) ProblemController {
	return ProblemController{usecase: usecase, tokenGateway: tokenGateway}
}

func (controller *ProblemController) GetAllProblems(ctx *gin.Context) {
	result, _ := controller.usecase.GetAllProblems()
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) GetProblemById(ctx *gin.Context) {
	problemId := ctx.Param("id")
	result, _ := controller.usecase.GetProblemByID(ctx, problemId)
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
	result, _ := controller.usecase.CreateProblem(ctx, &body)
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
	result, _ := controller.usecase.UpdateProblem(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) GetAllProblemsByUserId(ctx *gin.Context) {
	userId, _ := controller.tokenGateway.GetUserId(ctx)
	result, err := controller.usecase.GetAllProblemsByUserId(ctx, userId)
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
	result, _ := controller.usecase.DeleteProblem(ctx, body.ID)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) DeleteAllProblemsByUserId(ctx *gin.Context) {
	userId, _ := controller.tokenGateway.GetUserId(ctx)
	result, _ := controller.usecase.DeleteAllProblemsByUserId(ctx, userId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}
