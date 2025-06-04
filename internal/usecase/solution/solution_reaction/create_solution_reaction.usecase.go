package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	dto "go-resolution-api/internal/dto/solution_reaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSolutionReactionUsecase struct {
	solutionReactionRepository repository.SolutionReactionRepository
	solutionRepository         repository.SolutionRepository
	tokenGateway               gateway.TokenGateway
	idGeneratorGateway         gateway.IDGeneratorGateway
}

func NewCreateSolutionReactionUsecase(
	solutionReactionRepository repository.SolutionReactionRepository,
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,

) CreateSolutionReactionUsecase {
	return CreateSolutionReactionUsecase{
		solutionReactionRepository: solutionReactionRepository,
		solutionRepository:         solutionRepository,
		tokenGateway:               tokenGateway,
		idGeneratorGateway:         idGeneratorGateway,
	}
}

func (usecase *CreateSolutionReactionUsecase) Execute(ctx *gin.Context, input *dto.CreateSolutionReactionRequest) (*entity.SolutionReaction, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	solution, err := usecase.solutionRepository.GetById(input.SolutionID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return nil, err
	}

	if solution.UserID != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized user")
		return nil, nil
	}

	newId := usecase.idGeneratorGateway.Generate()

	newSolutionReaction := entity.SolutionReaction{
		ID:           newId,
		UserID:       userId,
		SolutionID:   input.SolutionID,
		ReactionType: input.ReactionType,
	}

	result, err := usecase.solutionReactionRepository.Create(&newSolutionReaction)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Create Failed")
		return nil, err
	}

	return result, nil
}
