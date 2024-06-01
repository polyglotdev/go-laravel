package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before routes

	//  add routes
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/jet", func(w http.ResponseWriter, r *http.Request) {
		a.App.Render.JetPage(w, r, "testjet", nil, nil)
	})

	// static routes
	fileSever := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileSever))

	return a.App.Routes
}
