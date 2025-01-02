package main

import (
	"github.com/patrykp4888/golang-social-network/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8090"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
