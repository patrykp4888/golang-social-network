package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.Environment,
		"version": app.config.ApiVersion,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Fatalf("app health check failed")
	}
}
