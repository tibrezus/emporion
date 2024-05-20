package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tibrezus/emporion/config"
)

func NewPostgresStorage() (*sql.DB, error) {
	envs := config.Envs

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envs.Host, envs.Port, envs.User, envs.Password, envs.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}
