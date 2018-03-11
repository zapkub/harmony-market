package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func createGraphQLSchema() (graphql.Schema, error) {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema, err
}

func CreateGraphQLHandler() (*handler.Handler, error) {
	schema, err := createGraphQLSchema()
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return h, err
}
