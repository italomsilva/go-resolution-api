package usecase

import (
	"fmt"
	"go-resolution-api/application/problems/dto"
	"go-resolution-api/application/problems/model"
	"go-resolution-api/response"
	"go-resolution-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *ProblemUseCase) UpdateProblem(ctx *gin.Context, input *dto.ReqUpdateProblem) (*model.Problem, error) {
	_, exists := utils.GetUserId(ctx)
	if !exists {
		response.SendError(ctx, http.StatusUnauthorized, "Authentication required")
		return nil, fmt.Errorf("authentication required")
	}

	problem, err := usecase.problemRepository.GetProblemById(input.ID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Problem Not Found")
		return nil, err
	}

	if input.Title != nil && *input.Title != "" {
		problem.Title = *input.Title
	}

	if input.Description != nil {
		problem.Description = *input.Description
	}

	if input.Location != nil && *input.Location != "" {
		problem.Location = *input.Location
	}

	if input.Status != nil {
		problem.Status = *input.Status
	}

	result, err := usecase.problemRepository.UpdateProblem(problem.ID, problem)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "update problem error")
		return nil, err
	}

	return result, nil

}
