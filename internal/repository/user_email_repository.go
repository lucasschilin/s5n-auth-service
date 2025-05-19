package repository

import (
	"database/sql"

	"github.com/lucasschilin/schily-users-api/internal/model"
)

type UserEmailRepository interface {
	GetByAddress(address string) (*model.UserEmail, error)
}

type userEmailRepository struct {
	DB *sql.DB
}

func NewUserEmailRepository(db *sql.DB) UserEmailRepository {
	return &userEmailRepository{
		DB: db,
	}
}

func (r *userEmailRepository) GetByAddress(address string) (*model.UserEmail, error) {
	return nil, nil
}
