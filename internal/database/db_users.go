package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/lucasschilin/schily-users-api/internal/config"
)

var DBUsers *sql.DB

func ConnectDBUsers(config *config.DBUsers) {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to users db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("users db not reachable: %v", err)
	}

	DBUsers = db
}
