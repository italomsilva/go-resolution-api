package usecase

import "go-resolution-api/application/solutions/repository"

type SolutionUseCase struct{
	solutionRepository *repository.SolutionRepository
}

func NewSolutionUseCase(solutionRepository *repository.SolutionRepository) SolutionUseCase {
	return SolutionUseCase{solutionRepository: solutionRepository}
}