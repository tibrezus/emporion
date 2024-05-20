package main

import (
	"log"

	"github.com/tibrezus/emporion/cmd/api"
	"github.com/tibrezus/emporion/config"
	"github.com/tibrezus/emporion/db"
)

func main() {

	config.SetEnv()

	db, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
