package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rizbo-dev/social-api/internal/env"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
