package movie

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Repository struct{}

const TmbdUrlTrendingMovieWeek = "https://api.themoviedb.org/3/trending/movie/week?api_key=f74026db6e599702db5d73c37ea43aa6"

func (r Repository) SaveMovieData() ([]byte, error) {
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
