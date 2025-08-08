package useremail

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

type Repository interface {
	GetByAddress(address *string) (*model.UserEmail, error)
	CreateWithTX(
		tx *sql.Tx, userID *string, address *string, verifyToken *string,
	) (*model.UserEmail, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		DB: db,
	}
}
