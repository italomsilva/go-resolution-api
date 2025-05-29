package infra

import (
	"database/sql"
	"fmt"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type ProblemRepository struct {
	connection *sql.DB
}

func NewProblemRepository(connection *sql.DB) repository.ProblemRepository {
	return &ProblemRepository{connection: connection}
}

func (problemRepository *ProblemRepository) fromDatabase(rows *sql.Rows) []entity.Problem {
	var problemList []entity.Problem
	for rows.Next() {
		var problemObj entity.Problem
		err := rows.Scan(
			&problemObj.ID,
			&problemObj.Title,
			&problemObj.Description,
			&problemObj.Location,
			&problemObj.Status,
			&problemObj.CreatedAt,
			&problemObj.UserID)
		if err != nil {
			fmt.Println(err)
			return []entity.Problem{}
		}
		problemList = append(problemList, problemObj)
	}
	return problemList
}

func (problemRepository *ProblemRepository) GetAll() ([]entity.Problem, error) {
	query := `SELECT * FROM problem`
	rows, err := problemRepository.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []entity.Problem{}, err
	}

	result := problemRepository.fromDatabase(rows)
	if len(result) == 0 {
		fmt.Println("Error fetch Users")
		return []entity.Problem{}, err
	}

	return result, nil
}

func (problemRepository *ProblemRepository) GetAllByUserId(userID string) ([]entity.Problem, error) {
	query := `SELECT * FROM problem WHERE user_id = $1`
	rows, err := problemRepository.connection.Query(query, userID)
	if err != nil {
		fmt.Println(err)
		return []entity.Problem{}, err
	}

	result := problemRepository.fromDatabase(rows)
	if len(result) == 0 {
		fmt.Println("Error fetch Users")
		return []entity.Problem{}, err
	}

	return result, nil
}

func (problemRepository *ProblemRepository) GetById(id string) (*entity.Problem, error) {
	query := `SELECT * FROM problem WHERE id = $1`
	rows, err := problemRepository.connection.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := problemRepository.fromDatabase(rows)
	if len(result) == 0 {
		return nil, fmt.Errorf("problem not found")
	}

	return &result[0], nil
}

func (problemRepository *ProblemRepository) Create(data *entity.Problem) (*entity.Problem, error) {
	query := `
	INSERT INTO problem
		(id, title, description, location, status, created_at, user_id)
	VALUES
	 ($1, $2, $3, $4, $5, $6, $7)`
	_, err := problemRepository.connection.Query(query,
		data.ID,
		data.Title,
		data.Description,
		data.Location,
		data.Status,
		data.CreatedAt,
		data.UserID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (problemRepository *ProblemRepository) Update(id string, data *entity.Problem) (*entity.Problem, error) {
	query := `
	UPDATE problem
	SET
		title = $1,
		description = $2,
		location = $3,
		status = $4
	WHERE id = $5`
	_, err := problemRepository.connection.Query(query,
		data.Title,
		data.Description,
		data.Location,
		data.Status,
		data.ID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (problemRepository *ProblemRepository) Delete(id string) (bool, error) {
	query := `DELETE FROM problem WHERE id = $1`
	_, err := problemRepository.connection.Query(query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (problemRepository *ProblemRepository) DeleteAllByUserId(userId string) (int, error) {
	query := `DELETE FROM problem WHERE user_id = $1`
	rows, err := problemRepository.connection.Exec(query, userId)
	if err != nil {
		return 0, err
	}
	deletedCounter, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(deletedCounter), nil
}
