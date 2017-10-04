package main

import (
	"sample-go-graphql-server/static/schema"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// define schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.RootQuery,
	})

	if err != nil {
		panic(err)
	}

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)
	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
