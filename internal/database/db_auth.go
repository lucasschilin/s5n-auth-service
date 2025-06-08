package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/lucasschilin/s5n-auth-service/internal/config"
)

func ConnectDBAuth(config *config.DBAuth) *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to auth db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("auth db not reachable: %v", err)
	}

	return db
}
