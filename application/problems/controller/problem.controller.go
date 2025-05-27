package controller

import (
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/application/problems/usecase"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemController struct {
	usecase *usecase.ProblemUseCase
}

func NewProblemController(usecase *usecase.ProblemUseCase) ProblemController {
	return ProblemController{usecase: usecase}
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
	body := dto.NewReqCreateProblem()
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
	body := dto.NewReqUpdateProblem()
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
	userId, _ := utils.GetUserId(ctx)
	result, err := controller.usecase.GetAllProblemsByUserId(ctx, userId)
	if err == nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}

func (controller *ProblemController) DeleteProblem(ctx *gin.Context) {
	body := dto.NewReqDeleteProblem()
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
	userId, _ := utils.GetUserId(ctx)
	result, _ := controller.usecase.DeleteAllProblemsByUserId(ctx, userId)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}
