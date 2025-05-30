package infra

import (
	"database/sql"
	"fmt"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type ProblemSectorRepository struct {
	connection *sql.DB
}

func NewProblemSectorRepository(connection *sql.DB) repository.ProblemSectorRepository {
	return &ProblemSectorRepository{connection: connection}
}

func (repository *ProblemSectorRepository) fromDatabase(rows *sql.Rows) []entity.ProblemSector {
	var problemSectorList []entity.ProblemSector
	for rows.Next() {
		var problemSectorObj entity.ProblemSector
		err := rows.Scan(
			&problemSectorObj.ID,
			&problemSectorObj.ProblemID,
			&problemSectorObj.SectorID)
		if err != nil {
			return []entity.ProblemSector{}
		}
		problemSectorList = append(problemSectorList, problemSectorObj)
	}
	return problemSectorList
}

func (repository *ProblemSectorRepository) Create(data *entity.ProblemSector) (*entity.ProblemSector, error) {
	query := `
	INSERT INTO problem_sector
		(id, problem_id, sector_id)
	VALUES
		($1, $2, $3)`
	_, err := repository.connection.Query(query, data.ID, data.ProblemID, data.SectorID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repository *ProblemSectorRepository) Delete(id int) (bool, error) {
	query := `DELETE FROM problem_sector WHERE id = $1`
	_, err := repository.connection.Query(query, id)
	if err != nil {
		return false, err
	}
	return true, err
}

func (repository *ProblemSectorRepository) DeleteAllByProblemId(problemID string) (int, error) {
	query := `DELETE FROM problem_sector WHERE problem_id = $1`
	rows, err := repository.connection.Exec(query, problemID)
	if err != nil {
		return 0, err
	}
	deletedCounter, err := rows.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(deletedCounter), err

}

func (repository *ProblemSectorRepository) GetAll() ([]entity.ProblemSector, error) {
	query := `SELECT * FROM problem_sector`
	rows, err := repository.connection.Query(query)
	if err != nil {
		return []entity.ProblemSector{}, err
	}
	result := repository.fromDatabase(rows)
	return result, nil
}

func (repository *ProblemSectorRepository) GetAllByProblemId(problemID int) ([]entity.ProblemSector, error) {
	query := `SELECT * FROM problem_sector WHERE problem_id = $1`
	rows, err := repository.connection.Query(query, problemID)
	if err != nil {
		return []entity.ProblemSector{}, err
	}
	result := repository.fromDatabase(rows)
	return result, nil
}

func (repository *ProblemSectorRepository) GetById(id int) (*entity.ProblemSector, error) {
	query := `SELECT * FROM problem_sector WHERE id = $1`
	rows, err := repository.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	result := repository.fromDatabase(rows)
	if len(result) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return &result[0], nil
}
