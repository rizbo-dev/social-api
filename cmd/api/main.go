package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rizbo-dev/social-api/internal/env"
	"github.com/rizbo-dev/social-api/internal/store"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config,
		store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
