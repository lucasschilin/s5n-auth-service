package password

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

func (r *repository) GetByUser(
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
