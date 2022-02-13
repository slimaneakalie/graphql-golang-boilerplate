package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql"
	"github.com/slimaneakalie/graphql-golang-boilerplate/internal/pkg/graphql/resolver"
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
	rootResolver := &resolver.Resolver{}
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
