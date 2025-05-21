package repository

import (
	"database/sql"

	"github.com/lucasschilin/schily-users-api/internal/model"
)

type UserRepository interface {
	CreateWithTX(
		tx *sql.Tx, newID string, username string,
	) (*model.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateWithTX(
	tx *sql.Tx, newID string, username string,
) (*model.User, error) {
	var user model.User

	if err := tx.QueryRow(
		"INSERT INTO users (id, username) VALUES ($1, $2) RETURNING id, username, created_at, updated_at",
		newID,
		username,
	).Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
