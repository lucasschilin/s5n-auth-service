package repository

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

type UserEmailRepository interface {
	GetByAddress(address *string) (*model.UserEmail, error)
	CreateWithTX(
		tx *sql.Tx, userID *string, address *string, verifyToken *string,
	) (*model.UserEmail, error)
}

type userEmailRepository struct {
	DB *sql.DB
}

func NewUserEmailRepository(db *sql.DB) UserEmailRepository {
	return &userEmailRepository{
		DB: db,
	}
}

func (r *userEmailRepository) GetByAddress(
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

func (r *userEmailRepository) CreateWithTX(
	tx *sql.Tx, userID *string, address *string, verifyToken *string,
) (*model.UserEmail, error) {
	var newUserEmail model.UserEmail

	err := tx.QueryRow(
		`INSERT INTO users_emails (id, "user", address, verify_token) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, "user", address, verify_token, created_at, updated_at`,
		newID(), userID, address, verifyToken,
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
