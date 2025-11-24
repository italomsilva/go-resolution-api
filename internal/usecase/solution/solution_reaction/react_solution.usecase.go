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

type ReactSolutionUsecase struct {
	solutionReactionRepository repository.SolutionReactionRepository
	solutionRepository         repository.SolutionRepository
	tokenGateway               gateway.TokenGateway
	idGeneratorGateway         gateway.IDGeneratorGateway
}

func NewReactSolutionUsecase(
	solutionReactionRepository repository.SolutionReactionRepository,
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,

) ReactSolutionUsecase {
	return ReactSolutionUsecase{
		solutionReactionRepository: solutionReactionRepository,
		solutionRepository:         solutionRepository,
		tokenGateway:               tokenGateway,
		idGeneratorGateway:         idGeneratorGateway,
	}
}

func (usecase *ReactSolutionUsecase) Execute(ctx *gin.Context, input *dto.CreateSolutionReactionRequest) (*entity.SolutionReaction, error) {
	userId, _ := usecase.tokenGateway.GetUserId(ctx)

	solution, err := usecase.solutionRepository.GetById(input.SolutionID)
	if err != nil {
		response.SendError(ctx, http.StatusNotFound, "Solution not found")
		return nil, err
	}

	hasReacted, err := usecase.solutionReactionRepository.GetBySolutionIdAndUserId(solution.ID, userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Internal Server Error")
		return nil, err
	}

	reaction := entity.SolutionReaction{
		UserID:       userId,
		SolutionID:   input.SolutionID,
		ReactionType: input.ReactionType,
	}
	
	if hasReacted == nil {
		newId := usecase.idGeneratorGateway.Generate()
		reaction.ID = newId;

		result, err := usecase.solutionReactionRepository.Create(&reaction)
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "Create Failed")
			return nil, err
		}

		return result, nil
	} else {
		reaction.ID = hasReacted.ID;
		_, err := usecase.solutionReactionRepository.Update(hasReacted.ID, reaction);
		if err != nil {
			response.SendError(ctx, http.StatusInternalServerError, "Update Failed")
			return nil, err
		}
		return &reaction, nil;
	}

}
