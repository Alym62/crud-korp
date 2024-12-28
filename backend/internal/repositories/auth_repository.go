package repositories

import (
	"database/sql"

	"github.com/Alym62/crud-korp/internal/models"
)

type AuthRepository struct {
	connection *sql.DB
}

func NewAuthRepository(connection *sql.DB) AuthRepository {
	return AuthRepository{
		connection: connection,
	}
}

func (ar *AuthRepository) Login(email string) (*models.User, error) {
	var u models.User

	query, err := ar.connection.Prepare(
		"SELECT id, email, password, position, role, created_at, updated_at, removed FROM users " +
			"WHERE removed = false AND email = $1")

	if err != nil {
		return nil, err
	}

	err = query.QueryRow(email).Scan(
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
