package main

import (
	"log"
	"net/http"
	"os"

	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/api"
	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/schema"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var LOGGER zerolog.Logger

func main() {
	LOGGER = zerolog.New(os.Stdout).With().Timestamp().
		Str("app", "go-graphql-rst-api").
		Logger().With().Caller().Logger()

	userType := schema.GetUserType()

	schemaObj, err := schema.GetSchema(userType)
	if err != nil {
		LOGGER.Err(err).Msg(err.Error())
	}
	userController := &api.UserController{Schema: schemaObj}
	router := mux.NewRouter()
	router.Handle("/graphql", userController.GetUser())

	LOGGER.Info().Msg("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
