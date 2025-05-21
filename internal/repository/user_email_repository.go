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
	var userEmail model.UserEmail

	if err := r.DB.QueryRow("SELECT * FROM users_emails WHERE address = $1",
		address).Scan(&userEmail); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &userEmail, nil

}
