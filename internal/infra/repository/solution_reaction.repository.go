package infra

import (
	"database/sql"
	"go-resolution-api/internal/domain/entity"
	"go-resolution-api/internal/domain/repository"
)

type SolutionReactionRepository struct {
	connection *sql.DB
}

func NewSolutionReactionRepository(connection *sql.DB) repository.SolutionReactionRepository {
	return &SolutionReactionRepository{connection: connection}
}

func (repository *SolutionReactionRepository) fromDatabase(rows *sql.Rows) []entity.SolutionReaction {
	var solutionReactionList []entity.SolutionReaction
	for rows.Next() {
		var solutionReactionObj entity.SolutionReaction
		err := rows.Scan(
			&solutionReactionObj.ID,
			&solutionReactionObj.UserID,
			&solutionReactionObj.SolutionID,
			&solutionReactionObj.ReactionType)
		if err != nil {
			return []entity.SolutionReaction{}
		}
		solutionReactionList = append(solutionReactionList, solutionReactionObj)
	}
	return solutionReactionList
}

func (repository *SolutionReactionRepository) Create(data *entity.SolutionReaction) (*entity.SolutionReaction, error) {
	query := `
	INSERT INTO solution_reation
		(id, user_id, solution_id, reaction_type)
	VALUES
		($1, $2, $3, $4)`
	_, err := repository.connection.Query(query, data.ID, data.UserID, data.SolutionID, data.ReactionType)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repository *SolutionReactionRepository) Delete(id string) (bool, error) {
	query := `DELETE FROM solution_reaction WHERE id = $1`

	_, err := repository.connection.Query(query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repository *SolutionReactionRepository) DeleteAllBySolutionId(solutionId string) (int, error) {
	query := `DELETE FROM solution_reaction WHERE solution_id = $1`

	rows, err := repository.connection.Exec(query, solutionId)
	if err != nil {
		return 0, err
	}

	deletedCounter, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(deletedCounter), nil
}

func (repository *SolutionReactionRepository) GetAllBySolutionId(solutionId string) ([]entity.SolutionReaction, error) {
	query := `SELECT * FROM solution_reaction WHERE solution_id = $1`

	rows, err := repository.connection.Query(query, solutionId)
	if err != nil {
		return []entity.SolutionReaction{}, err
	}

	result := repository.fromDatabase(rows)
	if len(result) == 0 {
		return []entity.SolutionReaction{}, err
	}

	return result, nil
}

func (repository *SolutionReactionRepository) GetByID(id string) (*entity.SolutionReaction, error) {
	query := `SELECT * FROM solution_reaction WHERE id = $1`

	rows, err := repository.connection.Query(query, id)
	if err != nil {
		return nil, err
	}

	result := repository.fromDatabase(rows)
	if len(result) == 0 {
		return nil, err
	}

	return &result[0], nil
}

func (repository *SolutionReactionRepository) GetReactionsBySolutionIdAndUserId(solutionId string, userId string) (int, int, entity.ReactionType, error) {
	query := `SELECT * FROM solution_reaction WHERE solution_id = $1`

	rows, err := repository.connection.Query(query, solutionId)
	if err != nil {
		return 0, 0, entity.ReactionTypeNone, err
	}

	result := repository.fromDatabase(rows)
	if len(result) == 0 {
		return 0, 0, entity.ReactionTypeNone, nil
	}

	myReaction := entity.ReactionTypeNone

	likes := 0
	dislikes := 0

	for _, reaction := range result {
		if reaction.ReactionType == entity.ReactionTypeLike {
			likes++
		} else if reaction.ReactionType == entity.ReactionTypeDeslike {
			dislikes++
		}
		if reaction.UserID == userId {
			myReaction = reaction.ReactionType
		}
	}

	return likes, dislikes, myReaction, nil
}

func (repository *SolutionReactionRepository) GetBySolutionIdAndUserId(solutionId string, userId string) (*entity.SolutionReaction, error) {
	query := `SELECT * FROM solution_reaction WHERE solution_id = $1 AND user_id = $2`

	rows, err := repository.connection.Query(query, solutionId, userId)
	if err != nil {
		return nil, err
	}

	result := repository.fromDatabase(rows)
	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (repository *SolutionReactionRepository) Update(id string, reaction entity.SolutionReaction) (bool, error) {
	query := `
	UPDATE solution_reaction
	SET reaction_type = $1
	WHERE id = $2`

	_, err := repository.connection.Exec(query, reaction.ReactionType, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
