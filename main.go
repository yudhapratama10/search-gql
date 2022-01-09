package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	gql_handler "github.com/yudhapratama10/search-gql/graph/handler"
	"github.com/yudhapratama10/search-gql/graph/repository"
	"github.com/yudhapratama10/search-gql/graph/usecase"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("No .env file specified.")
// 	}
// }

func main() {
	httpRouter := chi.NewRouter()
	//db := db.InitDB()

	httpRouter.Use(middleware.Logger)
	httpRouter.Use(render.SetContentType(render.ContentTypeJSON))
	httpRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"}, // Only allowing POST/GET because GQL only accept POST & GET
		AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	repo := repository.NewFootballRepository()
	usecase := usecase.NewFootballUsecase(repo)
	schema := gql_handler.NewSchema(gql_handler.NewResolver(usecase))

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Query(),
		//Mutation: schema.Mutation(),
	})

	if err != nil {
		log.Fatal(err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	http.Handle("/graphql", gqlHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
