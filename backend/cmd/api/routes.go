package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			data := map[string]string{
				"status":  "ok",
				"version": app.version,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})
	})

	return mux
}
