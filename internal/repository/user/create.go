package user

import (
	"database/sql"

	"github.com/aidarkhanov/nanoid"
	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

func (r *repository) CreateWithTX(
	tx *sql.Tx, username *string,
) (*model.User, error) {
	var newUser model.User

	newID := nanoid.New()

	if err := tx.QueryRow(
		"INSERT INTO users (id, username) VALUES ($1, $2) RETURNING id, username, created_at, updated_at",
		newID, username,
	).Scan(&newUser.ID, &newUser.Username, &newUser.CreatedAt, &newUser.UpdatedAt); err != nil {
		return nil, err
	}

	return &newUser, nil
}
