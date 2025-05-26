package usecase

import "go-resolution-api/application/problems/repository"

type ProblemUseCase struct {
	problemRepository *repository.ProblemRepository
}

func NewProblemUseCase(problemRepository *repository.ProblemRepository) ProblemUseCase {
	return ProblemUseCase{problemRepository: problemRepository}
}
