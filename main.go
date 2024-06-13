package main

import (
	"OzonTest/internal/database/in_memory"
	"OzonTest/internal/database/postgres"
	"OzonTest/internal/graph/generated"
	"log"
	"net/http"
	"os"

	"OzonTest/internal/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var repo graph.Repository
	// Выбор хранилища на основе переменной окружения
	storageType := os.Getenv("STORAGE_TYPE")
	switch storageType {
	case "postgres":
		repo = postgres.NewRepository()
	default:
		repo = in_memory.NewRepository()
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Repo: repo}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
