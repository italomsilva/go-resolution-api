package usecase

import (
	"fmt"
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/application/problems/model"
	"go-resolution-api/gateway"
	"go-resolution-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) CreateProblem(ctx *gin.Context, input *dto.ReqCreateProblem) (*model.Problem, error) {
	problem := model.NewProblem()
	id := gateway.GenerateUUID()
	userId, exists := ctx.Get("userId")
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Login again")
		return nil, fmt.Errorf("token not found")
	}
	problem.ID = id
	problem.UserID = fmt.Sprintf("%v", userId)
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
