package infra

import (
	"database/sql"
	"fmt"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type SectorRepository struct {
	connection *sql.DB
}

func NewSectorRepository(connection *sql.DB) repository.SectorRepository {
	return &SectorRepository{connection: connection}
}

func (repository *SectorRepository) fromDatabase(rows *sql.Rows) []entity.Sector {
	var sectorList []entity.Sector
	for rows.Next() {
		var sectorObj entity.Sector
		err := rows.Scan(
			&sectorObj.ID,
			&sectorObj.Name,
			&sectorObj.Description)
		if err != nil {
			return []entity.Sector{}
		}
		sectorList = append(sectorList, sectorObj)
	}
	return sectorList
}
func (repository *SectorRepository) GetAll() ([]entity.Sector, error) {
	query := `SELECT * FROM sector`
	rows, err := repository.connection.Query(query)
	if err != nil {
		return []entity.Sector{}, err
	}
	result := repository.fromDatabase(rows)
	return result, nil
}

func (repository *SectorRepository) GetById(id int) (*entity.Sector, error) {
	query := `SELECT * FROM sector WHERE id = $1`
	rows, err := repository.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	result := repository.fromDatabase(rows)
	if len(result) == 0 {
		return nil, fmt.Errorf("sector not found")
	}

	return &result[0], nil
}

func (repository *SectorRepository) GetByIds(idList []int) ([]entity.Sector, error) {
	if len(idList) == 0 {
		return []entity.Sector{}, nil
	}

	var placeholders string
	args := make([]any, len(idList))
	for i, id := range idList {
		placeholders += fmt.Sprintf("$%d,", i+1)
		args[i] = id
	}
	placeholders = placeholders[:len(placeholders)-1]

	query := fmt.Sprintf(`SELECT * FROM sector WHERE id IN (%s)`, placeholders)

	rows, err := repository.connection.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := repository.fromDatabase(rows)
	return result, nil
}
