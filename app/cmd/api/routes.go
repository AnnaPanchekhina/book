// cmd/api/routes.go

package main

import (
   "github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
   router := chi.NewRouter()

   router.Get("/v1/healthcheck", app.healthcheckHandler)
   router.Post("/v1/books", app.createBookHandler)
   router.Get("/v1/books/{id}", app.showBookHandler)

   return router
}