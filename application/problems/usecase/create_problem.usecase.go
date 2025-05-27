package usecase

import (
	"fmt"
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/application/problems/model"
	"go-resolution-api/gateway"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) CreateProblem(ctx *gin.Context, input *dto.ReqCreateProblem) (*model.Problem, error) {
	problem := model.NewProblem()
	id := gateway.GenerateUUID()
	userId, exists := utils.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return nil, fmt.Errorf("authentication required")
	}
	problem.ID = id
	problem.UserID = userId
	problem.Title = input.Title
	problem.Description = input.Description
	problem.Location = input.Location
	if input.Status != nil {
		problem.Status = *input.Status
	}

	result, err := usecase.problemRepository.CreateProblem(&problem)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Create Problem Error")
		return nil, err
	}

	return result, nil

}
