package controller

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution_reaction"
	usecase "go-resolution-api/internal/usecase/solution/solution_reaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SolutionReactionController struct {
	tokenGateway           gateway.TokenGateway
	reactSolutionUsecase usecase.ReactSolutionUsecase
	deleteSolutionReactionUsecase usecase.DeleteSolutionReactionUsecase
}

func NewSolutionReactionController(
	tokenGateway gateway.TokenGateway,
	reactSolutionUsecase usecase.ReactSolutionUsecase,
	deleteSolutionReactionUsecase usecase.DeleteSolutionReactionUsecase,
) SolutionReactionController {
	return SolutionReactionController{
		tokenGateway:           tokenGateway,
		reactSolutionUsecase: reactSolutionUsecase,
		deleteSolutionReactionUsecase: deleteSolutionReactionUsecase,
	}
}

func (controller *SolutionReactionController) CreateSolutionReaction(ctx *gin.Context) {
	body := dto.CreateSolutionReactionRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	result, _ := controller.reactSolutionUsecase.Execute(ctx, &body)
	if result != nil {
		response.SendSucess(ctx, http.StatusCreated, result, "")
	}
}

func (controller *SolutionReactionController) DeleteSolutionReaction(ctx *gin.Context) {
	body := dto.DeleteSolutionReactionRequest{}
	err := ctx.BindJSON(&body)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	result, _ := controller.deleteSolutionReactionUsecase.Execute(ctx, body.ID)
	if result != nil {
		response.SendSucess(ctx, http.StatusOK, result, "")
	}
}