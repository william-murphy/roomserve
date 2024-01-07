package main

import (
	"log"
	"net/http"

	"roomserve/config"
	"roomserve/database"
	"roomserve/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.Initialize()
	database.Connect()
	r := chi.NewRouter()
	router.Initialize(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
