package main

import (
	"encoding/json"
	"net/http"
)

// envelope - обертка для JSON ответов
type envelope map[string]interface{}

// Добавляем метод writeJSON в структуру application
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	return err
}

// Методы для обработки ошибок
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Printf("server error: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logger.Printf("error writing JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}