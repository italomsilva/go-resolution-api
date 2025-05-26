package usecase

import "go-resolution-api/application/problems/model"

func (usecase *ProblemUseCase) GetAllProblems() ([]model.Problem, error) {
	return usecase.problemRepository.GetAllProblems()
}