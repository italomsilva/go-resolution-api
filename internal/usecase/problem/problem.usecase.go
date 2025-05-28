package usecase

import (
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
)

type ProblemUseCase struct {
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
	idGeneratorGateway gateway.IDGeneratorGateway
}

func NewProblemUseCase(
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
	idGeneratorGateway gateway.IDGeneratorGateway,
) ProblemUseCase {
	return ProblemUseCase{
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
		idGeneratorGateway: idGeneratorGateway,
	}
}
