package usecase

import (
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type GetAllProblemsUsecase struct {
	problemRepository  repository.ProblemRepository
}

func NewGetAllProblemsUsecase(
	problemRepository repository.ProblemRepository,
) GetAllProblemsUsecase {
	return GetAllProblemsUsecase{
		problemRepository:  problemRepository,
	}
}



func (usecase *GetAllProblemsUsecase) Execute() ([]entity.Problem, error) {
	return usecase.problemRepository.GetAll()
}