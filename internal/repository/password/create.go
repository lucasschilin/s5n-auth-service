package password

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

func (r *repository) CreateWithTX(
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
