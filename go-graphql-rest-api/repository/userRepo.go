package repository

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	// In a real application, you would typically fetch the user data from a database or other data source
	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, fmt.Errorf("Missing id Argument")
	}
	return NewRandomUser(id), nil
}

func NewRandomUser(id string) User {
	return User{
		ID:    id,
		Name:  "Alice",
		Email: "alice@example.com",
	}
}
