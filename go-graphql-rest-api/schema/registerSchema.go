package schema

import (
	"github.com/digvijay17july/golang-projects/go-graphql-rest-api/repository"

	"github.com/graphql-go/graphql"
)

func GetUserType() *graphql.Object {
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"age": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
	return userType
}

func GetSchema(userType *graphql.Object) (graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type:        userType,
					Description: "Get a user by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.ID),
						},
					},
					Resolve: repository.GetUser,
				},
			},
		}),
	})

	return schema, err
}
