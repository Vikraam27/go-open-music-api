package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vikraam27/go-open-music-api/exceptions"
	"github.com/Vikraam27/go-open-music-api/handlers"
	"github.com/gorilla/mux"
)

type ServerError struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
type rootHandler func(http.ResponseWriter, *http.Request) error

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	serverErrorRes := ServerError{
		Status:     "error",
		StatusCode: 500,
		Message:    "Server Error",
	}

	if err == nil {
		return
	}

	log.Printf("An error accured: %v", err)

	clientError, ok := err.(exceptions.ClientError)

	if !ok {
		w.WriteHeader(500)
		blob, _ := json.Marshal(serverErrorRes)
		w.Write(blob)
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		blob, _ := json.Marshal(serverErrorRes)
		w.Write(blob)
		return
	}

	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}

func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.Handle("/albums", rootHandler(handlers.CreateAlbumHandler)).Methods("POST")
	routes.Handle("/albums/{id}", rootHandler(handlers.GetAlbumDetailHandler)).Methods("GET")
	routes.Handle("/albums/{id}", rootHandler(handlers.UpdateAlbumHandler)).Methods("PUT")

	return routes
}
