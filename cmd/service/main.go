package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql/resolver"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/service/book/metadata"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/service/book/reviews"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
)

const (
	RuntimeEnvVarName = "ENV"
	productionRuntime = "Production"
	portEnvVarName    = "PORT"
	defaultPort       = "8080"
	metadataApiURL    = "https://www.googleapis.com/books/v1/volumes"
	reviewsApiURL     = "https://www.goodreads.com/book/review_counts.json"
)

func main() {
	graphqlHandler := newGraphqlHandler()
	http.Handle("/query", graphqlHandler)

	runtime := os.Getenv(RuntimeEnvVarName)
	if runtime != productionRuntime {
		playgroundHandler := playground.Handler("GraphQL playground", "/query")
		http.Handle("/", playgroundHandler)
	}

	httpServerPort := getHttpServerPort()
	log.Printf("The http server is running on port %s", httpServerPort)
	log.Fatal(http.ListenAndServe(":"+httpServerPort, nil))
}

func newGraphqlHandler() http.Handler {
	resolverServices := &resolver.Services{
		Metadata: metadata.NewService(metadataApiURL),
		Reviews:  reviews.NewService(reviewsApiURL),
	}

	rootResolver := &resolver.Resolver{
		Services: resolverServices,
	}
	config := graphql.Config{
		Resolvers: rootResolver,
	}

	executableSchema := graphql.NewExecutableSchema(config)
	return handler.NewDefaultServer(executableSchema)
}

func getHttpServerPort() string {
	port := os.Getenv(portEnvVarName)
	if port == "" {
		port = defaultPort
	}

	return port
}
