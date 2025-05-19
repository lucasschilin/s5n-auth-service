package repository

import (
	"database/sql"

	"github.com/lucasschilin/schily-users-api/internal/model"
)

type PasswordRepository interface {
}

type passwordRepository struct {
	DB *sql.DB
}

func NewPasswordRepository(db *sql.DB) PasswordRepository {
	return &passwordRepository{
		DB: db,
	}
}

func (r *passwordRepository) GetByEmailAddress(address string) (*model.Password, error) {
	return nil, nil
}
