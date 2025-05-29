package infra

import (
	"database/sql"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type SolutionRepository struct {
	connection *sql.DB
}

func NewSolutionRepository(databaseConnection *sql.DB) repository.SolutionRepository {
	return &SolutionRepository{connection: databaseConnection}
}
func (repository *SolutionRepository) Delete(id string) (bool, error) {
	panic("unimplemented")
}

func (repository *SolutionRepository) fromDatabase(rows *sql.Rows) ([]entity.Solution, error) {
	var solutionsList []entity.Solution
	for rows.Next() {
		solutionObj := entity.Solution{}
		err := rows.Scan(
			&solutionObj.ID,
			&solutionObj.Title,
			&solutionObj.Description,
			&solutionObj.EstimatedCost,
			&solutionObj.Approved,
			&solutionObj.CreatedAt,
			&solutionObj.ProblemId,
			&solutionObj.UserId,
		)
		if err != nil {
			return []entity.Solution{}, err
		}

		solutionsList = append(solutionsList, solutionObj)
	}

	return solutionsList, nil
}

func (repository *SolutionRepository) GetAllByProblemId(problemId string) ([]entity.Solution, error) {
	query := `SELECT * FROM solution WHERE problem_id = $1`
	rows, err := repository.connection.Query(query, problemId)
	if err != nil {
		return []entity.Solution{}, err
	}

	solutions, err := repository.fromDatabase(rows)
	if err != nil || len(solutions) == 0 {
		return []entity.Solution{}, err
	}

	return solutions, nil

}

func (repository *SolutionRepository) GetById(id string) (*entity.Solution, error) {
	query := `SELECT * FROM solution WHERE id = $1`
	rows, err := repository.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	solutions, err := repository.fromDatabase(rows)
	if err != nil || len(solutions) == 0 {
		return nil, err
	}

	return &solutions[0], nil

}

func (repository *SolutionRepository) Create(data *entity.Solution) (*entity.Solution, error) {
	query := `
	INSERT INTO solution
		(id, title, description, estimated_cost, approved, created_at, problem_id, user_id)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := repository.connection.Query(query,
		data.ID,
		data.Title,
		data.Description,
		data.EstimatedCost,
		data.Approved,
		data.CreatedAt,
		data.ProblemId,
		data.UserId,
	)
	if err != nil {
		return nil, err
	}
	return data, nil

}
