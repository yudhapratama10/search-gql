package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/yudhapratama10/search-gql/gql/delivery"
	"github.com/yudhapratama10/search-gql/gql/repository"
	"github.com/yudhapratama10/search-gql/gql/usecase"
)

func main() {
	repo := repository.NewFootballRepository()
	usecase := usecase.NewFootballUsecase(repo)
	schema := delivery.NewSchema(delivery.NewResolver(usecase))

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Query(),
		//Mutation: schema.Mutation(),
	})

	if err != nil {
		log.Fatal(err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:     &graphqlSchema,
		GraphiQL:   false,
		Pretty:     true,
		Playground: true,
	})

	http.Handle("/graphql", CorsMiddleware(gqlHandler))
	log.Println("Starting Service-GQL at port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// CORS Handler
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
