package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/garethsharpe/gql/generated"
	"github.com/garethsharpe/gql/resolvers"
	"github.com/garethsharpe/gql/utils"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mw := utils.ChainMiddleware(
		utils.WithLogging,
		utils.WithTracing,
		utils.WithHeaders,
		utils.WithAppSvc,
	)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", mw(handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

