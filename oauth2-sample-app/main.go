package main

import (
	"log"
	"net/http"
	"os"

	security "github.com/digvijay17july/golang-projects/oauth2-sample-app/service"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var LOGGER zerolog.Logger

func main() {
	LOGGER = zerolog.New(os.Stdout).With().Timestamp().
		Str("app", "oauth2-example").
		Logger().With().Caller().Logger()
	service := security.Service{LOGGER: &LOGGER}
	router := mux.NewRouter()
	router.HandleFunc("/oauth2/fallback", service.FallbackOauth2)
	router.HandleFunc("/welcome", service.Authorise(Welcome())).Methods("GET")

	LOGGER.Info().Msg("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Welcome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		var response = "{Type: 'success', Data: 'data'}"

		w.Write([]byte(response))

	})
}
