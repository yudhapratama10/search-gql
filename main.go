package main

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file specified.")
	}
}

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

	httpRouter.Mount("/graphql", gqlHandler)
	//httpRouter.Get("/", productHandler.GetRoutes())

}
