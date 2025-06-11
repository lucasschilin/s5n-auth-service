package useremail

import (
	"database/sql"

	"github.com/aidarkhanov/nanoid"
	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

func (r *repository) CreateWithTX(
	tx *sql.Tx, userID *string, address *string, verifyToken *string,
) (*model.UserEmail, error) {
	var newUserEmail model.UserEmail

	newID := nanoid.New()

	err := tx.QueryRow(
		`INSERT INTO users_emails (id, "user", address, verify_token) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, "user", address, verify_token, created_at, updated_at`,
		newID, userID, address, verifyToken,
	).Scan(
		&newUserEmail.ID,
		&newUserEmail.User,
		&newUserEmail.Address,
		&newUserEmail.VerifyToken,
		&newUserEmail.CreatedAt,
		&newUserEmail.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newUserEmail, nil

}
