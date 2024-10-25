package main

import "log"

func main() {
	config := config{
		addr: ":8080",
	}
	app := &application{
		config,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
