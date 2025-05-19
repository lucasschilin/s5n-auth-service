package repository

import (
	"database/sql"

	"github.com/lucasschilin/schily-users-api/internal/model"
)

type UserRepository interface {
	GetByEmailAddress(address string) (*model.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetByEmailAddress(address string) (*model.User, error) {
	return nil, nil
}
