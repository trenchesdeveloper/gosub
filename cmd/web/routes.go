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

	mux.Use(app.SessionLoad)

	mux.Get("/", app.HomePage)
	mux.Get("/login", app.Login)
	mux.Post("/login", app.PostLogin)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegister)
	mux.Get("/activate-account", app.ActivateAccount)

	return mux
}
