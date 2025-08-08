package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/lucasschilin/s5n-auth-service/internal/config"
	"github.com/lucasschilin/s5n-auth-service/pkg/logger"
)

func ConnectDBAuth(l logger.Logger, config *config.DBAuth) *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.Name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		l.Errorf(err, "failed to connect to auth_db")
		panic(true)

	}

	if err = db.Ping(); err != nil {
		l.Errorf(err, "auth_db not reachable")
		panic(true)

	}

	l.Info("connected to auth_db")
	return db
}
