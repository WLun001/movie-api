package movie

import (
	"errors"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Repository struct{}

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
			return data, nil
		}
	}
}
