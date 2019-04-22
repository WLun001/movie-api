package movie

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/joho/godotenv"
	"github.com/wlun/movie-api/database"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Repository struct{}

const dbName = "movie"
const trendingCollection = "trending"

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
