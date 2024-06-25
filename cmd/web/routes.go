package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)

func (app *Config) Routes() http.Handler {
	mux := chi.NewRouter()

	// Set middleware
	mux.Use(middleware.Recoverer)

	mux.HandleFunc("/", app.HomePage)

	return mux
}
