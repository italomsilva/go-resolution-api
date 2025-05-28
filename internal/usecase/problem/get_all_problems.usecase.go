package usecase

import "go-resolution-api/internal/domain/entity"


func (usecase *ProblemUseCase) GetAllProblems() ([]entity.Problem, error) {
	return usecase.problemRepository.GetAllProblems()
}