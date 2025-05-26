package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/lucasschilin/schily-users-api/test/helper"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDBUsers struct {
	DB        *sql.DB
	Terminate func()
}

func SetupTestConnectDBUsers(ctx context.Context) (*TestDBUsers, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:17",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "db_users",
		},
		// WaitingFor: wait.ForListeningPort("5432/tcp"),
		WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
			return fmt.Sprintf(
				"postgres://testuser:testpass@localhost:%s/db_users?sslmode=disable",
				port.Port(),
			)
		}).WithStartupTimeout(30 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"postgres://testuser:testpass@%s/db_users?sslmode=disable",
		endpoint,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = helper.RunMigrationsUp(
		dsn,
		"./../../migrations/db_users/",
	)
	if err != nil {
		return nil, err
	}

	terminate := func() {
		db.Close()
		container.Terminate(ctx)
	}

	return &TestDBUsers{
		DB:        db,
		Terminate: terminate,
	}, nil

}

func (t *TestDBUsers) Close() {
	t.Terminate()
}
