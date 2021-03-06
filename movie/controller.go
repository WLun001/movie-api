package movie

import (
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) SaveMovieData(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	response, err := c.Repository.SaveMovieData()
	if err != nil {
		errByte, _ := json.Marshal(Message{err.Error()})
		writeResponse(&w, errByte, http.StatusInternalServerError)
	} else {
		jsonResponse, _ := json.Marshal(response)
		writeResponse(&w, jsonResponse)
	}
}

func (c Controller) TrendingMovie(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	title := r.URL.Query().Get("title")
	trendingMovies, err := c.Repository.GetTrendingMovies(title)
	if err != nil {
		errByte, _ := json.Marshal(Message{err.Error()})
		writeResponse(&w, errByte, http.StatusInternalServerError)
	} else if len(trendingMovies) <= 0 {
		errByte, _ := json.Marshal(Message{"no data found"})
		writeResponse(&w, errByte, http.StatusNotFound)
	} else {
		data, _ := json.Marshal(trendingMovies)
		writeResponse(&w, data)
	}
}

func (c Controller) DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	if err := c.Repository.DeleteAll(); err != nil {
		errByte, _ := json.Marshal(Message{err.Error()})
		writeResponse(&w, errByte, http.StatusInternalServerError)
	} else {
		data, _ := json.Marshal(Message{"deleted all data"})
		writeResponse(&w, data)
	}
}

func logRequest(r *http.Request) {
	log.Printf("API (request) - host: %s, method: %s, path: %s, query: %s, user agent: %s",
		r.Host,
		r.Method,
		r.URL.Path,
		r.URL.RawQuery,
		r.UserAgent())
}

func writeResponse(w *http.ResponseWriter, data []byte, statusCode ...int) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")

	if len(statusCode) == 0 {
		(*w).WriteHeader(http.StatusOK)
	} else {
		(*w).WriteHeader(statusCode[0])
	}

	if _, err := (*w).Write(data); err != nil {
		log.Printf("API (write) - error on writing data: %s", err)
	}
}
