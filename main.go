package main

import (
	"github.com/gorilla/handlers"
	"github.com/wlun/movie-api/movie"
	"log"
	"net/http"
)

func main() {

	router := movie.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})
	allowedHeader := handlers.AllowedHeaders([]string{"Accept", "Content-Type"})

	log.Fatal(http.ListenAndServe(":5000",
		handlers.CORS(allowedOrigins, allowedMethods, allowedHeader)(router)))
}
