package password

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

type Repository interface {
	CreateWithTX(
		tx *sql.Tx, userID string, password string,
	) (*model.Password, error)
	GetByUser(userID string) (*model.Password, error)
	UpdateByUser(userID string, newPassword string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		DB: db,
	}
}
