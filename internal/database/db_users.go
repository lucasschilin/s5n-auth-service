package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/lucasschilin/s5n-auth-service/internal/config"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
)

func ConnectDBUsers(l logger.Logger, config *config.DBUsers) *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		l.Errorf(err, "failed to connect to users_db")
		panic(true)

	}

	if err = db.Ping(); err != nil {
		l.Errorf(err, "users_db not reachable")
		panic(true)

	}

	l.Info("connected to users_db")
	return db
}
