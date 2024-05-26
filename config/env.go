package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
}

var Envs EnvConfig

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Envs.Host = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %v", err)
	}
	Envs.Port = uint16(port)
	Envs.User = os.Getenv("DB_USER")
	Envs.Password = os.Getenv("DB_PASSWORD")
	Envs.DBName = os.Getenv("DB_NAME")
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
