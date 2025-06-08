package repository

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

type PasswordRepository interface {
	CreateWithTX(
		tx *sql.Tx, userID string, password string,
	) (*model.Password, error)
	GetByUser(userID string) (*model.Password, error)
	UpdateByUser(userID string, newPassword string) error
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
	tx *sql.Tx, userID string, password string,
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

func (r *passwordRepository) GetByUser(
	userID string,
) (*model.Password, error) {
	var password model.Password

	if err := r.DB.QueryRow(
		`SELECT * FROM passwords WHERE "user" = $1`,
		userID,
	).Scan(
		&password.User,
		&password.Password,
		&password.CreatedAt,
		&password.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return &password, nil
}

func (r *passwordRepository) UpdateByUser(
	userID string, newPassword string,
) error {
	if _, err := r.DB.Exec(
		`UPDATE passwords SET password = $1 WHERE "user" = $2`,
		newPassword, userID,
	); err != nil {
		return err
	}

	return nil
}
