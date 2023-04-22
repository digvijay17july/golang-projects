package schema

import (
	"io/ioutil"

	repository "github.com/digvijay17july/golang-projects/go-graphql-rest-api/repository"
	"github.com/graphql-go/graphql"
)

func RegisterSchema() ([]byte, error) {
	schemaData, err := ioutil.ReadFile("schema.graphql")
	return schemaData, err
}

func GetSchema() (graphql.Schema, error) {
	// Parse the schema definition into a GraphQL schema object
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type:        repository.User,
					Description: "Get a user by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.ID),
						},
					},
					Resolve: repository.GetUser(),
				},
			},
		}),
		Mutation: nil,
		Types:    []graphql.Type{userType},
	})
	return schema, err
}
