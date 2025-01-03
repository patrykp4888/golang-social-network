package main

import (
	"log"

	"github.com/patrykp4888/golang-social-network/config"
	"github.com/patrykp4888/golang-social-network/internal/store"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	storage := store.NewStorage(nil)

	app := &application{
		config: *cfg,
		store:  *storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
