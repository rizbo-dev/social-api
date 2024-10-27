package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rizbo-dev/social-api/internal/db"
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
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		config.db.addr,
		config.db.maxOpenConnections,
		config.db.maxIdleConnections,
		config.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("database connection pool estaplished")

	store := store.NewStorage(db)

	app := &application{
		config,
		store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
