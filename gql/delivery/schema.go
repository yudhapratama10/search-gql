package delivery

import (
	"github.com/graphql-go/graphql"
)

// NewSchema initializes Schema struct which takes resolver as the argument.
func NewSchema(articleResolver Resolver) Schema {
	return Schema{
		articleResolver: articleResolver,
	}
}

// Schema is struct which has method for Query and Mutation. Please init this struct using constructor function.
type Schema struct {
	articleResolver Resolver
}

var SearchResultGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResult",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"tournaments": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"has_stadium": &graphql.Field{
				Type: graphql.Boolean,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// Query initializes config schema query for graphql server.
func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Search": &graphql.Field{
				Type:        graphql.NewList(SearchResultGraphQL),
				Description: "Search Football Club",
				Args: graphql.FieldConfigArgument{
					"keyword": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"hasstadium": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"take": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: s.articleResolver.Search,
			},
			"Autocomplete": &graphql.Field{
				Type:        graphql.NewList(SearchResultGraphQL),
				Description: "Autocomplete for Query Football Club",
				Args: graphql.FieldConfigArgument{
					"keyword": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.articleResolver.Autocomplete,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
