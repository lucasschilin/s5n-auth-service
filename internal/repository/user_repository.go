package repository

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

type UserRepository interface {
	GetByID(id *string) (*model.User, error)
	GetByUsername(username *string) (*model.User, error)
	CreateWithTX(
		tx *sql.Tx, username *string,
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

func (r *userRepository) GetByID(id *string) (*model.User, error) {
	var user model.User
	if err := r.DB.QueryRow(
		"SELECT id, username, created_at, updated_at FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByUsername(username *string) (*model.User, error) {
	var user model.User
	if err := r.DB.QueryRow(
		"SELECT id, username, created_at, updated_at FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreateWithTX(
	tx *sql.Tx, username *string,
) (*model.User, error) {
	var newUser model.User

	if err := tx.QueryRow(
		"INSERT INTO users (id, username) VALUES ($1, $2) RETURNING id, username, created_at, updated_at",
		newID(), username,
	).Scan(&newUser.ID, &newUser.Username, &newUser.CreatedAt, &newUser.UpdatedAt); err != nil {
		return nil, err
	}

	return &newUser, nil
}
