package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/patrykp4888/golang-social-network/internal/store"
	"net/http"
	"strconv"
)

type CreatePostPayload struct {
	Content string   `json:"content"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJSON(w, *r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Content: payload.Content,
		Title:   payload.Title,
		// TODO: Change UserId after auth
		UserID: 1,
		Tags:   payload.Tags,
	}

	if err := app.store.Posts.Create(r.Context(), post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postId")

	idParamInt, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := app.store.Posts.GetByID(r.Context(), idParamInt)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, post)
}
