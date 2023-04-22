package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type UserController struct {
	Schema graphql.Schema
}

func (userController *UserController) GetUser() http.Handler {
	return handler.New(&handler.Config{
		Schema:   &userController.Schema,
		Pretty:   true,
		GraphiQL: false,
	})
}
