package main

import (
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql/resolver"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := graphql.Config{
		Resolvers: &resolver.Resolver{},
	}
	executableSchema := graphql.NewExecutableSchema(config)
	graphqlServer := handler.NewDefaultServer(executableSchema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlServer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
