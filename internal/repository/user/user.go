package user

import (
	"database/sql"

	"github.com/lucasschilin/s5n-auth-service/internal/model"
)

type Repository interface {
	GetByID(id *string) (*model.User, error)
	GetByUsername(username *string) (*model.User, error)
	CreateWithTX(
		tx *sql.Tx, username *string,
	) (*model.User, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		DB: db,
	}
}
