package usecase

import (
	"fmt"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/gateway"
	"go-resolution-api/internal/domain/repository"
	"go-resolution-api/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProblemsByUserIdUsecase struct {
	problemRepository  repository.ProblemRepository
	tokenGateway       gateway.TokenGateway
}

func NewGetAllProblemsByUserIdUsecase(
	problemRepository repository.ProblemRepository,
	tokenGateway gateway.TokenGateway,
) GetAllProblemsByUserIdUsecase {
	return GetAllProblemsByUserIdUsecase{
		problemRepository:  problemRepository,
		tokenGateway:       tokenGateway,
	}
}


func (usecase *GetAllProblemsByUserIdUsecase) Execute(ctx *gin.Context, userId string) ([]entity.Problem, error) {
	problems := []entity.Problem{}

	userIdToken, _ := usecase.tokenGateway.GetUserId(ctx)
	if  userIdToken != userId {
		response.SendError(ctx, http.StatusUnauthorized, "Unauthorized User")
		return problems, fmt.Errorf("unauthorized user")
	}

	problems, err := usecase.problemRepository.GetAllByUserId(userId)
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "Fetch Users Failed")
		return problems, err
	}
	return problems, nil

}
