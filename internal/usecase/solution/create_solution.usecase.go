package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
)

type CreateSolutionUsecase struct {
	solutionRepository repository.SolutionRepository
	tokenGateway       gateway.TokenGateway
	idGeneratorGateway gateway.IDGeneratorGateway
}

func NewCreateSolutionUsecase(
	solutionRepository repository.SolutionRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,

) CreateSolutionUsecase {
	return CreateSolutionUsecase{
		solutionRepository: solutionRepository,
		tokenGateway: tokenGateway,
		idGeneratorGateway: idGeneratorGateway,
	}
}

func (usecase *CreateSolutionUsecase) Execute(){
	
}
