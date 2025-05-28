package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API     *API
	DBUsers *DBUsers
	DBAuth  *DBAuth
	JWT     *JWT
}

type API struct {
	Host string
	Port string
}

type DBUsers struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type DBAuth struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}
type JWT struct {
	SecretKey string
}

func Load() *Config {
	_ = godotenv.Load()

	APIConfig := API{
		Host: os.Getenv("API_HOST"),
		Port: os.Getenv("API_PORT"),
	}
	DBUsersConfig := DBUsers{
		Username: os.Getenv("DB_USERS_USERNAME"),
		Password: os.Getenv("DB_USERS_PASSWORD"),
		Host:     os.Getenv("DB_USERS_HOST"),
		Port:     os.Getenv("DB_USERS_PORT"),
		Name:     os.Getenv("DB_USERS_NAME"),
	}
	DBAuthConfig := DBAuth{
		Username: os.Getenv("DB_AUTH_USERNAME"),
		Password: os.Getenv("DB_AUTH_PASSWORD"),
		Host:     os.Getenv("DB_AUTH_HOST"),
		Port:     os.Getenv("DB_AUTH_PORT"),
		Name:     os.Getenv("DB_AUTH_NAME"),
	}
	JWTConfig := JWT{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}

	return &Config{
		API:     &APIConfig,
		DBUsers: &DBUsersConfig,
		DBAuth:  &DBAuthConfig,
		JWT:     &JWTConfig,
	}

}
