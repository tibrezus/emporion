package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tibrezus/emporion/cmd/api"
	"github.com/tibrezus/emporion/config"
	"github.com/tibrezus/emporion/db"
)

func main() {
	config.initConfig()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Envs.User,
		config.Envs.Password,
		config.Envs.Host,
		config.Envs.Port,
		config.Envs.DBName)

	ctx := context.Background()
	dbPool, err := db.NewPostgresStorage(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}

	server, err := api.NewAPIServer(":8080", dbPool)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
