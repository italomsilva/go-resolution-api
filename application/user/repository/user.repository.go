package repository

import (
	"database/sql"
	"fmt"
	"go-resolution-api/application/user/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{connection: connection}
}

func (userRepository *UserRepository) fromDatabase(rows *sql.Rows) []model.User {
	var userList []model.User
	for rows.Next() {
		var userObj model.User
		err := rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Email,
			&userObj.Document,
			&userObj.Profile,
			&userObj.Login,
			&userObj.Password,
			&userObj.Token)
		if err != nil {
			fmt.Println(err)
			return []model.User{}
		}
		userList = append(userList, userObj)
	}
	return userList
}

func (userRepository *UserRepository) GetUsers() ([]model.User, error) {
	query := `SELECT * FROM "user"`
	rows, err := userRepository.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}
	defer rows.Close()

	result := userRepository.fromDatabase(rows)
	if len(result) == 0 {
		println("Error fetch users")
		return []model.User{}, err
	}

	return result, nil
}

func (userRepository *UserRepository) GetUserById(id string) (*model.User, error) {
	query := `SELECT * FROM "user" WHERE id = $1`
	rows, err := userRepository.connection.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	result := userRepository.fromDatabase(rows)
	if len(result) == 0 {
		return nil, err
	}

	return &result[0], nil
}

func (userRepository *UserRepository) GetUserByLogin(login string) (*model.User, error) {
	query := `SELECT * FROM "user" WHERE login = $1`
	rows, err := userRepository.connection.Query(query, login)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := userRepository.fromDatabase(rows)

	defer rows.Close()

	if len(result) == 0 {
		println("Error fetch users")
		return nil, err
	}

	return &result[0], nil
}

func (userRepository *UserRepository) GetUserByDocument(document string) (*model.User, error) {
	query := `SELECT * FROM "user" WHERE document = $1`
	rows, err := userRepository.connection.Query(query, document)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := userRepository.fromDatabase(rows)

	defer rows.Close()

	if len(result) == 0 {
		println("Error fetch users")
		return nil, err
	}

	return &result[0], nil
}

func (userRepository *UserRepository) CreateUser(input *model.User) (*model.User, error) {
	query := `INSERT INTO "user" (id, name, email, document, profile, login, password, token)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id`

	var userId string
	err := userRepository.connection.QueryRow(query,
		input.ID,
		input.Name,
		input.Email,
		input.Document,
		input.Profile,
		input.Login,
		input.Password,
		input.Token).Scan(&userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result, err := userRepository.GetUserById(userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func (userRepository *UserRepository) UpdateUser(id string, data *model.User) (*model.User, error) {
	query := `
	UPDATE "user"
	SET name = $1,
		email = $2,
		document = $3,
		profile = $4,
		login = $5,
		password = $6,
		token = $7
	WHERE id = $8
	RETURNING id;
  `

	var userId string
	err := userRepository.connection.QueryRow(query,
		data.Name,
		data.Email,
		data.Document,
		data.Profile,
		data.Login,
		data.Password,
		data.Token,
		id).Scan(&userId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, err

}

func (userRepository *UserRepository) DeleteUser(id string) (bool, error) {
	query := `DELETE FROM "user" WHERE id = $1`

	_, err := userRepository.connection.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return false, err
	}

	return true, nil
}
