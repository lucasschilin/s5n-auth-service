package repository

import (
	"database/sql"

	"github.com/lucasschilin/schily-users-api/internal/model"
)

type PasswordRepository interface {
	CreateWithTX(
		tx *sql.Tx, userID *string, password string,
	) (*model.Password, error)
}

type passwordRepository struct {
	DB *sql.DB
}

func NewPasswordRepository(db *sql.DB) PasswordRepository {
	return &passwordRepository{
		DB: db,
	}
}

func (r *passwordRepository) CreateWithTX(
	tx *sql.Tx, userID *string, password string,
) (*model.Password, error) {
	var newPassword model.Password

	if err := tx.QueryRow(
		`INSERT INTO passwords ("user", password) VALUES ($1, $2) RETURNING *`,
		userID, password,
	).Scan(
		&newPassword.User,
		&newPassword.Password,
		&newPassword.CreatedAt,
		&newPassword.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &newPassword, nil
}
