package movie

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/joho/godotenv"
	"github.com/wlun/movie-api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Repository struct{}

const dbName = "movie"
const trendingCollection = "trending"
const imageBaseUrl = "https://image.tmdb.org/t/p/w500"

func (r Repository) SaveMovieData() ([]byte, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TmbdUrlTrendingMovieWeek := os.Getenv("TMDB")

	response, err := http.Get(TmbdUrlTrendingMovieWeek)
	if err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			return nil, errors.New(err.Error())
		} else {
			if err := r.saveToDatabase(data); err != nil {
				return nil, errors.New(err.Error())
			}
			return data, nil
		}
	}
}

func (r Repository) GetTrendingMovies(keyword string) (TrendingMovies, error) {
	var trendingMovies TrendingMovies
	dbClient := database.Mongo
	collection := dbClient.Database(dbName).Collection(trendingCollection)
	filter := bson.D{}
	if keyword != "" {
		filter = bson.D{{
			"title",
			primitive.Regex{Pattern: keyword, Options: " i"},
		}}
	}
	if cursor, err := collection.Find(context.TODO(), filter); err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	} else {
		for cursor.Next(context.TODO()) {
			var trendingMovie TrendingMovie
			err := cursor.Decode(&trendingMovie)
			if err != nil {
				return nil, errors.New(err.Error())
			} else {
				modifiedMovie := trendingMovie
				modifiedMovie.BackdropPath = imageBaseUrl + "/" + trendingMovie.BackdropPath
				modifiedMovie.PosterPath = imageBaseUrl + "/" + trendingMovie.PosterPath
				trendingMovies = append(trendingMovies, modifiedMovie)
			}
		}
		if err := cursor.Err(); err != nil {
			log.Println(err)
			return nil, errors.New(err.Error())
		}
		if err := cursor.Close(context.TODO()); err != nil {
			log.Println(err)
			return nil, errors.New(err.Error())
		}
	}
	return trendingMovies, nil
}

func (r Repository) saveToDatabase(data []byte) error {
	dbClient := database.Mongo
	collection := dbClient.Database(dbName).Collection(trendingCollection)

	var response Response
	err := json.Unmarshal(data, &response)
	if err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}
	trendingMovie := response.Results
	var movies []interface{}
	for _, movie := range trendingMovie {
		movies = append(movies, movie)
	}
	_, err = collection.InsertMany(context.TODO(), movies)
	if err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}
	return nil
}
