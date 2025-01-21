package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func readJSON(w http.ResponseWriter, r http.Request, data any) error {
	maxBytes := 1_048_578 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type errorResponse struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &errorResponse{Error: message})
}
