package main

import (
	"net/http"

	"github.com/Rajatdey12/goLangStarter/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes to create all the url mappings for app..
func routes() http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Get("/ws", handlers.Ws)
	router.Get("/about", handlers.About)

	return router
}
