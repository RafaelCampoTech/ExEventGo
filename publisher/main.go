package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Print("Starting publisher service on port " + LoadConfig().HTTPPort)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	RegisterRoutes(r)
	log.Println("Publisher service started on port " + LoadConfig().HTTPPort)
	log.Fatal(http.ListenAndServe(":"+LoadConfig().HTTPPort, r), "Error starting publisher service")
	log.Println("Publisher service stopped")
}
