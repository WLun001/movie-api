package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/wlun/movie-api/database"
	"github.com/wlun/movie-api/movie"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"
)

func initDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUri := os.Getenv("DBURi")
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	log.Println("db client created")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("db client connected")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db client ping")

	database.Mongo = client
}

func main() {
	//initDatabase()

	router := movie.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})
	allowedHeader := handlers.AllowedHeaders([]string{"Accept", "Content-Type"})

	log.Fatal(http.ListenAndServe(":5000",
		handlers.CORS(allowedOrigins, allowedMethods, allowedHeader)(router)))
}
