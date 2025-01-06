package main

import (
	"github.com/patrykp4888/golang-social-network/config"
	"github.com/patrykp4888/golang-social-network/internal/db"
	"github.com/patrykp4888/golang-social-network/internal/store"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	dbConn, err := db.NewConnection(cfg.DB.Address, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTime)
	if err != nil {
		log.Fatalf("unable to reach db connection: %e", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatalf("unable to close db connection")
		}
	}()

	log.Println("database connection pool established")

	storage := store.NewStorage(dbConn)

	app := &application{
		config: *cfg,
		store:  *storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
