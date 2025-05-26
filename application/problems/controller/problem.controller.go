package controller

import (
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/application/problems/usecase"
	"go-resolution-api/response"
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
	}
	result, _ := controller.usecase.CreateProblem(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "")
	}
}


