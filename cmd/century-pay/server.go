package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func initWebServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	http.ListenAndServe(":8080", r)
}
