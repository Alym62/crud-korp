package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Alym62/crud-korp/internal/models"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetList() ([]models.User, error) {
	query := "SELECT id, email, position, role, created_at, updated_at, removed FROM users WHERE removed = false"
	rows, err := ur.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []models.User{}, err
	}

	var userList []models.User
	var user models.User

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Position,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Removed,
		)

		if err != nil {
			fmt.Println(err)
			return []models.User{}, err
		}

		userList = append(userList, user)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) Create(user *models.User) (models.User, error) {
	var u models.User

	query, err := ur.connection.Prepare("INSERT INTO users (email, password, position, role, created_at, updated_at, removed)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, email, position, role, created_at, updated_at, removed")

	if err != nil {
		return models.User{}, err
	}

	err = query.QueryRow(
		user.Email,
		user.Password,
		user.Position,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
		user.Removed,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Position,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Removed,
	)
	if err != nil {
		return models.User{}, err
	}

	query.Close()

	return u, nil
}

func (ur *UserRepository) GetById(id uint) (*models.User, error) {
	var u models.User

	query, err := ur.connection.Prepare(
		"SELECT id, email, password, position, role, created_at, updated_at, removed FROM users " +
			"WHERE removed = false AND id = $1")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(id).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Position,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Removed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &u, nil
}

func (ur *UserRepository) DeleteById(id uint) (*models.User, error) {
	var u models.User

	query, err := ur.connection.Prepare(
		"UPDATE users SET removed = true, updated_at = $1 " +
			"WHERE removed = false AND id = $2 RETURNING id, email, position, role, created_at, updated_at, removed")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(time.Now(), id).Scan(
		&u.ID,
		&u.Email,
		&u.Position,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Removed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &u, nil
}

func (ur *UserRepository) Update(id uint, user *models.User) (*models.User, error) {
	var u models.User

	query, err := ur.connection.Prepare(
		"UPDATE users SET email = $1, password = $2, position = $3, role = $4, updated_at = $5 " +
			"WHERE removed = false AND id = $6 RETURNING id, email, position, role, created_at, updated_at, removed")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(user.Email, user.Password, user.Position, user.Role, time.Now(), id).Scan(
		&u.ID,
		&u.Email,
		&u.Position,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Removed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &u, nil
}
