package usecase

import "go-resolution-api/internal/infra"


type SolutionUseCase struct{
	solutionRepository *infra.SolutionRepository
}

func NewSolutionUseCase(solutionRepository *infra.SolutionRepository) SolutionUseCase {
	return SolutionUseCase{solutionRepository: solutionRepository}
}