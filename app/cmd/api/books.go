package main

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
)

// Обработчик для создания книги
func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title  string `json:"title"`
        Author string `json:"author"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, "Invalid JSON")
        return
    }

    // Здесь логика сохранения в БД
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(input)
}

// Обработчик для получения книги по ID
func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil || id < 1 {
        app.errorResponse(w, r, http.StatusNotFound, "Invalid ID")
        return
    }

    // Здесь логика получения книги из БД
    book := map[string]interface{}{
        "id":     id,
        "title":  "Example Book",
        "author": "Author Name",
    }
    
    json.NewEncoder(w).Encode(book)
}