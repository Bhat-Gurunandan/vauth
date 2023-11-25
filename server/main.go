package main

import (
	"log"
	"net/http"
	"vauth/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", handler.Login)

	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r)) // go run app.go -port=:3000
}
