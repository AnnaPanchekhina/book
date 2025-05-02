package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}


	err = app.writeJSON(w, http.StatusCreated, envelope{"book": input}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		app.errorResponse(w, r, http.StatusNotFound, "Invalid ID parameter")
		return
	}

	book := envelope{
		"id":     id,
		"title":  "Example Book",
		"author": "Author Name",
	}

	err = app.writeJSON(w, http.StatusOK, book, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}