package useremail

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

func (r *repository) GetByAddress(
	address *string,
) (*model.UserEmail, error) {
	var userEmail model.UserEmail

	if err := r.DB.QueryRow("SELECT * FROM users_emails WHERE address = $1",
		address).Scan(
		&userEmail.ID, &userEmail.User, &userEmail.Address,
		&userEmail.VerifyToken, &userEmail.VerifiedAt,
		&userEmail.CreatedAt, &userEmail.UpdatedAt, &userEmail.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &userEmail, nil
}
