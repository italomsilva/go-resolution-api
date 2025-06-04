package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteSolutionReactionUsecase struct {
	solutionReactionRepository repository.SolutionReactionRepository
	solutionRepository         repository.SolutionRepository
	tokenGateway               gateway.TokenGateway
}

func NewDeleteSolutionReactionUsecase(
	solutionReactionRepository repository.SolutionReactionRepository,
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,

) DeleteSolutionReactionUsecase {
	return DeleteSolutionReactionUsecase{
		solutionReactionRepository: solutionReactionRepository,
		solutionRepository:         solutionRepository,
		tokenGateway:               tokenGateway,
	}
}

func (usecase *DeleteSolutionReactionUsecase) Execute(ctx *gin.Context, id string) (*bool, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	solutionReaction, err := usecase.solutionReactionRepository.GetByID(id)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Solution reaction not found")
		return nil, err
	}

	if solutionReaction.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, nil
	}

	_, err = usecase.solutionRepository.GetById(solutionReaction.SolutionID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return nil, err
	}

	result, err := usecase.solutionReactionRepository.Delete(id)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Create Failed")
		return nil, err
	}

	return &result, nil
}
