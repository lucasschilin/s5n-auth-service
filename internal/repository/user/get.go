package user

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

func (r *repository) GetByID(id *string) (*model.User, error) {
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

func (r *repository) GetByUsername(username *string) (*model.User, error) {
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
