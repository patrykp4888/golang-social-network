package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/patrykp4888/golang-social-network/config"
	"github.com/patrykp4888/golang-social-network/internal/store"
)

type application struct {
	config config.Config
	store  store.Storage
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	// context timeout
	router.Use(middleware.Timeout(time.Second * 60))

	router.Route("/v1", func(router chi.Router) {
		router.Get("/health", app.healthCheckHandler)

		router.Route("/posts", func(router chi.Router) {
			router.Post("/create", app.createPostHandler)

			router.Route("/{postId}", func(router chi.Router) {
				router.Get("/", app.getPostHandler)
			})
		})
	})

	return router
}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         app.config.Address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.Address)

	return server.ListenAndServe()
}
